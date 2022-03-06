package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	argNum := len(os.Args)
	if argNum < 2 {
		log.Printf("Usage:%s command\n", os.Args[0])
		os.Exit(1)
	}

	arg := []string{}
	if argNum > 2 {
		arg = os.Args[2:]
	}

	mainOutput(UsePipe(os.Args[1], arg...))
}

func mainOutput(out []byte, err error) {
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The output of command %q is\n%s\n", os.Args[1], out)
}

// 直接给 Cmd.Stdout 赋值
func FillStd(name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)
	var out = new(bytes.Buffer)

	cmd.Stdout = out
	cmd.Stderr = out

	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

func UseOutput(name string, arg ...string) ([]byte, error) {
	return exec.Command(name, arg...).Output()
}

// 使用 Pipe
func UsePipe(name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err = cmd.Start(); err != nil {
		return nil, err
	}

	var out = make([]byte, 0, 1024)
	for {
		tmp := make([]byte, 128)
		n, err := stdout.Read(tmp)
		out = append(out, tmp[:n]...)
		if err != nil {
			break
		}
	}

	if err = cmd.Wait(); err != nil {
		return nil, err
	}

	return out, nil
}
