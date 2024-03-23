# Tradescantia
Tradescantia is the open-source Cryptomus webhook caller written in native Go.

## Features

- Lightweight
- Easy to use
- 0 dependencies

## Install & usage

```
# Clone this repository
git clone https://github.com/ex1d3/Tradescantia
cd tradescantia

# Build project
go build

# Initialize and edit config
./tradescantia && vi config.json

# Call test webhook

./tradescantia --invoiceUUID=123 --currency=USDT --network=tron --status=paid
```
