EXAMPLE_BUILD_TOOL = wasmserve
VERSION = null

.PHONY: setup example release

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

release:
	@echo
	@echo " ----------------------------------------------------"
	@echo "|              Publishing $(NAME)...              |"
	@echo " ----------------------------------------------------"
	@echo

	GOPROXY=proxy.golang.org go list -m github.com/dimkauzh/webzen@$(VERSION)
