package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func compute(name string, t int, pwg *sync.WaitGroup) {
	fmt.Printf("start to compute %s...\n", name)
	var td time.Duration
	if t == 0 {
		td = time.Duration(rand.Intn(2)) * time.Second
	} else {
		td = time.Duration(t) * time.Second
	}
	time.Sleep(td)
	fmt.Print("compute done.\n")
	pwg.Done()

}

func test_goroutine() {
	nameSlice := make([]string, 0)
	nameSlice = append(nameSlice, "xasxa", "xasxsa")
	var wg sync.WaitGroup
	for _, name := range nameSlice {
		wg.Add(1)
		go compute(name, 2, &wg)
	}
	wg.Wait()
	// var input string
	// fmt.Scanln(&input)
	fmt.Println("done")
}
