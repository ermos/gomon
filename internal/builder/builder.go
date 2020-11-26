package builder

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func Build (ch chan bool) {
	done := make(chan bool)
	chs := make(chan bool)
	start := time.Now()

	go func() {
		for {
			select {
			case <- ch:
				chs <- true
			}
		}
	}()

	go func() {
		for {
			select {
			case <- done:
				chs <- false
			}
		}
	}()

	fmt.Println("[gomon] Build binary..")
	l:for {
		execute(nil, done, "go", "run", "./")
		select {
		case v := <- chs:
			if v {
				fmt.Println("[gomon] Rebuild binary..")
			} else {
				break l
			}
		}
	}
	fmt.Printf("[gomon] Your program is done, time elapsed : %s\n", time.Since(start))
}

func execute(dir *string, done chan bool, name string, args... string) {
	ch := make(chan error)

	cmd := exec.Command(name, args...)
	if dir != nil {
		cmd.Dir = *dir
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	go func(){
		ch <- cmd.Run()
	}()

	select{
	case err := <- ch:
		if err != nil {
			log.Fatalf("command failed with %s", err)
		} else {
			done <- true
		}
	}
}

func clear () {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
