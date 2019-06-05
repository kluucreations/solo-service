SWAGGER_VERSION = v0.18.0

.PHONY: setup gen

setup:
	curl -L https://github.com/go-swagger/go-swagger/releases/download/${SWAGGER_VERSION}/swagger_linux_amd64 > ./go-swagger
	chmod 755 ./go-swagger

clean:
	rm -rf ./gen

gen: clean
	mkdir ./gen
	./go-swagger generate server -t gen -f ./swagger/api.json --exclude-main -A solo-service
	go get -u -v -f ./gen/...