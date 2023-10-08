home-directory=/home/hello/cosmos-checkers

mock-expected-keepers:
	mockgen -source=x/cosmoscheckers/types/expected_keepers.go \
		-package testutil \
		-destination=x/cosmoscheckers/testutil/expected_keepers_mocks.go 

install-protoc-gen-ts:
	mkdir -p scripts/protoc
	cd scripts && npm install
	curl -L https://github.com/protocolbuffers/protobuf/releases/download/v21.5/protoc-21.5-linux-x86_64.zip -o scripts/protoc/protoc.zip
	cd scripts/protoc && unzip -o protoc.zip
	rm scripts/protoc/protoc.zip
	ln -s $(home-directory)/scripts/protoc/bin/protoc /usr/local/bin/protoc

cosmos-version = v0.45.4

download-cosmos-proto:
	mkdir -p proto/cosmos/base/query/v1beta1
	curl https://raw.githubusercontent.com/cosmos/cosmos-sdk/${cosmos-version}/proto/cosmos/base/query/v1beta1/pagination.proto -o proto/cosmos/base/query/v1beta1/pagination.proto
	mkdir -p proto/google/api
	curl https://raw.githubusercontent.com/cosmos/cosmos-sdk/${cosmos-version}/third_party/proto/google/api/annotations.proto -o proto/google/api/annotations.proto
	curl https://raw.githubusercontent.com/cosmos/cosmos-sdk/${cosmos-version}/third_party/proto/google/api/http.proto -o proto/google/api/http.proto
	mkdir -p proto/gogoproto
	curl https://raw.githubusercontent.com/cosmos/cosmos-sdk/${cosmos-version}/third_party/proto/gogoproto/gogo.proto -o proto/gogoproto/gogo.proto

gen-protoc-ts: 
	mkdir -p ./client/src/types/generated/
	ls proto/cosmoscheckers | xargs -I {} protoc \
		--plugin="./scripts/node_modules/.bin/protoc-gen-ts_proto" \
		--ts_proto_out="./client/src/types/generated" \
		--proto_path="./proto" \
		--ts_proto_opt="esModuleInterop=true,forceLong=long,useOptionals=messages" \
		cosmoscheckers/{}
	
install-and-gen-protoc-ts: download-cosmos-proto install-protoc-gen-ts gen-protoc-ts

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ./build/cosmos-checkersd-linux-amd64 ./cmd/cosmos-checkersd/main.go
	GOOS=linux GOARCH=arm64 go build -o ./build/cosmos-checkersd-linux-arm64 ./cmd/cosmos-checkersd/main.go

do-checksum-linux:
	cd build && sha256sum \
		cosmos-checkersd-linux-amd64 cosmos-checkersd-linux-arm64 \
		> cosmos-checkers-checksum-linux

build-with-checksum: build-linux do-checksum-linux

build-dockerfiles:
	docker build -f prod-sim/Dockerfile-cosmos-checkersd-debian . -t cosmos-checkersd_i --no-cache
	docker build -f prod-sim/Dockerfile-tmkms-debian . -t tmkms_i:v0.12.2 --no-cache

clean-validators:
	echo "You are removing all validators, but you must have to use sudo"

	echo desk-alice'\n'desk-bob'\n'node-carol'\n'sentry-alice'\n'sentry-bob'\n'val-alice'\n'val-bob'\n'kms-alice \
    | xargs -I {} \
    sudo rm -rf $(home-directory)/prod-sim/{}

create-validators:
	mkdir -p prod-sim/kms-alice
	mkdir -p prod-sim/node-carol
	mkdir -p prod-sim/sentry-alice
	mkdir -p prod-sim/sentry-bob
	mkdir -p prod-sim/val-alice
	mkdir -p prod-sim/val-bob
	mkdir -p prod-sim/desk-alice
	mkdir -p prod-sim/desk-bob

	docker run --name cosmos-checkers-container cosmos-checkersd_i 

	echo desk-alice'\n'desk-bob'\n'node-carol'\n'sentry-alice'\n'sentry-bob'\n'val-alice'\n'val-bob \
    | xargs -I {} \
    docker cp cosmos-checkers-container:/root/.cosmos-checkers/config \
	$(home-directory)/prod-sim/{}

	echo desk-alice'\n'desk-bob'\n'node-carol'\n'sentry-alice'\n'sentry-bob'\n'val-alice'\n'val-bob \
    | xargs -I {} \
    docker cp cosmos-checkers-container:/root/.cosmos-checkers/data \
	$(home-directory)/prod-sim/{}

	docker stop cosmos-checkers-container
	docker remove cosmos-checkers-container

build-all-steps-dockerfiles: build-dockerfiles create-validators
