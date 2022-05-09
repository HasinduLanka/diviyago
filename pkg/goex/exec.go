package goex

import (
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

var wsroot string = "."

func ExcecTask(endTask chan bool, input []byte, program string, arg ...string) ([]byte, error) {

	println("Excecute Task " + program + " " + strings.Join(arg, " "))

	cmd, cmdErr := getCmd(program, arg...)

	if cmdErr != nil {
		return nil, cmdErr
	}

	cmd.Dir = wsroot

	// cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	stdin, stdinerr := cmd.StdinPipe()

	if stdinerr != nil {
		return nil, stdinerr
	}

	stdout, stdouterr := cmd.StdoutPipe()

	if stdouterr != nil {
		return nil, stdouterr
	}

	var output []byte

	go func() {
		stdoutput, stdoutputErr := io.ReadAll(stdout)
		if stdoutputErr != nil {
			log.Println("Error reading stdout: ", stdoutputErr)
		}

		output = stdoutput
	}()

	starterr := cmd.Start()
	if starterr != nil {
		log.Println(program+" >> command execute error: ", starterr)
	}

	go func() {
		_, stdinWrErr := stdin.Write(input)
		if stdinWrErr != nil {
			log.Println(program+" >> stdin write error : ", stdinWrErr)
		}

		stdin.Close()
	}()

	if endTask != nil {
		go func() {
			Kill := <-endTask

			if Kill {
				log.Println("killing ", program, cmd.Process.Signal(os.Kill))
			} else {
				log.Println("interrupting ", program, cmd.Process.Signal(os.Interrupt))
			}
		}()
	}

	waiterr := cmd.Wait()

	if waiterr != nil {
		log.Println(program+" >> command wait error: ", waiterr)
	}

	return output, nil
}

func ExcecProgramToString(program string, arg ...string) (string, error) {
	args := strings.Join(arg, " ")
	log.Println(program + " >> Excecute " + program + " " + args)

	cmd, cmdErr := getCmd(program, arg...)

	if cmdErr != nil {
		return "", cmdErr
	}

	// stdin, errStdin := cmd.StdinPipe()

	cmd.StdinPipe()
	cmd.Dir = wsroot
	// configure `Stdout` and `Stderr`
	cmd.Stderr = os.Stdout
	ret, err := cmd.Output()

	out := string(ret)
	return out, err
}

func getCmd(fileName string, arg ...string) (*exec.Cmd, error) {

	cmd := exec.Command(fileName, arg...)
	return cmd, nil
}
