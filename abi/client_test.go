package abi

import (
    "testing"
    "context"
    "math/big"
    "github.com/ethereum/go-ethereum/common"
)

func TestBlockNumber(t *testing.T) {
    mantaClient := MustNewClient("https://pacific-rpc.manta.network/http")
    t.Log("mataClient is :", mantaClient)

    block, err := mantaClient.Client.BlockByNumber(context.Background(), new(big.Int).SetUint64(3198801))
    if err != nil {
        t.Error(err)
    }
    t.Log(block)

    block1, err := mantaClient.BlockNumber(context.Background())
    if err != nil {
        t.Error(err)
    }
    t.Log(block1)

    block2, err := mantaClient.BlockByHash(context.Background(), common.HexToHash("0xb4a6fe3bf511b9bac8c4ad55a2daa88f49d6e63ee0821e14bd6c45facbd880b0"))
    if err != nil {
        t.Error(err)
    }
    t.Log(block2)
}
