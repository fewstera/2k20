2k20:
	CC=/Users/aidanf/Development/personal/raspbian-sdk/prebuilt/bin/cglang \
	CGO_CFLAGS="--sysroot=/Users/aidanf/Development/personal/raspbian-sdk/sysroot" \
	CGO_LDFLAGS="--sysroot=/Users/aidanf/Development/personal/raspbian-sdk/sysroot -L/Users/aidanf/Development/personal/raspbian-sdk/sysroot/usr/lib/gcc/arm-linux-gnueabihf/8" \
	CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=6 go build -o 2k20 ./cmd/2k20/main.go

.PHONY: fmt
fmt:
	gofmt -s -w .

.PHONY: clean
clean:
	rm -f ./2k20

deploy: clean 2k20
	scp -q ./2k20 pi@pi-zero-wh.local:

run:
	go run ./cmd/2k20/main_darwin.go
