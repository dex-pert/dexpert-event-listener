package tokenfactory

import "dexpert-event-listener/config"

type Proxy struct {
    chains  map[int]*config.Chain
    mapping map[int]*Client
}

func NewProxy(chains map[int]*config.Chain) (*Proxy, error) {
    chainsLen := len(chains)
    p := &Proxy{
        chains:  make(map[int]*config.Chain, chainsLen),
        mapping: make(map[int]*Client, chainsLen),
    }
    for k, v := range chains {
        p.chains[k] = chains[k]
        client := MustNewClient(v.URL)
        p.mapping[k] = client
    }
    return p, nil
}

func (p *Proxy) WithChainID(chainID int) *Client {
    return p.mapping[chainID]
}
