.PHONY: run
run:
	go run main.go

.PHONY: tokenfactory
tokenfactory:
	cd abi && abigen --abi tokenfactory/tokenfactory.abi --pkg tokenfactory --type tokenfactory --out ./tokenfactory/tokenfactory.go

.PHONY: universalrouter
universalrouter:
	cd abi && abigen --abi universalrouter/universalrouter.abi --pkg universalrouter --type universalrouter --out ./universalrouter/universalrouter.go

.PHONY: uniswapv2router
uniswapv2router:
	cd abi && abigen --abi uniswapv2router/uniswapv2router.abi --pkg uniswapv2router --type uniswapv2router --out ./uniswapv2router/uniswapv2router.go
