

# ServiceScan Tool

ServiceScan is a command-line tool written in Go that helps you find subdomains and their CNAME records for a given domain. It uses the `net` package to perform DNS lookups and can utilize `assetfinder` to discover subdomains.

## Features

- Subdomain enumeration for a single domain using `net.LookupCNAME`.
- Subdomain enumeration for a single domain using `assetfinder` (optional flag).
- Subdomain enumeration for multiple domains using a file (each domain in a separate line).
- Identification of third-party CNAMEs.

## Installation

1. Install Go on your machine if you haven't already. You can download it from the official website: https://golang.org/dl/

2. Clone the repository:

   ```bash
   git clone https://github.com/your-username/ServiceScan.git
   cd ServiceScan
   ```

3. Build the executable:

   ```bash
   go build
   ```

## Usage

### Subdomain Enumeration for a Single Domain

To perform subdomain enumeration for a single domain, use the `-d` flag followed by the domain name. By default, the tool will use `net.LookupCNAME` to find subdomains.

```bash
./ServiceScan -d example.com
```

To use `assetfinder` for subdomain enumeration, include the `-assetfinder` flag:

```bash
./ServiceScan -d example.com -assetfinder
```

### Subdomain Enumeration for Multiple Domains

Create a text file (e.g., domains.txt) with each domain in a separate line. Then, use the `-f` flag followed by the file name to enumerate subdomains for multiple domains.

```bash
./ServiceScan -f domains.txt
```

## Output

The tool will display the subdomains along with their CNAME records. Third-party CNAMEs will be highlighted in red.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- The tool uses the `assetfinder` tool for subdomain enumeration. Check it out here: https://github.com/tomnomnom/assetfinder

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## Disclaimer

This tool is intended for legal and authorized use only. Always ensure that you have permission to scan and analyze the target domain's subdomains before using this tool. The authors are not responsible for any misuse or illegal activities.
```

Remember to replace `<your-username>` in the installation section with your GitHub username.

Feel free to customize the README to suit your tool's specific features and requirements. The README serves as the first point of contact for users and potential contributors, so make sure to include enough information to help them understand your tool and how to use it.
