package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	pid, _, _ := syscall.Syscall(syscall.SYS_GETPID, 0, 0, 0)
	uid, _, _ := syscall.Syscall(24, 0, 0, 0)

	fmt.Println("my pid is", pid)
	fmt.Println("user id", uid)

	message := []byte{'H', 'e', 'l', 'l', 'o', '!', '\n'}
	fileDescriptor := 1

	syscall.Write(fileDescriptor, message)

	command := "/bin/ls"
	env := os.Environ()

	syscall.Exec(command, []string{"ls", "-a", "-x"}, env)

}
