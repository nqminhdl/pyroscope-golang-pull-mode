package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	_ "net/http/pprof"
)

func hardWork(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Start: %v\n", time.Now())

	a := []string{}
	for i := 0; i < 500000; i++ {
		a = append(a, "aaaa")
	}

	time.Sleep(2 * time.Second)
	fmt.Printf("End: %v\n", time.Now())
}

func main() {
	var wg sync.WaitGroup

	go func() {
		fmt.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()

	wg.Add(1)
	wg.Add(1)
	go hardWork(&wg)
	wg.Wait()
}
