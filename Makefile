build:
	go build -o ./bin/gochk -ldflags "-s -w" main.go

buildwin:
	GOOS=windows go build -o ./bin/gochk.exe -ldflags "-s -w" main.go

run:
	go run main.go
