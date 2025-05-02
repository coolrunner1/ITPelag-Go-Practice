package limiter

import (
	"fmt"
	"github.com/go-errors/errors"
	"sync"
	"time"
)

type LeakyBucket struct {
	bucketSize uint32
	//Queue of packets
	bucket []uint32
	mutex  sync.Mutex
	ticker *time.Ticker
	done   chan bool
}

func NewLeakyBucket(bucketSize uint32, tickRate time.Duration) *LeakyBucket {
	return &LeakyBucket{
		bucketSize: bucketSize,
		bucket:     make([]uint32, 0),
		ticker:     time.NewTicker(tickRate * time.Millisecond),
		done:       make(chan bool),
	}
}

func (lb *LeakyBucket) AddPacket(packetSize uint32) error {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()

	currentBucketSize := lb.getCurrentBucketSize()
	if currentBucketSize == lb.bucketSize {
		return errors.New("Bucket is full")
	} else if currentBucketSize+packetSize > lb.bucketSize {
		return errors.New("Packet is too large")
	}
	lb.pushPacket(packetSize)
	return nil
}

func (lb *LeakyBucket) Run() {
	for {
		select {
		case <-lb.done:
			fmt.Println("Leaky Bucket stopped")
			lb.ticker.Stop()
			return
		case <-lb.ticker.C:
			lb.mutex.Lock()
			if len(lb.bucket) > 0 {
				packet := lb.popPacket()
				fmt.Printf("Packet with a size of %d has been popped\n", packet)
			}
			lb.mutex.Unlock()
		}
	}
}

func (lb *LeakyBucket) Stop() {
	lb.done <- true
}

func (lb *LeakyBucket) getCurrentBucketSize() uint32 {
	var bucketSize uint32
	for _, el := range lb.bucket {
		bucketSize += el
	}
	return bucketSize
}

func (lb *LeakyBucket) pushPacket(packetSize uint32) {
	lb.bucket = append(lb.bucket, packetSize)
}

func (lb *LeakyBucket) popPacket() uint32 {
	if len(lb.bucket) == 0 {
		fmt.Println("Bucket is empty")
		return 0
	}
	el := lb.bucket[0]
	lb.bucket = lb.bucket[1:]
	return el
}
