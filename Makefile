include .env
export

.PHONY: test
test:
	TF_ACC=1 go test ./...
	
.PHONY: generate-client
generate-client:
	rm -rf gen && mkdir -p gen/
	rm -rf api/v1 && mkdir -p api/v1/spec
	cp ../api/*.yml api/v1/spec/
	docker run --rm \
		-v "${PWD}/api/v1/spec:/local/api:ro" \
		-v "${PWD}/gen:/local/gen" \
		-u "$(shell id -u):$(shell id -g)" \
		openapitools/openapi-generator-cli:v7.3.0 generate -g go --additional-properties=enumClassPrefix=true \
		-i /local/api/udoma-openapi.yml \
		-o /local/gen \
		--additional-properties=packageName=v1
	rm -rf gen/go.*  gen/git_push.sh gen/test gen/docs
	cp -R gen/* api/v1
	rm -rf gen api/v1/spec
	go fmt ./... && go vet ./...


.PHONY: sanitize
sanitize:
	go fmt ./...
	go vet ./...

.PHONY: generate-docs
generate-docs:
	tfplugindocs generate

.PHONY: build
build:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -installsuffix cgo -ldflags "-X main.version=${VERSION} -X main.builddate=`date -u +.%Y%m%d.%H%M%S` -s -w" -o terraform-provider-udoma main.go
