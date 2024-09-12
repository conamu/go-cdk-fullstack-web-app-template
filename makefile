.SILENT:
.PHONY:
build:
	echo "**Building binaries**"
	arch=arm64 ./scripts/build-container-binaries.sh
	echo "**Building Lambdas**"
	arch=arm64 ./scripts/build-lambda-payloads.sh

bootstrap:
	cdk bootstrap

destroy:
	cdk destroy

deploy: build
	./deploy-all.sh