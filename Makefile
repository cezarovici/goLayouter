.PHONY: test
test: ## Run tests with check race and coverage
	@go test -failfast -count=1 ./... -json -cover -race | tparse -smallscreen

.PHONY: benchmark
benchmark:
	@go test  ./... -bench=.