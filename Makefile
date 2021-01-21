all:
	@echo "'make run_no_2' to running no_2"
	@echo "'make run_no_4' to running no_4"
run_no_2:
	@echo "Running no_2"
	@echo "Download dependencies"
	@go mod download
	@echo "Add vendor"
	@go mod vendor
	@echo "Running app"
	@go run ./no_2/main/*.go

run_no_4:
	@echo "Running no_4"
	@go run ./no_4/no_4.go