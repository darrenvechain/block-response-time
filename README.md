# Block Response Time

A simple benchmarking tool to measure VeChain node block fetching response times, comparing sequential vs concurrent request performance.

## Overview

This tool fetches blocks from a VeChain node and measures:
- **Sequential fetching**: Blocks fetched one after another, reporting total time and average time per block
- **Concurrent fetching**: Blocks fetched in parallel using goroutines, reporting total time

## Installation

```bash
go build -o block-response-time
```

## Usage

```bash
./block-response-time [flags]
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-network-url` | `https://mainnet.vechain.org` | URL of the VeChain node to connect to |
| `-blocks` | `100` | Number of blocks to fetch |

### Examples

Benchmark against mainnet (default):
```bash
./block-response-time
```

Benchmark against testnet:
```bash
./block-response-time -network-url https://testnet.vechain.org
```

Fetch 500 blocks:
```bash
./block-response-time -blocks 500
```

## Sample Output

```
2024/01/15 10:30:00 Fetched 100 blocks sequentially in 5.2s (avg 52ms per block)
2024/01/15 10:30:01 Fetched 100 blocks concurrently in 850ms
```

## Requirements

- Go 1.25.1+
