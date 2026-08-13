WebScan

WebScan adalah tool web assessment berbasis Go untuk melakukan analisis informasi publik sebuah website. Tool ini mengumpulkan informasi seperti HTTP response, TLS, security header, teknologi yang digunakan, endpoint, metadata HTML, serta membuat laporan hasil scan.

Fitur

- HTTP status & response analysis
- HTTPS & TLS information
- DNS lookup
- Redirect chain detection
- Security header analysis
  - Content-Security-Policy
  - Strict-Transport-Security
  - X-Content-Type-Options
  - X-Frame-Options
  - Referrer-Policy
  - Permissions-Policy
- Cookie analysis
- Technology fingerprinting
- HTML metadata extraction
- JavaScript & asset analysis
- Endpoint inventory
- Risk scoring
- Security findings
- JSON output
- HTML report
- Scan history
- Compare hasil scan
- Multi target scan

---

Installation

Requirements

- Go 1.20+
- Linux / macOS / Termux

Build

Clone repository:
```bash
git clone https://github.com/Elfajree07/webscan.git
cd webscan
```
Build:
```bash
go build -o webscan .
```
Test instalasi:
```bash
./webscan --version
```
Output:

WebScan v2.x.x

---

Basic Usage

Scan website
```bash
./webscan https://example.com
```
Output akan menampilkan:

- Target
- Status HTTP
- Server
- HTTPS
- TLS
- DNS
- Security header
- Score

---

JSON Output

Untuk mendapatkan hasil dalam format JSON:
```bash
./webscan --json https://example.com
```
Simpan hasil:
```bash
./webscan --json https://example.com > scan.json
```
---

HTML Report

Membuat laporan HTML:
```bash
./webscan --report https://example.com
```
Default output:

webscan-report.html

Custom nama report:
```bash
./webscan --report --output report.html https://example.com
```
Report berisi:

- Risk summary
- Security score
- Findings
- Technologies
- Endpoints
- Metadata
- TLS information
- History timeline
- Compare result

---

Scan Profile

Quick Scan

Scan dasar:
```bash
./webscan --profile quick https://example.com
```
Full Scan

Scan lengkap:
```bash
./webscan --profile full https://example.com
```
Default profile:

full

---

Compare Scan

Membandingkan dua hasil scan:

Scan pertama:
```bash
./webscan --json https://example.com > old.json
```
Scan kedua:
```bash
./webscan --json https://example.com > new.json
```
Compare:
```bash
./webscan --compare old.json new.json
```
Contoh output:

WebScan Compare
--------------------------------
Score: 50 -> 75

Changes:
[+] Security score meningkat

---

Multiple Target Scan

Buat file target:
```bash
nano targets.txt
```
Isi:

https://example.com
https://example.org / ini hanya contohnya!

Jalankan:
```bash
./webscan --list targets.txt
```
Simpan summary:
```bash
./webscan --list targets.txt --summary summary.json
```
---

Configuration

Buat konfigurasi default:
```bash
./webscan --init-config
```
File:

webscan.json

Gunakan config:
```bash
./webscan --config webscan.json https://example.com
```
---

Command Reference

Command| Fungsi
"--version"| Menampilkan versi
"--json"| Output JSON
"--report"| Membuat HTML report
"--output"| Nama file report
"--profile"| Mode quick/full
"--compare"| Membandingkan scan
"--list"| Scan banyak target
"--summary"| Simpan summary
"--config"| Menggunakan config
"--init-config"| Membuat config
"--timeout"| Mengatur timeout
"--workers"| Jumlah worker

---

Project Structure

webscan/
├── main.go
├── scanner/
│   ├── technology detection
│   ├── asset analysis
│   └── endpoint extraction
├── report.go
├── findings.go
├── risk.go
├── score.go
├── history.go
├── compare.go
├── config.go
└── go.mod

---

Example Workflow

Scan website:
```bash
./webscan --json https://example.com > first.json
```
Buat report:
```bash
./webscan --report https://example.com
```
Beberapa waktu kemudian scan ulang:
```bash
./webscan --json https://example.com > second.json
```
Bandingkan:
```bash
./webscan --compare first.json second.json
```
---

Development

Run test:
```bash
go test ./...
```
Format code:
```bash
gofmt -w *.go scanner/*.go
```
Build:
```bash
go build -o webscan .
```
---

Disclaimer

WebScan dibuat untuk tujuan edukasi, audit keamanan yang memiliki izin, dan monitoring aset milik sendiri.

Gunakan tool secara bertanggung jawab dan selalu pastikan memiliki izin sebelum melakukan scanning terhadap sistem pihak lain.

---

License

MIT License
