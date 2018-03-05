# vault-plugin-spiffe-auth
A gRPC based Vault plugin to allow authentication using a SPIFFE SVID

For initial plugin design please see [Design.md](Design.md)

This repository contains sample code for a HashiCorp Vault Auth Plugin. It is
both a real custom Vault auth method, and an example of how to build, install,
and maintain your own Vault auth plugin.

For more information on building a plugin, see the [accompanying blog post](https://www.hashicorp.com/blog/building-a-vault-secure-plugin).

## Generating a SVID
To generate an SVID use the example spire docker image

```bash
$ make generate_svid
```

This process will generate an SPIFFE SVID and output into the ./example_svid folder

A spire demo server with a running agent and server can be run by using the following command
```bash
$ make run_spire_docker
```

## Setup

The setup guide assumes some familiarity with Vault and Vault's plugin
ecosystem. You must have a Vault server already running, unsealed, and
authenticated.

1. Download and decompress the latest plugin binary from the Releases tab on
GitHub. Alternatively you can compile the plugin from source.

1. Move the compiled plugin into Vault's configured `plugin_directory`:

```sh
$ mv vault-auth-example /etc/vault/plugins/vault-auth-example
```

1. Calculate the SHA256 of the plugin and register it in Vault's plugin catalog.
If you are downloading the pre-compiled binary, it is highly recommended that
you use the published checksums to verify integrity.

```sh
$ export SHA256=$(shasum -a 256 "/etc/vault/plugins/vault-auth-example" | cut -d' ' -f1)

$ vault write sys/plugins/catalog/example-auth-plugin \
    sha_256="${SHA256}" \
    command="vault-auth-example"
```

1. Mount the auth method:

```sh
$ vault auth-enable \
    -path="example" \
    -plugin-name="example-auth-plugin" plugin
```

## Authenticating with the Shared Secret

To authenticate, the user supplies the shared secret:

```sh
$ vault write auth/example/login password="super-secret-password"
```

The response will be a standard auth response with some token metadata:

```text
Key             	Value
---             	-----
token           	b62420a6-ee83-22a4-7a15-a908af658c9f
token_accessor  	9eff2c4e-e321-3903-413e-a5084abb631e
token_duration  	30s
token_renewable 	true
token_policies  	[default my-policy other-policy]
token_meta_fruit	"banana"
```

## Should I Use This?

No, please do not. This is an example Vault Plugin that should be use for
learning purposes. Having a shared phrase that gives anyone access to Vault is
highly discouraged and a security anti-pattern. This code should be used for
educational purposes only.

## License

This code is licensed under the MPLv2 license.
Useful Links:
* https://www.hashicorp.com/blog/building-a-vault-secure-plugin
* https://github.com/hashicorp/vault/tree/master/logical/plugin
* https://github.com/hashicorp/go-plugin
* https://www.vaultproject.io/api/secret/identity/index.html
* https://github.com/hashicorp/vault/blob/master/logical/auth.go#L61
* https://github.com/hashicorp/vault/blob/master/helper/identity/types.proto

