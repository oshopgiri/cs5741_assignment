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

func producer(bowl chan string, waitGroup *sync.WaitGroup) {
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= 100; i++ {
		produce := string(Alphabets[rand.Intn(len(Alphabets))])
		fmt.Printf("Produced letter number %v: '%v'\n", i, produce)
		bowl <- produce
	}

	waitGroup.Done()
}

func consumer(bowl chan string, waitGroup *sync.WaitGroup) {
	for i := 1; i <= 100; i++ {
		consume := <-bowl
		fmt.Printf("Consumed letter number %v: '%v'\n", i, consume)
	}

	waitGroup.Done()
}

func main() {
	bowl := make(chan string)
	waitGroup := new(sync.WaitGroup)

	waitGroup.Add(1)
	go producer(bowl, waitGroup)

	time.Sleep(5)

	waitGroup.Add(1)
	go consumer(bowl, waitGroup)

	waitGroup.Wait()
}
