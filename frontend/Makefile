.PHONY: build clean deploy

vendor:
	npm install

validate:
	@[ "${ENV}" ] || ( echo "Error: Please provide target environment. ex: 'make deploy ENV=dev'"; exit 1 )

fetchEnvVars:
	$(eval API_BASE_URL := $(shell aws ssm get-parameter --name "/crispurl/${ENV}/API_BASE_URL" --with-decryption --query Parameter.Value --output text))

build:
	npm install &&\
	NEXT_PUBLIC_API_BASE_URL=${API_BASE_URL} npm run build &&\
	npm run export

s3deploy:
	aws s3 sync ./out s3://crispurl-landingpage-${ENV}/

clean:
	rm -rf ./out

deploy: validate fetchEnvVars build s3deploy clean
