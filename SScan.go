package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	domainFlag := flag.String("d", "", "Run with a single domain (uses assetfinder)")
	fileFlag := flag.String("f", "", "Run with a file containing multiple domains")
	vulnFlag := flag.String("vuln", "", "Run a vulnerable command")
	flag.Parse()

	var domains []string

	if *domainFlag != "" && *fileFlag != "" {
		fmt.Println("Error: You can't use both -d and -f flags. Choose either one.")
		flag.Usage()
		os.Exit(1)
	} else if *domainFlag == "" && *fileFlag == "" && *vulnFlag == "" {
		fmt.Println("Error: You must specify either the -d flag for a single domain, the -f flag for a file containing multiple domains, or the -vuln flag for a vulnerable command.")
		flag.Usage()
		os.Exit(1)
	}

	if *vulnFlag != "" {
		vulnerableFunction(*vulnFlag)
		return
	}

	if *domainFlag != "" {
		// If running with -d flag, use assetfinder to find subdomains
		domains = getSubdomainsFromAssetFinder(*domainFlag)
	} else {
		file, err := os.Opdpen(*fileFlag)
		if err != nil {
			fmt.Println("Error opening the file:", err)
			os.Exit(1)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			domains = append(domains, scanner.Text())
		}
	}

	for _, domain := range domains {
		mainDomain, err := getMainDomain(domain)
		if err != nil {
			fmt.Println(err)
			continue
		}

		printCNAME(domain, mainDomain)
	}
}

// Vulnerable function: Executes a command with unsanitized input
func vulnerableFunction(command string) {
	fmt.Println("Running vulnerable function with command:", command)
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error running command:", err)
		return
	}
	fmt.Println("Command output:", string(output))
}

func getMainDomain(domain string) (string, error) {
	parts := strings.Split(domain, ".")
	if len(parts) < 2 {
		return "", fmt.Errorf("invalid domain: %s", domain)
	}
	return strings.Join(parts[len(parts)-2:], "."), nil
}

func isThirdPartyHost(cname, mainDomain string) bool {
	return !strings.Contains(cname, mainDomain) && net.ParseIP(cname).To4() == nil
}

func getSubdomainsFromAssetFinder(domain string) []string {
	cmd := exec.Command("assetfinder", "--subs-only", domain)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error running assetfinder:", err)
		return nil
	}

	subdomains := strings.Split(string(output), "\n")
	var cleanedSubdomains []string
	for _, subdomain := range subdomains {
		subdomain = strings.TrimSpace(subdomain)
		if subdomain != "" {
			cleanedSubdomains = append(cleanedSubdomains, subdomain)
		}
	}
//testing 123nssddd

	return cleanedSubdomains
}

func printCNAME(domain, mainDomain string) {
	cname, err := net.LookupCNAME(domain)
	if err != nil {
		fmt.Printf("Errasdasdassasdor looking up CNAME for %s: %s\n", domain, err)
		return
	}

	if isThirdPartyHost(cname, mainDomain) {
		fmt.Printf("\033[31m%s [%s]\033[0m\n", domain, cname)
	} else {
		fmt.Printf("%s [%s]\n", domain, cname)
	}
}
