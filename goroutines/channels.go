package goroutines

import (
	"fmt"
	"time"
)

// Channel send msg to channel
func Channel() {
	message := make(chan string)
	go func() { message <- "ping" }()
	fmt.Println(message)
	msg := <-message
	fmt.Println(msg)
}

// ChannelBuff channel with buff
func ChannelBuff() {
	messages := make(chan int, 2)
	messages <- 1
	messages <- 2
	fmt.Println(messages)
	fmt.Println(<-messages)
	fmt.Println(messages)
	fmt.Println(<-messages)
}

func worker(done chan bool) {
	fmt.Print("Working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

// ChannelSync use channel communicate
func ChannelSync() {
	done := make(chan bool, 1)
	go worker(done)
	<-done
}

func receiver(ping chan<- string, msg string) {
	ping <- msg
}

func deliver(ping <-chan string, pong chan<- string) {
	msg := <-ping
	pong <- msg
}

// ChannelTrans driection of chan
func ChannelTrans() {
	ping := make(chan string, 1)
	pong := make(chan string, 2)
	receiver(ping, "hello")
	deliver(ping, pong)
	// fmt.Println(<-ping)
	fmt.Println(<-pong)
}

func counter(in chan<- int) {
	for i := 0; i < 10; i++ {
		in <- i
		fmt.Println("in", i)
		time.Sleep(time.Second * 2)
	}
	close(in)
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i
		fmt.Println("out", i)
		time.Sleep(time.Second * 2)
	}
	close(out)
}

// ChannelTrans2 direction of chan
func ChannelTrans2() {
	in := make(chan int, 10)
	out := make(chan int, 10)
	go counter(in)
	go squarer(out, in)
	for i := range out {
		fmt.Println("result", i)
	}
}

// ChannelSelect select channel
func ChannelSelect() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("c1: ", msg1)
		case msg2 := <-c2:
			fmt.Println("c2: ", msg2)
		}
	}
}

// ChannelTimeout channel with timeout
func ChannelTimeout() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "c1"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "c2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout")
	}
}

// ChannelClose close channel
func ChannelClose() {
	jobs := make(chan int, 3)
	done := make(chan bool, 1)
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("receive", j, more)
			} else {
				fmt.Println("receive all")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		jobs <- i
	}

	for i := 0; i < 2; i++ {
		jobs <- i
	}
	close(jobs)
}

// ChannelRange range of channel
func ChannelRange() {
	channel := make(chan int, 2)
	channel <- 1
	channel <- 2
	close(channel)
	for i := range channel {
		fmt.Println(i)
	}
}

// ChannelRange2 range of channel test
func ChannelRange2() {
	channel := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
			fmt.Println("In channel", i)
			time.Sleep(time.Second * 2)
		}
		close(channel)
	}()

	for j := range channel {
		fmt.Println("Out channel", j)
	}
}

func worker2(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second * 10)
		result <- j
		fmt.Println("worker", id, "finished job", j)
	}
}

// WorkerPool worker pool
func WorkerPool() {
	jobs := make(chan int, 10)
	result := make(chan int, 10)

	for i := 0; i < 3; i++ {
		go worker2(i, jobs, result)
	}

	for j := 0; j < 5; j++ {
		jobs <- j
	}
	close(jobs)

	for j := 0; j < 5; j++ {
		<-result
	}
}
