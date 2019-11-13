package lab

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func producerChannel(bowl chan string, waitGroup *sync.WaitGroup) {
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= 100; i++ {
		produce := string(Alphabets[rand.Intn(len(Alphabets))])
		fmt.Printf("Produced letter number %v: '%v'\n", i, produce)
		bowl <- produce
	}

	waitGroup.Done()
}

func consumerChannel(bowl chan string, waitGroup *sync.WaitGroup) {
	for i := 1; i <= 100; i++ {
		consume := <-bowl
		fmt.Printf("Consumed letter number %v: '%v'\n", i, consume)
	}

	waitGroup.Done()
}

func ProducerConsumerChannel() {
	bowl := make(chan string)
	waitGroup := new(sync.WaitGroup)

	waitGroup.Add(1)
	go producerChannel(bowl, waitGroup)

	time.Sleep(5)

	waitGroup.Add(1)
	go consumerChannel(bowl, waitGroup)

	waitGroup.Wait()
}
