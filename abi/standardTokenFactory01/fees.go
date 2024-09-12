package standardTokenFactory01

import (
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/common"
    "github.com/pkg/errors"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "math/big"
)

func GetFees(contractAddress string, level, blockNumber *big.Int, client *ethclient.Client) (*big.Int, error) {
    tokenFactory, err := NewStandardTokenFactory01(common.HexToAddress(contractAddress), client)
    if err != nil {
        return nil, errors.Wrap(err, "failed to new token factory")
    }
    fees, err := tokenFactory.Fees(&bind.CallOpts{BlockNumber: blockNumber}, level)
    if err != nil {
        return nil, errors.Wrap(err, "failed to get tokenFactory fees")
    }
    return fees, nil
}
