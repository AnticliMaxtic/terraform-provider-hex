package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Try tofu first
	if _, err := exec.LookPath("tofu"); err == nil {
		fmt.Println("Found tofu, formatting with OpenTofu...")
		cmd := exec.Command("tofu", "fmt", "-recursive", "../examples/")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error running tofu fmt: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// Try terraform as fallback
	if _, err := exec.LookPath("terraform"); err == nil {
		fmt.Println("Found terraform, formatting with Terraform...")
		cmd := exec.Command("terraform", "fmt", "-recursive", "../examples/")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error running terraform fmt: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// Neither found
	fmt.Println("Neither tofu nor terraform found, skipping formatting")
}
