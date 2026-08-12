# WebScan

A lightweight Go-based website assessment CLI for authorized security testing.

## Features

- HTTP/HTTPS status
- Response time
- DNS resolution
- TLS version
- TLS certificate information
- Security header presence
- Redirect chain
- robots.txt check
- sitemap.xml check
- JSON output
- HTML report
- Custom report filename
- Configurable timeout

## Install

```bash
pkg update && pkg upgrade -y
pkg install git -y
git clone https://github.com/Elfajree07/webscan.git
cd webscan
```

## Build

```bash
go build -o webscan .
```

Usage

```bash
./webscan https://example.com
```

JSON:

```bash
./webscan --json https://example.com
```

HTML report:

```bash
./webscan --report https://example.com
```

Custom report:

```bash
./webscan --report --output example-report.html https://example.com
```

Version:

```bash
./webscan --version
```

## NOTES

WebScan is intended for passive/low-impact assessment of
systems you own or have explicit permission to test.

The security-header score indicates only whether selected
headers were observed. It is not a complete measure of website
security.

## License

MIT


## Cek semuanya

```bash
gofmt -w *.go
go test ./...
go build -o webscan .
./webscan --version
```
