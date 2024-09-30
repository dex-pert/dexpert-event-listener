package uniswapv2factoryaddress

import (
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/pkg/errors"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "math/big"
)

func GetPair(blockNumber *big.Int, contractAddress, arg0, arg1 string, client *ethclient.Client) (string, error) {
    uniswapv2Factoryaddress, err := NewUniswapv2factoryaddress(common.HexToAddress(contractAddress), client)
    if err != nil {
        return "", errors.Wrap(err, "failed to new token factory")
    }
    address, err := uniswapv2Factoryaddress.GetPair(&bind.CallOpts{
        BlockNumber: blockNumber,
    }, common.HexToAddress(arg0), common.HexToAddress(arg1))
    if err != nil {
        return "", errors.Wrap(err, "failed to get pair")
    }
    return address.String(), nil
}
