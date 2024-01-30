LAMBDA_BUILD_DIR=dist/functions/
GITHUB_USER=pevin
# build: lambdas

build:
	for d in handlers/* ; do \
		LAMBDA_DIR=$${d#*/} ;\
		echo $(LAMBDA_BUILD_DIR)$${LAMBDA_DIR} ; \
		GOOS=linux GOARCH=amd64 go build -o $(LAMBDA_BUILD_DIR)$${LAMBDA_DIR} handlers/$${LAMBDA_DIR}/main.go ; \
	done

deploy: build deploy-only

deploy-only:
	bin/deploy.sh

clean:
	rm -rf dist

# make new handler
handler-new:
	bin/new_handler.sh $(name)
