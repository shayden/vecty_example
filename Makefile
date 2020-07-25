www/lib.wasm: main.go www/wasm_exec.js
	@GOOS=js GOARCH=wasm go build -o www/lib.wasm main.go

www/wasm_exec.js:
	@cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" ./www

serve:
	@go run server.go

clean:
	@rm -f www/wasm_exec.js www/lib.wasm
