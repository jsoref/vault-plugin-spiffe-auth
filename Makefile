build:
	go build -o ./bin/spiffe-auth

build_arm:
	CGO_ENABLED=0	GOOS=linux GOARCH=arm GOARM=6 go build -o ./bin/spiffe-auth

functional_test:
	go test -v functional_tests/main_test.go
start_vault:
	./create_vault_config.sh
	
	vault server -dev -dev-root-token-id="root" -config=/tmp/vault.hcl &
	sleep 5
	VAULT_ADDR=http://127.0.0.1:8200 vault login root

	@echo ""
	@echo "Vault has been started in dev mode set the environment variable VAULT_ADDR=http://127.0.0.1:8200 and use the vault token listed in the above output"


install_plugin: build 
	VAULT_ADDR=http://127.0.0.1:8200 vault write sys/plugins/catalog/spiffe-auth \
	  sha_256="$(shell shasum -a 256 "./bin/spiffe-auth" | cut -d " " -f1)" \
 		command="vault-auth-spiffe"

run_spire_docker:
	docker run -it --name spire -v $(shell pwd)/example_svid:/out nicholasjackson/spire
	docker rm -f spire

generate_svid:
	@docker run -d -it --name spire -v $(shell pwd)/example_svid:/out nicholasjackson/spire > /dev/null
	@sleep 10
	docker exec -it spire /bin/bash -c 'spire-server entry create \
    -parentID spiffe://example.org/host \
    -spiffeID spiffe://example.org/host/workload \
    -selector unix:uid:`id -u workload`'
	@sleep 10
	docker exec -it spire /bin/bash -c 'su -c "spire-agent api fetch -write /out" workload'
	@echo "Output example SVID to ./example_svid"
	@docker rm -f spire > /dev/null
