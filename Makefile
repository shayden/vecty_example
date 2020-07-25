lib.wasm: main.go wasm_exec.js
	@GOOS=js GOARCH=wasm go build -o lib.wasm main.go

wasm_exec.js:
	@cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" .

serve:
	@go run server.go
clean:
	@rm -f wasm_exec.js lib.wasm
