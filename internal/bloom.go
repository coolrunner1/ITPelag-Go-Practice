package internal

import (
	"github.com/spaolacci/murmur3"
	"hash"
	"math"
)

type BloomFilter struct {
	bitsetSize     uint32
	numberOfHashes uint32
	bitset         []bool
	hashes         []hash.Hash32
}

func NewBloomFilter(expectedNumOfElements uint32, falsePositiveProbability float64) *BloomFilter {
	bitsetSize := calculateBitsetSize(expectedNumOfElements, falsePositiveProbability)
	numberOfHashes := calculateNumberOfHashes(expectedNumOfElements, bitsetSize)
	bitset := make([]bool, bitsetSize)
	hashes := make([]hash.Hash32, numberOfHashes)

	var i uint32

	for i = 0; i < numberOfHashes; i++ {
		hashes[i] = murmur3.New32WithSeed(i)
	}

	return &BloomFilter{
		bitsetSize:     bitsetSize,
		numberOfHashes: numberOfHashes,
		bitset:         bitset,
		hashes:         hashes,
	}
}

func calculateBitsetSize(expectedNumOfElements uint32, falsePositiveProbability float64) uint32 {
	return uint32(-1 * float64(expectedNumOfElements) * math.Log(falsePositiveProbability) / math.Pow(math.Log(2), 2))
}

func calculateNumberOfHashes(expectedNumOfElements, bitsetSize uint32) uint32 {
	return uint32(float64(bitsetSize/expectedNumOfElements) * math.Log(2))
}

func (bloomFilter *BloomFilter) Add(key []byte) {
	var i uint32
	for i = 0; i < bloomFilter.numberOfHashes; i++ {
		bloomFilter.hashes[i].Reset()
		_, _ = bloomFilter.hashes[i].Write(key)
		index := bloomFilter.hashes[i].Sum32() % bloomFilter.bitsetSize
		bloomFilter.bitset[index] = true
	}
}

func (bloomFilter *BloomFilter) Check(key []byte) bool {
	var i uint32
	for i = 0; i < bloomFilter.numberOfHashes; i++ {
		bloomFilter.hashes[i].Reset()
		_, _ = bloomFilter.hashes[i].Write(key)
		index := bloomFilter.hashes[i].Sum32() % bloomFilter.bitsetSize
		if !bloomFilter.bitset[index] {
			return false
		}
	}
	return true
}
