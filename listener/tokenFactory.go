package listener

import (
    "context"
    "log/slog"
    "math/big"
    el "github.com/x1rh/event-listener"

    "github.com/ethereum/go-ethereum/common"
    "github.com/pkg/errors"
    "github.com/ethereum/go-ethereum/ethclient"
)

type TokenFactoryEventListener struct {
    EL *el.EventListener
}

func NewTokenFactoryEventListener(ltCtx *LTContext) (IListener, error) {
    c := el.ChainConfig{
        ChainId:   ltCtx.Chain.ChainId,
        ChainName: ltCtx.Chain.ChainName,
        URL:       ltCtx.Chain.URL,
    }

    client, err := ethclient.Dial(c.URL)
    if err != nil {
        panic(err)
    }

    tokenFactorAddress := ltCtx.TokenFactorAddress
    abiStr := ltCtx.ABIStr
    blockNumber := big.NewInt(ltCtx.BlockNumber)
    step := big.NewInt(ltCtx.Step)
    tokenFactory, err := el.NewContract(tokenFactorAddress, abiStr, blockNumber, step)
    if err != nil {
        panic(err)
    }
    tokenFactory.SetLogHandler(logHandler(ltCtx, tokenFactory))
    _el, err := el.New(
        c,
        el.WithClient(client),
        el.WithContract(*tokenFactory),
    )
    if err != nil {
        return nil, errors.Wrap(err, "fail to new an EventListener object")
    }
    return &TokenFactoryEventListener{EL: _el}, nil
}

func (t *TokenFactoryEventListener) Start() {
    t.EL.Start()
}

// LogTokenCreated signature: TokenCreated(address,address,uint8,uint96,uint256)
type LogTokenCreated struct {
    Owner        common.Address // token owner
    Token        common.Address // token address
    TokenType    uint8
    TokenVersion *big.Int
    Level        *big.Int
}

func logHandler(ltCtx *LTContext, c *el.Contract) el.LogHandleFunc {
    return func(ctx context.Context, event *el.Event) error {
        switch event.Name {
        case "TokenCreated":
            var l LogTokenCreated
            if err := c.Abi.UnpackIntoInterface(&l, event.Name, event.Data); err != nil {
                return errors.Wrap(err, "fail to unpack log")
            }
            // handle indexed topic
            l.Owner = el.HashToAddress(event.IndexedParams[0])
            l.Token = el.HashToAddress(event.IndexedParams[1])
            slog.Info("TokenCreated event", slog.Any("event", l))

        default:
            // do nothing
        }
        return nil
    }
}
