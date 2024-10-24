# Crypto Mine CLI

Crypto Mine CLI is an application that extracts the data of various cryptocurrencies from the [Coin Market Cap](https://coinmarketcap.com/) and persists it in a CSV file in Donwloads Folder

## Requirements

- Golang 16+

## How to run

#### Download the project dependencies

```go
go mod tidy
go mod vendor
```

#### Execute the project

```go
go run cmd/main.go
```

The following message should appear in the terminal:

![run output](/docs/images/run-without-save.png)

By adding the ```--save``` you can save the results in a CSV file in your system's Downloads folder (You can also use ```-s```)

```go
go run cmd/main.go --save
```

The following message should appear in the terminal with a additional log:

![run output with save](/docs/images/run-with-save.png)