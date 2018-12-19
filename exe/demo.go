package main

import (
	"os/exec"
	"fmt"
	"time"
	"os"
)

func main() {
	cmd := exec.Command("cmd.exe")
	outPipe, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	cmd.Stdin = os.Stdin
	//inPipe, err := cmd.StdinPipe()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	err = cmd.Start()
	if err != nil {
		fmt.Println("cmd.Run()", err)
		return
	}
	//inPipe.Write([]byte("dir\n"))
	var buf = make([]byte, 1024)
	for ; ; {
		outPipe.Read(buf)
		time.Sleep(time.Second)
	}
	//fmt.Println([]byte(output))
}
