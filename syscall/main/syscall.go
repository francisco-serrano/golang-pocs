package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

func exampleA() {
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

func exampleB() {
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{Ptrace: true}

	var err error
	if err = cmd.Start(); err != nil {
		fmt.Println("Start:", err)
		return
	}

	err = cmd.Wait()

	fmt.Printf("State: %v\n", err)

	wpid := cmd.Process.Pid

	var r syscall.PtraceRegs
	if err = syscall.PtraceGetRegs(wpid, &r); err != nil {
		fmt.Println("PtraceGetRegs:", err)
		return
	}

	fmt.Printf("Registers: %#v\n", r)
	fmt.Printf("R15=%d, Gs=%d\n", r.R15, r.Gs)

	time.Sleep(2 * time.Second)
}

var maxSyscalls = 0

const SYSCALLFILE = "SYSCALLS"

func exampleC() {
	var SYSTEMCALLS []string

	f, err := os.Open(SYSCALLFILE)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, " ", "", -1)
		line = strings.Replace(line, "SYS_", "", -1)

		temp := strings.ToLower(strings.Split(line, "=")[0])

		SYSTEMCALLS = append(SYSTEMCALLS, temp)

		maxSyscalls++
	}

	COUNTER := make([]int, maxSyscalls)

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{Ptrace: true}

	err = cmd.Start()

	if err = cmd.Wait(); err != nil {
		fmt.Println("Wait:", err)
	}

	pid := cmd.Process.Pid

	fmt.Println("Process ID:", pid)

	var regs syscall.PtraceRegs
	before := true
	forCount := 0

	for {
		if before {
			if err := syscall.PtraceGetRegs(pid, &regs); err != nil {
				break
			}

			if regs.Orig_rax > uint64(maxSyscalls) {
				fmt.Println("Unknown:", regs.Orig_rax)
				return
			}

			COUNTER[regs.Orig_rax]++
			forCount++
		}

		if err = syscall.PtraceSyscall(pid, 0); err != nil {
			fmt.Println("PtraceSyscall:", err)
			return
		}

		if _, err := syscall.Wait4(pid, nil, 0, nil); err != nil {
			fmt.Println("Wait4:", err)
			return
		}

		before = !before
	}

	for i, x := range COUNTER {
		if x != 0 {
			fmt.Println(SYSTEMCALLS[i], "->", x)
		}
	}

	fmt.Println("Total system calls:", forCount)

}

func main() {
	//exampleA()
	//exampleB()
	exampleC()
}
