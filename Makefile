build:
	go build -o ./bin/gotls -ldflags "-s -w" ./cmd/

buildwin:
	GOOS=windows go build -o ./bin/gotls.exe -ldflags "-s -w" ./cmd/

run:
	go run ./cmd/
