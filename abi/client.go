package abi

import "github.com/ethereum/go-ethereum/ethclient"

type Client struct {
    *ethclient.Client
    ethNodeUrl string
}

func MustNewClient(url string) *Client {
    cli, err := ethclient.Dial(url)
    if err != nil {
        panic(err)
    }
    return &Client{ethNodeUrl: url, Client: cli}
}
