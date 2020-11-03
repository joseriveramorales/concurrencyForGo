package main

import (
	"fmt"
	"time"
)

// Count to the number specified
func count(value int) {
	for i := 0; i <= value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}
