package hello

import (
	"fmt"
	"time"
)

func SayHello() {
	fmt.Println("Hello")
	time.Sleep(2 * time.Second)
	fmt.Println("Good bye")
}