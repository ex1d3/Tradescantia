# Tradescantia
Tradescantia is the open-source Cryptomus webhook caller written in native Go.

## Features

- Lightweight
- Easy to use
- Useful for developers and QA`s.

## Demo

![](https://github.com/ex1d3/Tradescantia/blob/main/docs/demo.gif)

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

./tradescantia urlCallback=https://example.com/webhook invoiceUUID=123 
```
