EXAMPLE_BUILD_TOOL = go run github.com/hajimehoshi/wasmserve@latest

.PHONY: run_example

run_example:
	@echo
	@echo " ----------------------------------------------------"
	@echo "|              Running Example...                 |"
	@echo " ----------------------------------------------------"
	@echo
	
	$(EXAMPLE_BUILD_TOOL) ./example/
