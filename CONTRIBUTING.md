# Contributing to API-Gravekeeper

Thanks for your interest in contributing! Here's how you can help:

## Getting Started

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/your-feature`
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## Areas We Need Help

### New Language Support
Add scanners for additional frameworks:
- **FastAPI** (Python)
- **Django** (Python)
- **Express.js** (Node.js)
- **Spring Boot** (Java)
- **Go Chi/Gin**

See `internal/scanner/scanner.go` for the interface pattern.

### Log Format Support
Extend `internal/logs/parser.go` to support:
- CloudFront logs
- ALB/NLB logs
- Custom log formats

### Testing
- Unit tests for scanner regex patterns
- Integration tests with real log files
- Edge case testing for normalization logic

## Code Standards

- Follow Go conventions (gofmt, golint)
- Add comments for exported functions
- Keep functions focused and testable
- Write clear commit messages

## Testing Your Changes

```bash
go test ./...
go build -o api-gravekeeper
./api-gravekeeper scan -c ./test/fixtures -l ./test/logs/test.log
```

## Reporting Issues

Please include:
- Go version (`go version`)
- OS and architecture
- Steps to reproduce
- Sample code or logs (sanitized)

## License

All contributions are licensed under the MIT License.
