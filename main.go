package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func sendChanChanChan(c chan chan chan chan int) {
	for range factor {
		go func() {
			fmt.Println("starting 3chan producer")
			_3chan := make(chan chan chan int)
			sendChanChan(c, _3chan)
		}()
	}
}

func receiveChanChanChan(c chan chan chan chan int) {
	for _3chan := range c {
		fmt.Println("got message from 4chan")
		for range factor {
			fmt.Println("starting 3chan consumer")
			go receiveChanChan(_3chan)
		}
	}
}

func sendChanChan(_4chan chan chan chan chan int, _3chan chan chan chan int) {
	_4chan <- _3chan
	for range factor {
		go func() {
			fmt.Println("starting 2chan producer")
			_2chan := make(chan chan int)
			sendChan(_3chan, _2chan)
		}()
	}
}

func receiveChanChan(c chan chan chan int) {
	for _2chan := range c {
		fmt.Println("got message from 3chan")
		for range factor {
			fmt.Println("starting 2chan consumer")
			go receiveChan(_2chan)
		}
	}
}

func sendChan(_3chan chan chan chan int, _2chan chan chan int) {
	_3chan <- _2chan
	for range factor {
		go func() {
			fmt.Println("starting 1chan producer")
			_1chan := make(chan int)
			send(_2chan, _1chan)
		}()
	}
}

func receiveChan(c chan chan int) {
	for _1chan := range c {
		fmt.Println("got message from 2chan")
		for range factor {
			fmt.Println("starting 1chan consumer")
			go receive(_1chan)
		}
	}
}

func send(_2chan chan chan int, _1chan chan int) {
	_2chan <- _1chan
	for range factor {
		go func() {
			fmt.Println("starting int producer")
			for range factor {
				go func() {
					fmt.Println("sending int")
					_1chan <- 1
				}()
			}
		}()
	}
}

func receive(c chan int) {
	for s := range c {
		fmt.Println("received int")
		sum.Add(int32(s))
	}
}

const factor = 3

var sum = &atomic.Int32{}

func main() {
	_4chan := make(chan chan chan chan int)

	go sendChanChanChan(_4chan)
	go receiveChanChanChan(_4chan)

	time.Sleep(500 * time.Millisecond)

	fmt.Printf("%d ^ 5: %d", factor, sum.Load())
}
