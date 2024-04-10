build:
	sam build
.PHONY: build

deploy: 
	sam deploy --guided
.PHONY: deploy	

install_linter:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.57.2
.PHONY: install_linter

lint:
	golangci-lint run
.PHONY: lint
