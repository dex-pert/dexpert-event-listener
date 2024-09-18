package erc20

import (
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/pkg/errors"
    "github.com/ethereum/go-ethereum/ethclient"
)

func GetSymbolNameDecimalByAddress(tokenAddress common.Address, client *ethclient.Client) (symbol string, name string, decimals int32, err error) {
    erc20, err := NewERC20(tokenAddress, client)
    if err != nil {
        return "", "", 0, errors.Wrap(err, "failed to new erc20")
    }

    opt := bind.CallOpts{
        Pending:     false,
        From:        tokenAddress,
        BlockNumber: nil,
        BlockHash:   common.Hash{},
        Context:     nil,
    }

    symbol, err = erc20.Symbol(&opt)
    if err != nil {
        return "", "", 0, errors.Wrap(err, "failed to get erc20 symbol")
    }
    name, err = erc20.Name(&opt)
    if err != nil {
        return "", "", 0, errors.Wrap(err, "failed to get erc20 name")
    }
    _decimals, err := erc20.Decimals(&opt)
    if err != nil {
        return "", "", 0, errors.Wrap(err, "failed to get erc20 decimals")
    }
    decimals = int32(_decimals)
    return symbol, name, decimals, nil
}
