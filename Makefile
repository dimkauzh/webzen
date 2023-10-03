EXAMPLE_BUILD_TOOL = wasmserve

.PHONY: setup example build_test test_backend test1

setup:
	@echo
	@echo " ----------------------------------------------------"
	@echo "|              Setting up...                       |"
	@echo " ----------------------------------------------------"
	@echo

	go install github.com/hajimehoshi/wasmserve@latest


example:
	@echo
	@echo " ----------------------------------------------------"
	@echo "|              Running Example...                 |"
	@echo " ----------------------------------------------------"
	@echo
	
	$(EXAMPLE_BUILD_TOOL) ./example/

build_test:
	@echo
	@echo " ----------------------------------------------------"
	@echo "|              Building $(NAME)...                |"
	@echo " ----------------------------------------------------"
	@echo

	GOOS=js GOARCH=wasm go build -o build/main.wasm tests/backend/backend.go
	GOOS=js GOARCH=wasm go build -o build/main.wasm tests/test1/test1.go

test_backend:
	@echo
	@echo " ----------------------------------------------------"
	@echo "|              Building $(NAME)...                |"
	@echo " ----------------------------------------------------"
	@echo

	$(EXAMPLE_BUILD_TOOL) ./tests/backend/backend.go

test1:
	@echo
	@echo " ----------------------------------------------------"
	@echo "|              Running Test 1...                   |"
	@echo " ----------------------------------------------------"
	@echo

	$(EXAMPLE_BUILD_TOOL) ./tests/test1/test1.go
