package main

import (
	"fmt"
	"sync"
	"time"
)



func main() {
	go say("world")
	say("hello")
	fmt.Println("\na-- End")


	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y, "\nb-- End")


	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch, "\nc-- End")

	c = make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("\nd-- End")

	//The select statement lets a goroutine wait on multiple communication operations.
	c = make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
		
	}()
	fibonaccii(c, quit)
	fmt.Println("\ne-- End")


	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	br :=false
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			br = true
			break
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
		if br == true{
			break
		}
	}
	fmt.Println("\nf-- End")

	sc := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go sc.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(sc.Value("somekey"), "\ng-- End")




}


// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}



func fibonaccii(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return			
		}
	}
}


func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}


func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

