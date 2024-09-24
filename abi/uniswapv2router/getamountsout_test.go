package uniswapv2router

import (
    "testing"
    "github.com/ethereum/go-ethereum/ethclient"
    "math/big"
    "github.com/ethereum/go-ethereum/common"
)

func TestGetAmountsOutByAddressAndBlockNumber(t *testing.T) {
    cli, err := ethclient.Dial("https://evmmain-china.confluxrpc.com")
    if err != nil {
        t.Error(err)
        return
    }

    outAmount, err := GetAmountsOutByAddressAndBlockNumber("0x62b0873055Bf896DD869e172119871ac24aEA305",
        big.NewInt(105367840), big.NewInt(100000000000000), cli, []common.Address{common.HexToAddress("0x14b2D3bC65e74DAE1030EAFd8ac30c533c976A9b"),
            common.HexToAddress("0xfe97E85d13ABD9c1c33384E796F10B73905637cE")})
    if err != nil {
        t.Error(err)
        return
    }
    t.Log("outAmount: ", outAmount)
}
