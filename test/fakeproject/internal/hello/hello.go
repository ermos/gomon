package hello

import (
	"fmt"
	"time"
)

func SayHello(args []string) {
	fmt.Println("fewefw")
	i := 0
	for {
		fmt.Println(i)
		i++
		time.Sleep(1 * time.Second)
	}
}