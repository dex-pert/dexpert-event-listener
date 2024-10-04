package uniswapv2factoryaddress

import (
    "testing"
    "github.com/ethereum/go-ethereum/ethclient"
    "math/big"
)

func TestGetPair(t *testing.T) {
    cli, err := ethclient.Dial("https://mainnet-2.rpc.banelabs.org")
    if err != nil {
        t.Error(err)
        return
    }
    contarctAddress := "0x753df473702cB31BB81a93966e658e1AA4f10DD8"
    arg0, arg1 := "0xdE41591ED1f8ED1484aC2CD8ca0876428de60EfF", "0x68b55E582961968ef7758D8454D8A3e78c692e0B"
    resp, err := GetPair(big.NewInt(337089), contarctAddress, arg0, arg1, cli)
    if err != nil {
        t.Error(err)
        return
    }
    t.Log(resp)
}
