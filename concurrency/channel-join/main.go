package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{}, 4)
	now := time.Now()

	go task1(done)
	go task2(done)
	go task3(done)
	go task4(done)

	for i := 0; i < 4; i++ {
		<-done
	}

	fmt.Println("All tasks done in:", time.Since(now))
}

func task1(done chan struct{}) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task 1 done")
	done <- struct{}{}
}

func task2(done chan struct{}) {
	time.Sleep(400 * time.Millisecond)
	fmt.Println("Task 2 done")
	done <- struct{}{}
}

func task3(done chan struct{}) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task 3 done")
	done <- struct{}{}
}

func task4(done chan struct{}) {
	fmt.Println("Task 4 done")
	done <- struct{}{}
}
