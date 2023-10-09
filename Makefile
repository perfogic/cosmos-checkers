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
	docker build -f prod-sim/Dockerfile-tmkms-debian . -t tmkms_i:v0.12.2

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

	echo desk-alice'\n'desk-bob'\n'node-carol'\n'sentry-alice'\n'sentry-bob'\n'val-alice'\n'val-bob \
    | xargs -I {} \
	cosmos-checkersd init cosmos-checkers --chain-id "cosmos-checkers-1"\
		--home $(home-directory)/prod-sim/{} \

	echo desk-alice'\n'desk-bob'\n'node-carol'\n'sentry-alice'\n'sentry-bob'\n'val-alice'\n'val-bob \
    | xargs -I {} \
	sed -i 's/"stake"/"upawn"/g' $(home-directory)/prod-sim/{}/config/genesis.json\

	echo desk-alice'\n'desk-bob'\n'node-carol'\n'sentry-alice'\n'sentry-bob'\n'val-alice'\n'val-bob \
    | xargs -I {} \
	sed -Ei 's/([0-9]+)stake/\1upawn/g' $(home-directory)/prod-sim/{}/config/app.toml\

	echo desk-alice'\n'desk-bob'\n'node-carol'\n'sentry-alice'\n'sentry-bob'\n'val-alice'\n'val-bob \
    | xargs -I {} \
	sed -Ei 's/^chain-id = .*/chain-id = "cosmos-checkers-1"/g' $(home-directory)/prod-sim/{}/config/client.toml

add-keys-alice:
	cosmos-checkersd keys add alice --keyring-backend test --keyring-dir $(home-directory)/prod-sim/desk-alice/keys
add-keys-bob:
	cosmos-checkersd keys add bob --keyring-backend test --keyring-dir $(home-directory)/prod-sim/desk-bob/keys

prepare-tmkms:
	docker run --name tmkms-instance tmkms_i:v0.12.2
	docker cp tmkms-instance:/root/tmkms $(home-directory)/prod-sim/kms-alice
	docker remove tmkms-instance

import-concensus-key:
	cosmos-checkersd tendermint show-validator --home $(home-directory)/prod-sim/val-alice/ | tr -d '\n' | tr -d '\r' > $(home-directory)/prod-sim/desk-alice/config/pub_validator_key-val-alice.json

	cp $(home-directory)/prod-sim/val-alice/config/priv_validator_key.json $(home-directory)/prod-sim/desk-alice/config/priv_validator_key-val-alice.json 
	mv $(home-directory)/prod-sim/val-alice/config/priv_validator_key.json $(home-directory)/prod-sim/kms-alice/secrets/priv_validator_key-val-alice.json

	echo "===============================================================\n===============================================================\nType: tmkms softsign import secrets/priv_validator_key-val-alice.json secrets/val-alice-consensus.key, then chmod it\n"

	docker run --rm -it \
    -v $(home-directory)/prod-sim/kms-alice:/root/tmkms \
    -w /root/tmkms \
    tmkms_i:v0.12.2

	cp prod-sim/sentry-alice/config/priv_validator_key.json \
    prod-sim/val-alice/config/

init-balances:
	# address of alice from desk-alice
	cosmos-checkersd add-genesis-account cosmos1qkemfq9evv9ds6azcmlls89uf5qupchd3hq7sk 1000000000upawn --home=$(home-directory)/prod-sim/desk-alice

	# address of bob from desk-bob
	cosmos-checkersd add-genesis-account cosmos1d0z95l6p9u5eqkesg62nxfztqja9pnkxy8eg20 500000000upawn --home=$(home-directory)/prod-sim/desk-bob

init-stakes:
	cp prod-sim/val-bob/config/priv_validator_key.json \
    prod-sim/desk-bob/config/priv_validator_key.json

	cosmos-checkersd gentx bob 40000000upawn \
    --keyring-backend test --keyring-dir $(home-directory)/prod-sim/desk-bob/keys \
    --account-number 0 --sequence 0 \
    --chain-id cosmos-checkers-1 \
    --gas 1000000 \
	--home $(home-directory)/prod-sim/desk-bob \
    --gas-prices 0.1upawn

	mv prod-sim/desk-bob/config/genesis.json \
    prod-sim/desk-alice/config/

	cosmos-checkersd gentx alice 60000000upawn \
    --keyring-backend test --keyring-dir $(home-directory)/prod-sim/desk-alice/keys \
    --account-number 0 --sequence 0 \
    --pubkey $(cat $(home-directory)/prod-sim/desk-alice/config/pub_validator_key-val-alice.json) \
    --chain-id checkers-1 \
	--home $(home-directory)/prod-sim/desk-alice \
    --gas 1000000 \
    --gas-prices 0.1upawn

	cp prod-sim/desk-bob/config/gentx/gentx-* \
    prod-sim/desk-alice/config/gentx

	cosmos-checkersd collect-gentxs --home $(home-directory)/prod-sim/desk-alice

genesis-distribution:
	echo desk-bob'\n'node-carol'\n'sentry-alice'\n'sentry-bob'\n'val-alice'\n'val-bob \
		| xargs -I {} \
		cp prod-sim/desk-alice/config/genesis.json prod-sim/{}/config

simulate-validators-with-docker: build-dockerfiles create-validators add-keys-alice add-keys-bob
