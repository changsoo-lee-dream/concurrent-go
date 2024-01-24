package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var value int32
	atomic.AddInt32(&value, 5)

	fmt.Println(atomic.LoadInt32(&value))
}
