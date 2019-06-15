package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

/*
	系统函数调用
 */
func Shellout(command string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	var ShellToUse string;

	switch runtime.GOOS {
	case "linux":
		ShellToUse  = "bash"
	case "windows":
		ShellToUse = "powershell"
	default:
		ShellToUse = "bash"
	}

	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}

func main() {
	err, out, errout := Shellout("ls /")
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	fmt.Println(out)
	fmt.Println(errout)
}
