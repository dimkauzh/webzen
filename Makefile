EXAMPLE_BUILD_TOOL = go run github.com/hajimehoshi/wasmserve@latest

.PHONY: example test

example:
	@echo
	@echo " ----------------------------------------------------"
	@echo "|              Running Example...                 |"
	@echo " ----------------------------------------------------"
	@echo
	
	$(EXAMPLE_BUILD_TOOL) ./example/

test:
	@echo
	@echo " ----------------------------------------------------"
	@echo "|              Building $(NAME)...                |"
	@echo " ----------------------------------------------------"
	@echo

	GOOS=js GOARCH=wasm go build -o test/main.wasm example/main.go
