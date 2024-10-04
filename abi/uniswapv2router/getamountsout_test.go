package uniswapv2router

import (
    "testing"
    "github.com/ethereum/go-ethereum/ethclient"
    "math/big"
    "github.com/ethereum/go-ethereum/common"
)

func TestGetAmountsOutByAddressAndBlockNumber(t *testing.T) {
    cli, err := ethclient.Dial("https://mainnet-1.rpc.banelabs.org")
    if err != nil {
        t.Error(err)
        return
    }

    outAmount, err := GetAmountsOutByAddressAndBlockNumber("0x82b56Dd9c7FD5A977255BA51B96c3D97fa1Af9A9",
        big.NewInt(437883), big.NewInt(100000000000000), cli, []common.Address{common.HexToAddress("0xdE41591ED1f8ED1484aC2CD8ca0876428de60EfF"),
            common.HexToAddress("0x68b55E582961968ef7758D8454D8A3e78c692e0B")})
    if err != nil {
        t.Error(err)
        return
    }
    t.Log("outAmount: ", outAmount)
}

func TestGetAmountsOutNEOXByAddressAndBlockNumber(t *testing.T) {
    cli, err := ethclient.Dial("https://mainnet-2.rpc.banelabs.org")
    if err != nil {
        t.Error(err)
        return
    }

    ammountIn, ok := new(big.Int).SetString("417701490067272740000", 10)
    if !ok {
        t.Error("failed to convert big int")
        return
    }
    outAmount, err := GetAmountsOutByAddressAndBlockNumber("0x82b56Dd9c7FD5A977255BA51B96c3D97fa1Af9A9",
        big.NewInt(337089), ammountIn, cli, []common.Address{common.HexToAddress("0x0ef63483c492864d0da238d6f20ffd0a0fdc437a"), common.HexToAddress("0xdE41591ED1f8ED1484aC2CD8ca0876428de60EfF"),
            common.HexToAddress("0x68b55E582961968ef7758D8454D8A3e78c692e0B")})
    if err != nil {
        t.Error(err)
        return
    }
    t.Log("outAmount: ", outAmount)
}
