export PROD_PROJECT=prod_dexpert_event_listener

.PHONY: run
run:
	go run main.go

.PHONY: standardTokenFactory01
standardTokenFactory01:
	cd abi && abigen --abi standardTokenFactory01/standardTokenFactory01.abi --pkg standardTokenFactory01 --type standardTokenFactory01 --out ./standardTokenFactory01/standardTokenFactory01.go

.PHONY: dexpertUniversalRouter
dexpertUniversalRouter:
	cd abi && abigen --abi dexpertUniversalRouter/dexpertUniversalRouter.abi --pkg dexpertUniversalRouter --type dexpertUniversalRouter --out ./dexpertUniversalRouter/dexpertUniversalRouter.go

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

.PHONY: compose-up
compose-up:
	docker-compose -p $(PROD_PROJECT) -f docker-compose-prod.yaml build dexpert-event-listener && \
	docker-compose -p $(PROD_PROJECT) -f docker-compose-prod.yaml up -d

.PHONY: compose-down
compose-down:
	docker-compose -p $(PROD_PROJECT) -f docker-compose-prod.yaml down
