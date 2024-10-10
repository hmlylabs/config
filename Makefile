ARTIFACT_NAME=household-service

build:
	@go build -o build/${ARTIFACT_NAME}/${ARTIFACT_NAME} cmd/${ARTIFACT_NAME}.go

run:
	@go run cmd/main.go

test:
	@go clean -testcache && go test ./ -v

docker-build:
	@docker build . -t europe-west2-docker.pkg.dev/hmly-tech/api/${ARTIFACT_NAME}

docker-run:
	@docker run europe-west2-docker.pkg.dev/hmly-tech/api/${ARTIFACT_NAME}
