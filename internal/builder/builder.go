package builder

import "fmt"

func Build (ch chan bool) {
	for {
		//for {
		//	//
		//}
		fmt.Println("=> build")
		select {
		case <- ch:
		}
	}
}
