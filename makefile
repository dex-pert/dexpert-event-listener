.PHONY: run
run:
	go run main.go

.PHONY: tokenfactory
tokenfactory:
	cd abi && abigen --abi tokenfactory/tokenfactory.abi --pkg tokenfactory --type tokenfactory --out ./tokenfactory/tokenfactory.go
