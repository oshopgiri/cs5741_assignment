package lab

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func producerSharedMemory(bowl []string, waitGroup *sync.WaitGroup) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		produce := string(Alphabets[rand.Intn(len(Alphabets))])
		fmt.Printf("Produced letter number %v: '%v'\n", i+1, produce)
		bowl[i] = produce
	}

	waitGroup.Done()
}

func consumerSharedMemory(bowl []string, waitGroup *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		consume := bowl[i]
		if consume == "" {
			time.Sleep(5)
			i--
			continue
		}
		fmt.Printf("Consumed letter number %v: '%v'\n", i+1, consume)
	}

	waitGroup.Done()
}

func producerConsumerSharedMemory() {
	bowl := make([]string, 100)
	waitGroup := new(sync.WaitGroup)

	waitGroup.Add(1)
	go producerSharedMemory(bowl, waitGroup)

	time.Sleep(5)

	waitGroup.Add(1)
	go consumerSharedMemory(bowl, waitGroup)

	waitGroup.Wait()
}
