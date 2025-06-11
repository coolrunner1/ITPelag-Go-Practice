package middleware

import (
	"fmt"
	"github.com/coolrunner1/project/utils/filter"
	"github.com/coolrunner1/project/utils/limiter"
	"github.com/labstack/echo/v4"
	"sync"
	"time"
)

type LimiterMiddleware interface {
	Init(next echo.HandlerFunc) echo.HandlerFunc
}
type limiterMiddleware struct {
	bloomFilter  filter.BloomFilter
	leakyBuckets map[string]limiter.LeakyBucket
	mu           sync.Mutex
	bucketSize   uint32
	tickRate     time.Duration
}

func NewLimiterMiddleware(expectedNumOfUsers uint32, falsePositiveProbability float64, bucketSize uint32, tickRate time.Duration) LimiterMiddleware {
	return &limiterMiddleware{
		bloomFilter:  filter.NewBloomFilter(expectedNumOfUsers, falsePositiveProbability),
		leakyBuckets: make(map[string]limiter.LeakyBucket),
		bucketSize:   bucketSize,
		tickRate:     tickRate,
	}
}

func (lm *limiterMiddleware) Init(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ip := c.RealIP()

		if !lm.bloomFilter.Check([]byte(ip)) {
			lm.mu.Lock()
			lm.bloomFilter.Add([]byte(ip))
			fmt.Printf("New IP added to Bloom filter: %s\n", ip)
			lm.mu.Unlock()
		}

		lm.mu.Lock()
		bucket, exists := lm.leakyBuckets[ip]
		if !exists {
			bucket = limiter.NewLeakyBucket(lm.bucketSize, lm.tickRate)
			lm.leakyBuckets[ip] = bucket
			go bucket.Run()
		}
		lm.mu.Unlock()

		err := bucket.AddPacket(1)

		if err != nil {
			fmt.Printf("Rate limit exceeded for IP %s: %v\n", ip, err)
			return c.JSON(429, map[string]string{"message": "Too many requests"})
		}

		return next(c)
	}
}
