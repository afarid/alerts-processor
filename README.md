# Alert Processor

A Go-based tool to process and prioritize alert data from multiple monitoring systems.

## Features

- Parse alert JSON
- Filter by severity, service, time window
- Group related alerts
- Priority score calculation

## Usage

```bash
go run cmd/main.go --severity=critical --service=payment-processor --within=30
```

## Setup

```bash
go build -o alerts-processor cmd/main.go
```

## Tests

```bash
go test ./...
```

## CI/CD

GitHub Actions workflow runs on each PR and push.
