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
