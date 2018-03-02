package tests

import (
	"os"
	"testing"

	"github.com/DATA-DOG/godog"
)

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
	return godog.ErrPending
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
}
