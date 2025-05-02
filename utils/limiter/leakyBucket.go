package limiter

import (
	"fmt"
	"sync"
	"time"
)

type LeakyBucket struct {
	bucketSize uint32
	//Queue of packets
	bucket   []uint32
	mutex    sync.Mutex
	tickRate time.Duration
}

func NewLeakyBucket(bucketSize uint32) *LeakyBucket {
	return &LeakyBucket{
		bucketSize: bucketSize,
		bucket:     make([]uint32, bucketSize),
	}
}

func (lb *LeakyBucket) AddPacket(packetSize uint32) {
	if lb.getCurrentBucketSize() >= lb.bucketSize {
		fmt.Println("Bucket is full")
		return
	}
	lb.pushPacket(packetSize)
}

func (lb *LeakyBucket) Run() {
	for {
		if len(lb.bucket) == 0 {
			fmt.Println("Bucket is empty\n Exiting...")
			return
		}
	}
}

func (lb *LeakyBucket) getCurrentBucketSize() uint32 {
	var bucketSize uint32
	for i := range lb.bucket {
		bucketSize += lb.bucket[i]
	}
	return bucketSize
}

func (lb *LeakyBucket) pushPacket(packetSize uint32) {
	lb.bucket = append(lb.bucket, packetSize)
}

func (lb *LeakyBucket) popPacket() uint32 {
	el := lb.bucket[0]
	lb.bucket = lb.bucket[1:]
	return el
}
