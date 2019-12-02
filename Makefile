2k20:
	GOOS=linux GOARCH=arm GOARM=5 go build -o 2k20 ./cmd/2k20/main.go

.PHONY: fmt
fmt:
	gofmt -s -w .

.PHONY: clean
clean:
	rm ./2k20

deploy: clean 2k20
	scp -q ./2k20 pi@pi-zero-wh.local:
