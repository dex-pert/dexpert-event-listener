package standardTokenFactory01

import (
    "testing"
    "github.com/ethereum/go-ethereum/ethclient"
    "math/big"
)

func TestGetFee(t *testing.T) {
    cli, err := ethclient.Dial("https://evm.confluxrpc.com")
    if err != nil {
        t.Error(err)
        return
    }

    fee1, err := GetFees("0x9998Bc4B6e148FC4633b274E2CC040B19638F554", big.NewInt(1), big.NewInt(104091299), cli)
    if err != nil {
        t.Error(err)
        return
    }
    t.Log("fee1: ", fee1)

    fee2, err := GetFees("0x9998Bc4B6e148FC4633b274E2CC040B19638F554", big.NewInt(2), big.NewInt(104091310), cli)
    if err != nil {
        t.Error(err)
        return
    }
    t.Log("fee2: ", fee2)
}
