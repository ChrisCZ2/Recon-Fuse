package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

// Function to display ASCII banner
func printBanner() {
	// Read the banner.txt file
	data, err := ioutil.ReadFile("banner.txt")
	if err != nil {
		log.Fatalf("Error reading banner file: %v", err)
	}
	// Print the banner
	fmt.Println(string(data))
}

// Main function
func main() {
	// Display the ASCII banner at the start
	printBanner()

	// Parse command-line flags
	installTools := flag.Bool("install", false, "Install Nmap, ffuf, and KiteRunner")
	nmapArgs := flag.String("nmap", "", "Arguments for Nmap")
	ffufArgs := flag.String("ffuf", "", "Arguments for ffuf")
	kiteArgs := flag.String("kite", "", "Arguments for KiteRunner")
	flag.Parse()

	// Check if user provides additional arguments directly
	args := flag.Args()

	// Install tools if the flag is set
	if *installTools {
		installAllTools()
		return
	}

	// Check for tool arguments
	if len(args) > 0 {
		if *nmapArgs != "" || strings.Contains(args[0], "nmap") {
			runNmap(strings.Join(args, " "))
			return
		} else if *ffufArgs != "" || strings.Contains(args[0], "ffuf") {
			runFFUF(strings.Join(args, " "))
			return
		} else if *kiteArgs != "" || strings.Contains(args[0], "kr") {
			runKiteRunner(strings.Join(args, " "))
			return
		}
	}

	// Run Nmap if nmap flag is set
	if *nmapArgs != "" {
		runNmap(*nmapArgs)
	}

	// Run ffuf if ffuf flag is set
	if *ffufArgs != "" {
		runFFUF(*ffufArgs)
	}

	// Run KiteRunner if kite flag is set
	if *kiteArgs != "" {
		runKiteRunner(*kiteArgs)
	}
}

// Function to install all tools (Nmap, ffuf, and KiteRunner)
func installAllTools() {
	fmt.Println("Starting installation of Nmap, ffuf, and KiteRunner...")

	// Install Nmap
	installNmap()

	// Install ffuf
	installFFUF()

	// Install KiteRunner manually
	installKiteRunner()

	fmt.Println("Installation of tools completed.")
}

// Function to install Nmap
func installNmap() {
	fmt.Println("Installing Nmap...")

	cmd := exec.Command("sudo", "apt-get", "install", "-y", "nmap")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error installing Nmap: %v\nOutput: %s", err, output)
	}
	fmt.Println("Nmap installed successfully.")
}

// Function to install ffuf
func installFFUF() {
	fmt.Println("Installing ffuf...")

	cmd := exec.Command("go", "install", "github.com/ffuf/ffuf@latest")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error installing ffuf: %v\nOutput: %s", err, output)
	}
	fmt.Println("ffuf installed successfully.")
}

// Function to install KiteRunner (Manual Installation Instructions)
func installKiteRunner() {
	fmt.Println("KiteRunner requires manual installation. Please follow these steps:")
	fmt.Println("1. git clone https://github.com/assetnote/kiterunner.git")
	fmt.Println("2. cd kiterunner && make")
	fmt.Println("3. sudo mv dist/kr /usr/local/bin/")
	fmt.Println("4. Verify installation by running 'kr --help'")
}

// Function to run Nmap with provided arguments
func runNmap(args string) {
	fmt.Println("Running Nmap with args:", args)

	nmapCmdArgs := strings.Fields(args)
	nmapCmd := exec.Command("nmap", nmapCmdArgs...)

	output, err := nmapCmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error running Nmap: %v\nOutput: %s", err, output)
	}
	fmt.Printf("Nmap Output:\n%s\n", output)
}

// Function to run ffuf with provided arguments
func runFFUF(args string) {
	fmt.Println("Running ffuf with args:", args)

	ffufCmdArgs := strings.Fields(args)
	ffufCmd := exec.Command("ffuf", ffufCmdArgs...)

	output, err := ffufCmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error running ffuf: %v\nOutput: %s", err, output)
	}
	fmt.Printf("ffuf Output:\n%s\n", output)
}

// Function to run KiteRunner with provided arguments
func runKiteRunner(args string) {
	fmt.Println("Running KiteRunner with args:", args)

	kiteCmdArgs := strings.Fields(args)
	kiteCmd := exec.Command("kr", kiteCmdArgs...)

	output, err := kiteCmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error running KiteRunner: %v\nOutput: %s", err, output)
	}
	fmt.Printf("KiteRunner Output:\n%s\n", output)
}
