package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		parent()
	case "child":
		child()
	default:
		panic("please enter a correct command")
	}
}

func parent() {
	// fmt.Printf("Running %v\n", os.Args[2:])

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Hostname: %s\n", hostname)

	fmt.Printf("Parent running %v as %d\n", os.Args[2:], os.Getpid())

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		// UTS namespace which contains the hostname, lets us have our own hostname in the container
		Unshareflags: syscall.NEWNS, // dont share new mount namespace with the host
	}
	cmd.Run()
}

func child() {

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Hostname: %s\n", hostname)

	fmt.Printf("Child running %v as %d\n", os.Args[2:], os.Getpid())

	syscall.Sethostname([]byte("container"))

	syscall.Chroot("/docker-docker-but-worse")
	syscall.Chdir("/")
	syscall.Mount("proc", "proc", "proc", 0, "")

	new_hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("New Hostname: %s\n", new_hostname)

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
	syscall.Unmount("/proc", 0)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
