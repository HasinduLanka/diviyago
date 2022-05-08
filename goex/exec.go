package goex

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var wsroot string = "."

func ExcecProgram(program string, arg ...string) (string, error) {
	args := strings.Join(arg, " ")
	println("Excecute " + program + " " + args)

	cmd, cmdErr := GetCmdFromEmbeded(program, arg...)

	if cmdErr != nil {
		return "", cmdErr
	}

	cmd.Dir = wsroot
	// configure `Stdout` and `Stderr`
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	err := cmd.Run()
	// run command
	if err != nil {
		fmt.Println("Error:", err)
	}

	// out := string(ret)
	return "Done Excecute " + program + " " + args, err
}

func OpenProgram(program string, arg ...string) (string, error) {
	args := strings.Join(arg, " ")
	println("Excecute " + program + " " + args)

	cmd, cmdErr := GetCmdFromEmbeded(program, arg...)

	if cmdErr != nil {
		return "", cmdErr
	}

	cmd.Dir = wsroot

	err := cmd.Run()
	// run command
	if err != nil {
		fmt.Println("Error:", err)
	}

	// out := string(ret)
	return "Done Opening " + program + " " + args, err
}

func ExcecCmdTask(command string, endTask chan bool) (string, error) {
	return ExcecTask("sh", endTask, "-c", command)
}

func ExcecTask(program string, endTask chan bool, arg ...string) (string, error) {
	args := strings.Join(arg, " ")
	println("Excecute Task " + program + " " + args)

	cmd, cmdErr := GetCmdFromEmbeded(program, arg...)

	if cmdErr != nil {
		return "", cmdErr
	}

	cmd.Dir = wsroot
	// configure `Stdout` and `Stderr`
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	err := cmd.Start()
	// run command
	if err != nil {
		fmt.Println("Error:", err)
	}

	Kill := <-endTask

	if Kill {
		log.Println(cmd.Process.Signal(os.Kill))
	} else {
		log.Println(cmd.Process.Signal(os.Interrupt))
	}

	// out := string(ret)
	return "Done Excecute Task " + program + " " + args, err
}

func ExcecProgramToString(program string, arg ...string) (string, error) {
	args := strings.Join(arg, " ")
	println("Excecute " + program + " " + args)

	cmd, cmdErr := GetCmdFromEmbeded(program, arg...)

	if cmdErr != nil {
		return "", cmdErr
	}

	cmd.Dir = wsroot
	// configure `Stdout` and `Stderr`
	cmd.Stderr = os.Stdout
	ret, err := cmd.Output()

	out := string(ret)
	return out, err
}

func GetCmdFromEmbeded(fileName string, arg ...string) (*exec.Cmd, error) {
	// bin, binErr := GetFile(fileName)

	// bin, binErr :=
	// if binErr != nil {
	// 	return nil, binErr
	// }

	// exe, memexerr := memexec.New(bin)
	// if memexerr != nil {
	// 	return nil, memexerr
	// }

	// // defer exe.Close()

	// cmd := exe.Command(arg...)
	cmd := exec.Command(fileName, arg...)

	return cmd, nil

}
