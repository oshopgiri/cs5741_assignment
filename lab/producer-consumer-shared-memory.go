package lab

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	Alphabets = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func producer(bowl []string, waitGroup *sync.WaitGroup) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		produce := string(Alphabets[rand.Intn(len(Alphabets))])
		fmt.Printf("Produced letter number %v: '%v'\n", i+1, produce)
		bowl[i] = produce
	}

	waitGroup.Done()
}

func consumer(bowl []string, waitGroup *sync.WaitGroup) {
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

func main() {
	bowl := make([]string, 100)
	waitGroup := new(sync.WaitGroup)

	waitGroup.Add(1)
	go producer(bowl, waitGroup)

	time.Sleep(5)

	waitGroup.Add(1)
	go consumer(bowl, waitGroup)

	waitGroup.Wait()
}
