package lab

import (
	"log"
	"runtime"
	"sync"
	"time"
)

func appender(data []int, value int, waitGroup *sync.WaitGroup) {
	data = append(data, value)
	waitGroup.Done()
}

func VaryingThreads() {
	waitGroup := new(sync.WaitGroup)

	var data1 []int
	start := time.Now()
	for i := 0; i < 1000; i++ {
		waitGroup.Add(1)
		go appender(data1, i, waitGroup)
	}
	waitGroup.Wait()
	elapsed := time.Since(start)
	fmt.Printf("With %d threads took %s", runtime.GOMAXPROCS(0), elapsed)

	runtime.GOMAXPROCS(2)
	var data2 []int
	start = time.Now()
	for i := 0; i < 1000; i++ {
		waitGroup.Add(1)
		go appender(data2, i, waitGroup)
	}
	waitGroup.Wait()
	elapsed = time.Since(start)
	fmt.Printf("With %d threads took %s", runtime.GOMAXPROCS(0), elapsed)
}
