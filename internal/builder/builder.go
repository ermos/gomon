package builder

import (
	"errors"
	"fmt"
	"github.com/ermos/chalk"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

type runner struct {
	command 	*exec.Cmd
	startedAt	time.Time
	dst 		io.Writer
}

func Build (ch chan string, args []string) {
	r := runner{}
	stringArgs := strings.Join(args, " ")

	r.setWriter(os.Stdout)

	fmt.Printf("[%s] Starting compilation in watch mode..\n", time.Now().Format("2006-01-02 15:04:05"))
	l:for {
		clear()
		fmt.Println(chalk.Green("[gomon]"), chalk.Yellow("1.0.0"))
		fmt.Println(
			chalk.Green("[gomon]"),
			chalk.Yellow("watching dir :"),
			chalk.Magenta("toto,tata,ok"),
			)
		fmt.Println(
			chalk.Green("[gomon]"),
			chalk.Yellow("watching file's type :"),
			chalk.Magenta("go,js,json"),
			)
		fmt.Println(chalk.Green("[gomon]"), chalk.Blue("go run ./ ", stringArgs))
		err := r.Run("go", "run", "./", stringArgs)
		if err != nil {
			log.Fatal(err)
		}
		select {
		case v := <- ch:
			if v == "reload" {
				err = r.Kill()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(chalk.Green("[gomon]"), chalk.Magenta("Recompile.."))
			} else {
				break l
			}
		}
	}
	fmt.Println(chalk.Green("[gomon]"), chalk.Magentaf("Your program is done, time elapsed : %s", time.Since(r.startedAt).String()))
}

func clear () {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (r *runner) Run(name string, args... string) error {
	if r.command != nil && r.command.ProcessState != nil && r.command.ProcessState.Exited() {
		return errors.New("last process is not end")
	}

	r.command = exec.Command(name, args...)

	stdout, err := r.command.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := r.command.StderrPipe()
	if err != nil {
		return err
	}

	err = r.command.Start()
	if err != nil {
		return err
	}

	r.startedAt = time.Now()

	go io.Copy(r.dst, stdout)
	go io.Copy(r.dst, stderr)
	go r.command.Wait()

	return nil
}

func (r *runner) Kill() error {
	if r.command == nil || r.command.Process == nil  {
		return nil
	}

	if err := r.command.Process.Kill(); err != nil {
		return err
	}

	r.command = nil

	return nil
}

func (r *runner) setWriter(writer io.Writer) {
	r.dst = writer
}