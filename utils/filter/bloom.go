package filter

import (
	"fmt"
	"github.com/spaolacci/murmur3"
	"hash"
	"math"
	"sync"
)

type BloomFilter interface {
	Add(key []byte)
	Check(key []byte) bool
}

type bloomFilter struct {
	bitsetSize     uint32
	numberOfHashes uint32
	bitset         []bool
	bitsetMutex    sync.Mutex
}

func NewBloomFilter(expectedNumOfElements uint32, falsePositiveProbability float64) BloomFilter {
	bitsetSize := calculateBitsetSize(expectedNumOfElements, falsePositiveProbability)
	numberOfHashes := calculateNumberOfHashes(expectedNumOfElements, bitsetSize)
	bitset := make([]bool, bitsetSize)

	return &bloomFilter{
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

func (bf *bloomFilter) Add(key []byte) {
	var wg sync.WaitGroup
	var i uint32

	bitset := bf.bitset
	for i = 0; i < bf.numberOfHashes; i++ {
		wg.Add(1)
		go func(seed uint32) {
			defer wg.Done()
			hashEl := murmur3.New32WithSeed(seed)
			_, err := hashEl.Write(key)
			if err != nil {
				fmt.Println(err)
				return
			}
			index := hashEl.Sum32() % bf.bitsetSize

			bitset[index] = true
		}(i)
	}

	wg.Wait()

	bf.bitsetMutex.Lock()
	bf.bitset = bitset
	bf.bitsetMutex.Unlock()
}

func (bf *bloomFilter) Check(key []byte) bool {
	hashes := make([]hash.Hash32, bf.numberOfHashes)

	for i, hashElement := range hashes {
		hashElement = murmur3.New32WithSeed(uint32(i))
		_, err := hashElement.Write(key)
		if err != nil {
			fmt.Println(err)
			return false
		}
		index := hashElement.Sum32() % bf.bitsetSize
		if !bf.bitset[index] {
			return false
		}
	}
	return true
}
