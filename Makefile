.PHONY: run 
run:
	@echo 'Running application..'
	@go run ./cmd/api -port=3000 -env=production