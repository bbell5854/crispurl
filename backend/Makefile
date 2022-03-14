.PHONY: build clean deploy

vendor:
	go mod vendor

validate:
	@[ "${ENV}" ] || ( echo "Error: Please provide target environment. ex: 'make deploy ENV=dev'"; exit 1 )

build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/get-redirect/main ./functions/get-redirect/*.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/create-redirect/main ./functions/create-redirect/*.go

slsdeploy:
	sls deploy --verbose --conceal --stage $(ENV)

clean:
	rm -rf ./bin ./vendor

deploy: validate build slsdeploy clean
