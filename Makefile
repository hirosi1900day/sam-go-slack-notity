STACK_NAME := sam-app

build:
	sam build
.PHONY: build

deploy: 
	sam deploy \
		--stack-name ${STACK_NAME} \
		--s3-prefix ${STACK_NAME} \
		--s3-bucket sam-s3input-hisosi1900day00000 \
		--capabilities CAPABILITY_IAM \
		--region ap-northeast-1 \
		--no-confirm-changeset \
		--no-fail-on-empty-changeset \
		--no-progressbar
.PHONY: deploy	

install_linter:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.57.2
.PHONY: install_linter

lint:
	golangci-lint run
.PHONY: lint
