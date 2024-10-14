# Flipside Go SDK

This Go SDK provides a client for interacting with the Flipside Crypto API. It allows you to perform various operations related to Solana blockchain data analysis.

## Features

- Query first buyers of a specific token
- Get transfers between multiple addresses
- Retrieve first swap data for multiple addresses
- Execute custom SQL queries against Flipside's Solana dataset

## Installation

To use this SDK in your Go project, you can install it using:

```bash
go get github.com/franco-bianco/flipside-client
```

## Usage

### Initializing the Client

```go
import (
    "github.com/sirupsen/logrus"
    flipside "github.com/franco-bianco/flipside-client"
)

apiKey := "your-flipside-api-key"
log := logrus.New()
client := flipside.NewClient(apiKey, log)
```

### Getting First Buyers of a Token

```go
tokenAddress := "your-token-address"
limit := 10
swapData, err := client.GetFirstBuyers(tokenAddress, limit)
if err != nil {
    // Handle error
}
// Process swapData
```

### Getting Transfers Between Addresses

```go
addresses := []string{"address1", "address2", "address3"}
limit := 100
interactionsMap, err := client.GetTransfersBetweenAddresses(addresses, limit)
if err != nil {
    // Handle error
}
// Process interactionsMap
```

### Getting First Swaps for Multiple Addresses

```go
addresses := []string{"address1", "address2", "address3"}
firstSwaps, err := client.GetFirstSwaps(addresses)
if err != nil {
    // Handle error
}
// Process firstSwaps
```

## Data Structures

The SDK provides several data structures to represent the query results:

- `SwapData`: Represents data for a swap transaction
- `TransferData`: Represents data for a token transfer transaction
- `AddressInteractions`: Represents interactions for a single address
- `FirstSwapData`: Represents the first swap data for an address
