package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("halløj")
	cta := CreateCancellationToken("a")
	ctb := CreateCancellationToken("b")
	go f1(cta, "A")
	time.Sleep(50 * time.Millisecond)
	go f2(cta, "B")
	time.Sleep(50 * time.Millisecond)
	go f1(ctb, "C")
	time.Sleep(2 * time.Second)
	cta.Cancel()
	time.Sleep(2 * time.Second)
	ctb.Cancel()
}

func f1(ct *CancellationToken, name string) {
	for {
		fmt.Println(name, "f1")
		time.Sleep(time.Second/3)
		if ct.IsCancelled() {
			fmt.Println("ending", name, "f1")
			return
		}
	}
}

func f2(ct *CancellationToken, name string) {
	for {
		fmt.Println(name, "f2")
		time.Sleep(time.Second/3)
		if ct.IsCancelled() {
			fmt.Println("ending", name, "f2")
			return
		}
	}
}



type CancellationToken struct{
	ch chan struct{}
	name string
}

func CreateCancellationToken(name string) *CancellationToken {
	return &CancellationToken{
		ch: make(chan struct{}),
		name: name}

}
func (ct *CancellationToken) Cancel() {
	fmt.Println("cancelling token", ct.name)
	if ! ct.IsCancelled() {
		close(ct.ch)
	}
}
func (ct *CancellationToken) ThrowIfCancelled() {
	if ct.IsCancelled() {
		panic(fmt.Sprintf("Throw because cancelled %s", ct.name))
	}
}

func (ct *CancellationToken) IsCancelled() bool {
	// fmt.Println("checking ct", ct.name)
	select {
	case <-ct.ch:
		return true
	default:
		return false
	}
}
