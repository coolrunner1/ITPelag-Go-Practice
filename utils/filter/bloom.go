package filter

import (
	"fmt"
	"github.com/spaolacci/murmur3"
	"hash"
	"math"
	"sync"
)

type BloomFilter struct {
	bitsetSize     uint32
	numberOfHashes uint32
	bitset         []bool
	bitsetMutex    sync.Mutex
}

func NewBloomFilter(expectedNumOfElements uint32, falsePositiveProbability float64) *BloomFilter {
	bitsetSize := calculateBitsetSize(expectedNumOfElements, falsePositiveProbability)
	numberOfHashes := calculateNumberOfHashes(expectedNumOfElements, bitsetSize)
	bitset := make([]bool, bitsetSize)

	return &BloomFilter{
		bitsetSize:     bitsetSize,
		numberOfHashes: numberOfHashes,
		bitset:         bitset,
	}
}

func calculateBitsetSize(expectedNumOfElements uint32, falsePositiveProbability float64) uint32 {
	numerator := -1 * float64(expectedNumOfElements) * math.Log(falsePositiveProbability)
	return uint32(numerator / math.Pow(math.Log(2), 2))
}

func calculateNumberOfHashes(expectedNumOfElements, bitsetSize uint32) uint32 {
	return uint32(float64(bitsetSize/expectedNumOfElements) * math.Log(2))
}

func (bloomFilter *BloomFilter) Add(key []byte) {
	hashes := make([]hash.Hash32, bloomFilter.numberOfHashes)
	var wg sync.WaitGroup
	for i := range bloomFilter.numberOfHashes {
		wg.Add(1)
		go func() {
			defer wg.Done()
			hashes[i] = murmur3.New32WithSeed(i)
			_, err := hashes[i].Write(key)
			if err != nil {
				fmt.Println(err)
				return
			}
			index := hashes[i].Sum32() % bloomFilter.bitsetSize

			bloomFilter.bitsetMutex.Lock()
			bloomFilter.bitset[index] = true
			bloomFilter.bitsetMutex.Unlock()
		}()
	}
	wg.Wait()
}

func (bloomFilter *BloomFilter) Check(key []byte) bool {
	hashes := make([]hash.Hash32, bloomFilter.numberOfHashes)
	for i := range bloomFilter.numberOfHashes {
		hashes[i] = murmur3.New32WithSeed(i)
		_, err := hashes[i].Write(key)
		if err != nil {
			fmt.Println(err)
			return false
		}
		index := hashes[i].Sum32() % bloomFilter.bitsetSize
		if !bloomFilter.bitset[index] {
			return false
		}
	}
	return true
}
