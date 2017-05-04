package channels

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
