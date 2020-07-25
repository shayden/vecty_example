GOOS=js
GOARCH=wasm
export GOOS GOARCH

lib.wasm: main.go wasm_exec.js
	@go build -o lib.wasm main.go

wasm_exec.js:
	@cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" .

clean:
	@rm -f wasm_exec.js lib.wasm
