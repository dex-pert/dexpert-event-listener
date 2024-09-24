package abi

import (
    "testing"
    "context"
    "math/big"
)

func TestBlockNumber(t *testing.T) {

    tests := []struct {
        name        string
        url         string
        blocknumber uint64
    }{
        {name: "ethereum-mainnet", url: "https://eth.drpc.org", blocknumber: 20711519},
        {name: "ethereum-sepolia", url: "https://rpc.sepolia.org/", blocknumber: 6641775},
        {name: "Conflux-eSpace", url: "https://evm.confluxrpc.com", blocknumber: 104087180},
        {name: "bitlayer-mainnet", url: "https://rpc.bitlayer.org", blocknumber: 4453885},
        {name: "manta-mainnet", url: "https://manta-pacific.drpc.org", blocknumber: 3198801},
        {name: "neo-x-mainnet", url: "https://mainnet-1.rpc.banelabs.org", blocknumber: 330867},
    }

    for _, v := range tests {
        t.Logf("-------------------------- %s --------------------------\n", v.name)
        client := MustNewClient(v.url)
        t.Log("mataClient is :", client)

        block1, err := client.BlockNumber(context.Background())
        if err != nil {
            t.Error(err)
        }
        t.Log("the newest block number: ", block1)

        block, err := client.Client.HeaderByNumber(context.Background(), new(big.Int).SetUint64(v.blocknumber))
        if err != nil {
            t.Error(err)
            t.Log("failed...")
            continue
        }
        t.Log("success...")
        t.Log("the block is: ", block)
        t.Log("the block unix time is: ", block.Time)
    }
}

func TestBlockTxHash(t *testing.T) {
    client := MustNewClient("https://evmmain-china.confluxrpc.com")
    t.Log("mataClient is :", client)

    block1, err := client.BlockNumber(context.Background())
    if err != nil {
        t.Error(err)
    }
    t.Log("the newest block number: ", block1)

    block, err := client.Client.HeaderByNumber(context.Background(), big.NewInt(104288565))
    if err != nil {
        t.Error(err)
        t.Log("failed...")
    }
    t.Log("success...")
    t.Log("the block is: ", block)
    t.Log("the block unix time is: ", block.Time)
}
