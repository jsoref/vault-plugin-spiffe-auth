package tests

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/DATA-DOG/godog"
	"github.com/hashicorp/vault/api"
)

var vaultCommand *exec.Cmd
var vaultAPI *api.Client

func TestMain(m *testing.M) {
	format := "progress"
	for _, arg := range os.Args[1:] {
		if arg == "-test.v=true" { // go test transforms -v option
			format = "pretty"
			break
		}
	}
	status := godog.RunWithOptions("godog", func(s *godog.Suite) {
		godog.SuiteContext(s)
		FeatureContext(s)
	}, godog.Options{
		Format: format,
		Paths:  []string{"features"},
	})

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

func vaultIsRunningAndThePluginHasBeenLoaded() error {
	health, err := vaultAPI.Sys().Health()

	if err != nil {
		return err
	}

	if health.Sealed {
		return fmt.Errorf("Error Vault is Sealed")
	}

	// get SHA of plugin
	shasum, err := generatePluginSha()
	if err != nil {
		return err
	}

	// load plugin
	_, err = vaultAPI.Logical().Write(
		"sys/plugins/catalog/spiffe-auth",
		map[string]interface{}{
			"sha_256": shasum,
			"command": "spiffe-auth",
		},
	)
	if err != nil {
		return err
	}

	// enable the plugin
	err = vaultAPI.Sys().EnableAuthWithOptions(
		"spiffe",
		&api.EnableAuthOptions{
			Type:       "plugin",
			PluginName: "spiffe-auth",
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func iAuthenticateWithAValidSVID() error {
	return godog.ErrPending
}

func iExpectAValidVaultToken() error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^Vault is running and the plugin has been loaded$`, vaultIsRunningAndThePluginHasBeenLoaded)
	s.Step(`^I authenticate with a valid SVID$`, iAuthenticateWithAValidSVID)
	s.Step(`^I expect a valid Vault Token$`, iExpectAValidVaultToken)

	s.BeforeScenario(func(interface{}) {
		createConfig()

		vaultCommand = exec.Command("vault", "server", "-dev", "-dev-root-token-id=root", "-config=/tmp/vault.hcl")
		vaultCommand.Start()

		time.Sleep(5 * time.Second) // wait for startup

		c := api.Config{}
		c.Address = "http://localhost:8200"

		var err error
		vaultAPI, err = api.NewClient(&c)
		if err != nil {
			fmt.Println(err)
			return
		}

		vaultAPI.SetToken("root")
	})

	s.AfterScenario(func(interface{}, error) {
		vaultCommand.Process.Kill()
	})
}

func createConfig() {
	file, err := os.Create("/tmp/vault.hcl")
	if err != nil {
		fmt.Println(err)
		return
	}

	pluginFolder := filepath.Dir("../bin/")
	fmt.Fprintf(file, `plugin_directory = "%s"`, pluginFolder)
	file.Close()
}

func generatePluginSha() (string, error) {
	f, err := os.Open("../bin/spiffe-auth")
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
