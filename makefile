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

.PHONY: linuxrund
linuxrund:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./dexperteventlistener ./main.go && \
	nohup ./dexperteventlistener >> /var/log/dev_dexperteventlistener.log 2>&1 &


.PHONY: linuxstop
linuxstop:
	bash stop.sh

