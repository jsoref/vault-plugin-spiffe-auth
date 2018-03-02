Feature:
  In order to ensure the quality of the SPIFFE auth plugin
  As a developer
  I need to be able to test some features

Scenario: Succesful auth with SVID
    Given Vault is running and the plugin has been loaded
    When I authenticate with a valid SVID
    Then I expect a valid Vault Token
