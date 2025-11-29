# 🧟 API-Gravekeeper: The DevSecOps Tool for Hunting Zombie and Shadow APIs

[![Build Status](https://img.shields.io/badge/build-passing-brightgreen)](https://github.com/CyberError911/api-gravekeeper)
[![Go Version](https://img.shields.io/badge/go-1.20+-blue)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)
[![GitHub Stars](https://img.shields.io/github/stars/CyberError911/api-gravekeeper?style=social)](https://github.com/CyberError911/api-gravekeeper)

## The Problem

Your API surface is a graveyard. Traditional security scanners (SAST/DAST) don't catch two critical issues:

- **Zombie APIs**: Routes defined in your code but never accessed—dead weight that accumulates technical debt and increases your attack surface
- **Shadow APIs**: Endpoints users are hitting that your code doesn't account for—undocumented, unmonitored, a security blind spot

These anomalies represent real risk: forgotten credentials in dead code, unpatched vulnerabilities in unused endpoints, and shadow services operating outside your security model.

**API-Gravekeeper** solves this by comparing what's in your code against what's actually being used, identifying both problems in minutes.

## Key Features

🔍 **Accountability** — Git Blame Integration  
Every zombie route shows who wrote it and when. Bridge the gap between security findings and engineering accountability.

✅ **Accuracy** — Route Normalization  
Dynamic paths like `/users/123` and `/users/456` are intelligently normalized to `/users/:id`. No false positives from numeric IDs or UUIDs.

🛡️ **Security Insight** — Dual Detection  
Finds both unused code (Zombies) and undocumented endpoints (Shadows) in a single scan.

## Quick Start

### Installation

```bash
git clone https://github.com/CyberError911/api-gravekeeper.git
cd api-gravekeeper
go build -o api-gravekeeper
```

### Usage

```bash
./api-gravekeeper scan --code-dir ./my-api-repo/src --log-file ./nginx/access.log
```

### Flags

- `--code-dir`, `-c` (required): Root directory to scan for API definitions (e.g., `./src`)
- `--log-file`, `-l` (required): Path to access log file (e.g., `./logs/access.log`)

## Sample Output

```
Zombies (defined but never accessed):
  /admin/dashboard (./src/app.py:21) - alice@company.com (2024-08-15 09:32:17 +0000)
  /settings (./src/app.py:25) - bob@company.com (2024-09-10 14:05:42 +0000)
  /api/v1/legacy/export (./src/handlers.py:156) - charlie@company.com (2023-12-01 11:20:00 +0000)

Shadows (accessed but not defined):
  /api/v2/notifications
  /health
  /metrics
```

## How It Works

1. **Code Scanner**: Walks your codebase and extracts all defined API routes (currently supports Python `@app.route()`)
2. **Log Parser**: Reads access logs and extracts unique accessed paths
3. **Route Normalizer**: Converts dynamic segments (IDs, UUIDs) to standard placeholders for accurate matching
4. **Comparator**: Identifies zombies (defined but unused) and shadows (accessed but undefined)
5. **Git Blame**: For each zombie, provides authorship and commit date for accountability

## Supported Frameworks

**Current:**
- Flask (Python, via `@app.route()`)

**Planned:**
- FastAPI
- Django
- Express.js
- Go Chi/Gin

## Log Format Support

- **Nginx**: Common log format and combined log format
- **Apache**: Standard and combined formats

## Requirements

- Go 1.20 or later
- Git (for blame integration)
- Access to codebase and access logs

## Contributing

Contributions are welcome! Areas of focus:

- **New Language Support**: Add scanners for Django, FastAPI, Express.js, etc.
- **Log Parser Enhancements**: Support for additional log formats
- **Testing**: Unit tests and integration tests for edge cases

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## License

MIT License — see [LICENSE](LICENSE) for details.

---

**Maintained by**: [CyberError911](https://github.com/CyberError911)


