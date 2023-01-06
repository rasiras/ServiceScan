
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: SScan <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var subdomains []string
	for scanner.Scan() {
		subdomains = append(subdomains, scanner.Text())
	}

	mainDomain, err := getMainDomain(subdomains[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, subdomain := range subdomains {
		cname, err := net.LookupCNAME(subdomain)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if isThirdPartyHost(cname, mainDomain) {
			fmt.Printf("\033[31m%s [%s]\033[0m\n", subdomain, cname)
		} else {
			fmt.Printf("%s [%s]\n", subdomain, cname)
		}
	}
}

func getMainDomain(subdomain string) (string, error) {
	parts := strings.Split(subdomain, ".")
	if len(parts) < 2 {
		return "", fmt.Errorf("invalid subdomain: %s", subdomain)
	}
	return strings.Join(parts[len(parts)-2:], "."), nil
}

func isThirdPartyHost(cname, mainDomain string) bool {
	return !strings.Contains(cname, mainDomain) && net.ParseIP(cname).To4() == nil
}
