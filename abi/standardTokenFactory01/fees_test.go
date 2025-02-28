package standardTokenFactory01

import (
    "testing"
    "github.com/ethereum/go-ethereum/ethclient"
    "math/big"
)

func TestGetFee(t *testing.T) {
    cli, err := ethclient.Dial("https://evmmain-global.confluxrpc.com")
    if err != nil {
        t.Error(err)
        return
    }

    fee1, err := GetFees("0xF2F0cC2BAB14Abbdf3b2AFcAD345dE08225EBaf5", big.NewInt(2), big.NewInt(105367315), cli)
    if err != nil {
        t.Error(err)
        return
    }
    t.Log("fee1: ", fee1)
}
