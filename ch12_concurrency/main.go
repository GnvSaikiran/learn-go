package main

import (
	"fmt"
	"math"
	"sync"
)

func exercise1() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i*10 + 1
		}
	}()
	go func() {
		wg.Wait()
		close(ch)
	}()

	var wg2 sync.WaitGroup
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		for v := range ch {
			fmt.Println(v)
		}
	}()
	wg2.Wait()
}

func exercise2() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		for i := 0; i < 10; i++ {
			ch2 <- i*100 + i
		}
		close(ch2)
	}()

	count := 2
	for count != 0 {
		select {
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil
				count--
				continue
			}
			fmt.Printf("Goroutine: 1, value: %d\n", v)
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				count--
				continue
			}
			fmt.Printf("Goroutine: 2, value: %d\n", v)
		}
	}
}

func exercise3() {
	mapFunc := sync.OnceValue(func() map[int]float64 {
		sqrtMap := make(map[int]float64, 100_000)
		for i := 0; i < 100_000; i++ {
			sqrtMap[i] = math.Sqrt(float64(i))
		}
		return sqrtMap
	})
	m := mapFunc()
	for i := 0; i < len(m); i += 1000 {
		fmt.Printf("%d: %.2f\n", i, m[i])
	}
}

func main() {
	exercise1()
	exercise2()
	exercise3()
}
