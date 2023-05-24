package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu      sync.Mutex
	counter map[string]int
}

func (ctn *Container) inc(name string) {
	ctn.mu.Lock()

	defer ctn.mu.Unlock()

	ctn.counter[name]++
}

func minMut() {
	var wg sync.WaitGroup

	container := &Container{
		// No need to initialize the mutex here, as we can make use of its default value
		counter: map[string]int{
			"0": 0,
			"1": 0,
			"2": 0,
			"3": 0,
		},
	}

	incrementCounter := func(name string, amount int) {
		for i := 0; i < amount; i++ {
			container.inc(name)
		}
		wg.Done()
	}

	for key := range container.counter {
		wg.Add(1)
		go incrementCounter(key, 2000)
	}

	wg.Wait()

	fmt.Println(container)
}
