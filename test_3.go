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

func testGoroutine() {
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

func ping(pings chan<- string, s string) {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
		fmt.Print("\n")
		// td := time.Duration(1) * time.Second
		// time.Sleep(td)
		pings <- s
	}
	// pings <- s
}

func pong(pings <-chan string, pongs chan string) {
	fmt.Print("pong start\n")
	for i := 0; i < 10; i++ {
		td := time.Duration(10) * time.Second
		time.Sleep(td)
		s := <-pings
		fmt.Print("pong:", s, "\n")
		// pongs <- s
	}
	//close(pongs)
	// fmt.Println(<-pongs)
}

func testChannel() {
	// n := 3
	// // c := make(chan string) // 无缓冲信道
	// c := make(chan string, n+1) // 有缓冲信道
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func(n int, c chan string, pwg *sync.WaitGroup) {
	// 	if n == 0 {
	// 		n = rand.Intn(10)
	// 	}
	// 	fmt.Print("n:", n, "\n")
	// 	str := "channel-"
	// 	for i := 0; i < n; i++ {
	// 		c <- fmt.Sprintf("#%d %s %d", i, str, i)
	// 	}
	// 	wg.Done()
	// }(n, c, &wg)

	// wg.Wait()
	// for i := 0; i < n; i++ {
	// 	// for {
	// 	fmt.Printf("check: %s\n", <-c)
	// }

	pings := make(chan string)
	pongs := make(chan string)
	// pings := make(chan string, 1)
	// pongs := make(chan string, 1)
	go pong(pings, pongs)
	ping(pings, "ping message")
	for {
		if len(pongs) != 0 {
			time.Sleep(time.Duration(1) * time.Second)
			// fmt.Println("main:", <-pongs)
		}
	}
}
