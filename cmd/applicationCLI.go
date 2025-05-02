package cmd

import (
	"fmt"
	"github.com/coolrunner1/project/utils/filter"
	"github.com/coolrunner1/project/utils/limiter"
	"github.com/go-errors/errors"
	"github.com/spf13/cobra"
	"time"
)

func ApplicationCliInit() {
	rootCmd := &cobra.Command{
		Use:   "project",
		Short: "Use bloom argument to use the bloom filter",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Use bloom argument to use the bloom filter")
		},
	}

	bloomTestCmd := &cobra.Command{
		Use: "bloomTest",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello World")
			bloomFilter := filter.NewBloomFilter(10, 0.01)
			bloomFilter.Add([]byte("1232"))
			fmt.Println(bloomFilter.Check([]byte("test")))
			fmt.Println(bloomFilter.Check([]byte("dsfgfsdf324")))
			fmt.Println(bloomFilter.Check([]byte("1232")))
		},
	}

	rootCmd.AddCommand(bloomTestCmd)

	leakyBucketTestCmd := &cobra.Command{
		Use: "leakyBucketTest",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Leaky Bucket Test")
			leakyBucket := limiter.NewLeakyBucket(2000, 10)
			err := leakyBucket.AddPacket(100)
			if err != nil {
				fmt.Println(err.(*errors.Error).ErrorStack())
				return
			}
			err = leakyBucket.AddPacket(100)
			if err != nil {
				fmt.Println(err.(*errors.Error).ErrorStack())
				return
			}
			err = leakyBucket.AddPacket(1000)
			if err != nil {
				fmt.Println(err.(*errors.Error).ErrorStack())
				return
			}

			go leakyBucket.Run()
			time.Sleep(1 * time.Second)
			err = leakyBucket.AddPacket(200)
			if err != nil {
				fmt.Println(err.(*errors.Error).ErrorStack())
				return
			}
			err = leakyBucket.AddPacket(1000)
			if err != nil {
				fmt.Println(err.(*errors.Error).ErrorStack())
				return
			}
			time.Sleep(1 * time.Second)
			leakyBucket.Stop()
		},
	}

	rootCmd.AddCommand(leakyBucketTestCmd)

	bloomCmd := &cobra.Command{
		Use: "bloom",
		Run: func(cmd *cobra.Command, args []string) {
			var expectedNumOfElements uint32
			fmt.Println("Enter the expected number of elements")
			_, err := fmt.Scan(&expectedNumOfElements)
			if err != nil {
				fmt.Println(err)
				return
			}

			var falsePositiveProbability float64
			fmt.Println("Enter the probability of false positive results")
			_, err = fmt.Scan(&falsePositiveProbability)
			if err != nil {
				fmt.Println(err)
				return
			}
			bloomFilter := filter.NewBloomFilter(10, 0.01)
			for {
				fmt.Println("1 - Add a new rule")
				fmt.Println("2 - Check if the rule exists")
				fmt.Println("0 - Quit the application")

				var choice uint32
				_, err := fmt.Scan(&choice)
				if err != nil {
					fmt.Println(err)
					return
				}

				switch choice {
				case 1:
					fmt.Println("Enter the new rule")
					var newRule []byte
					_, err := fmt.Scan(&newRule)
					if err != nil {
						fmt.Println(err)
						return
					}
					bloomFilter.Add(newRule)
				case 2:
					fmt.Println("Enter a string to check")
					var newString []byte
					_, err := fmt.Scan(&newString)
					if err != nil {
						fmt.Println(err)
						return
					}
					if bloomFilter.Check(newString) {
						fmt.Println("Rule exists")
					} else {
						fmt.Println("Rule does not exist")
					}
				case 0:
					return
				default:
					fmt.Println("Invalid input")
				}
			}
		},
	}

	rootCmd.AddCommand(bloomCmd)

	leakyBucketCmd := &cobra.Command{
		Use: "leakyBucket",
		Run: func(cmd *cobra.Command, args []string) {
			var bucketSize uint32
			fmt.Println("Enter the bucket size")
			_, err := fmt.Scan(&bucketSize)
			if err != nil {
				fmt.Println(err)
				return
			}

			var tickRate time.Duration
			fmt.Println("Enter the tick rate")
			_, err = fmt.Scan(&tickRate)
			if err != nil {
				fmt.Println(err)
				return
			}
			leakyBucket := limiter.NewLeakyBucket(bucketSize, tickRate)
			for {
				fmt.Println("1 - Send a new packet")
				fmt.Println("2 - Start the leaky bucket")
				fmt.Println("3 - Stop the leaky bucket")
				fmt.Println("0 - Quit the application")

				var choice uint32
				_, err := fmt.Scan(&choice)
				if err != nil {
					fmt.Println(err)
					return
				}

				switch choice {
				case 1:
					fmt.Println("Enter the size of the packet")
					var packetSize uint32
					_, err := fmt.Scan(&packetSize)
					if err != nil {
						fmt.Println(err)
						return
					}
					err = leakyBucket.AddPacket(packetSize)
					if err != nil {
						fmt.Println(err.(*errors.Error).ErrorStack())
						return
					}
				case 2:
					go leakyBucket.Run()
				case 3:
					leakyBucket.Stop()
				case 0:
					return
				default:
					fmt.Println("Invalid input")
				}
			}
		},
	}

	rootCmd.AddCommand(leakyBucketCmd)

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		return
	}
}
