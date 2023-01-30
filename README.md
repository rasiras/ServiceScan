# ServiceScan

A tool for finding subdomains hosted on third party servers or uses a third party service.

**Description**

This tool takes a list of subdomains as input and checks the CNAME (Canonical Name) record for each one. If the CNAME record does not contain the main domain name and is not an IP address, the subdomain is considered to be hosted on a third party server and is highlighted in red.

**Install**

```
go install github.com/rasiras/ServiceScan/SScan@latest

```

**Usage**

To use the tool, run the following command:

```
go build subdomain-finder.go
./SScan <filename>

```
where ```<filename>``` is the path to a file containing the list of subdomains, one per line.

Example

Input:

```

sub1.example.com
sub2.example.com
sub3.example.com

```
  
Output 
```

sub1.example.com [example.com]
sub2.example.com [ezample.thirdparty.com]
sub3.example.com [1.2.3.4]
```
