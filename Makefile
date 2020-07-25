GOOS=js
GOARCH=wasm
export GOOS GOARCH

all: lib.wasm wasm_exec.js

lib.wasm: main.go
	@go build -o lib.wasm main.go

wasm_exec.js:
	@cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" .

clean:
	@rm wasm_exec.js lib.wasm
