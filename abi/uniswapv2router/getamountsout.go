package uniswapv2router

import (
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "math/big"
    "github.com/pkg/errors"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func GetAmountsOutByAddressAndBlockNumber(uniswapV2RouterAddress string, blockNumber, amountIn *big.Int, client *ethclient.Client, path []common.Address) (outAmounts []*big.Int, err error) {
    uniswapv2router, err := NewUniswapv2router(common.HexToAddress(uniswapV2RouterAddress), client)
    if err != nil {
        return nil, errors.Wrap(err, "new uniswapv2 router error")
    }
    outAmounts, err = uniswapv2router.GetAmountsOut(&bind.CallOpts{BlockNumber: blockNumber}, amountIn, path)
    if err != nil {
        return nil, errors.Wrap(err, "uniswapv2router getAmountsOut error")
    }
    return outAmounts, nil
}
