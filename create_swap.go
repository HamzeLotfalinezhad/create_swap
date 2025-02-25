package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func runCommand(cmd string, args ...string) string {
	command := exec.Command(cmd, args...)
	output, err := command.CombinedOutput()
	if err != nil {
		log.Fatalf("Error running %s: %v", cmd, err)
	}
	return strings.TrimSpace(string(output))
}

func createSwap(size string) {
	fmt.Println("\n🔍 Checking existing swap:")
	fmt.Println(runCommand("sudo", "swapon", "--show"))

	fmt.Printf("\n Creating %s swap file...\n", size)

	commands := [][]string{
		{"fallocate", "-l", size, "/swapfile"},
		{"chmod", "600", "/swapfile"},
		{"mkswap", "/swapfile"},
		{"swapon", "/swapfile"},
	}

	for _, cmd := range commands {
		runCommand(cmd[0], cmd[1:]...)
	}

	runCommand("bash", "-c", "echo '/swapfile none swap sw 0 0' | sudo tee -a /etc/fstab")
	runCommand("bash", "-c", "echo 'vm.swappiness=10' | sudo tee -a /etc/sysctl.conf")
	runCommand("sysctl", "-p")

	fmt.Println("\n Swap file created successfully!")
	fmt.Println("\n🔍 Checking updated swap:")
	fmt.Println(runCommand("sudo", "swapon", "--show"))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter swap size (e.g., 2G): ")
	size, _ := reader.ReadString('\n')
	size = strings.TrimSpace(size)

	createSwap(size)
}
