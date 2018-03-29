package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/dcos/dcos-cli/pkg/httpclient"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// These are the different auth types that DC/OS supports with the names that they'll be given from the providers
// endpoint.
const (
	LoginTypeDCOSUidPassword     = "dcos-uid-password"
	LoginTypeDCOSUidServiceKey   = "dcos-uid-servicekey"
	LoginTypeDCOSUidPasswordLDAP = "dcos-uid-password-ldap"
	LoginTypeSAMLSpInitiated     = "saml-sp-initiated"
	LoginTypeOIDCAuthCodeFlow    = "oidc-authorization-code-flow"
	LoginTypeOIDCImplicitFlow    = "oidc-implicit-flow"
)

var jsonOutput bool

// authListProvidersCmd represents the `dcos auth list-providers` subcommand.
var authListProvidersCmd = &cobra.Command{
	Use:  "list-providers",
	Args: cobra.MaximumNArgs(1),
	RunE: listProviders,
}

func init() {
	authCmd.AddCommand(authListProvidersCmd)
	authListProvidersCmd.Flags().BoolVar(&jsonOutput, "json", false, "returns providers in json format")
}

func listProviders(cmd *cobra.Command, args []string) error {
	var url string
	if len(args) == 0 {
		conf := attachedCluster().Config
		url = conf.URL()
	} else {
		url = args[0]
	}

	providers, err := getProviders(url)
	if err != nil {
		return err
	}

	if jsonOutput {
		// Re-marshal it into json with indents added in for pretty printing.
		out, err := json.MarshalIndent(providers, "", "\t")
		if err != nil {
			return err
		}
		fmt.Println(string(out))
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"PROVIDER ID", "AUTHENTICATION TYPE"})
		// Turn off wrapping because it seems to wrap even if the column is set to be wide enough.
		table.SetAutoWrapText(false)
		table.SetBorder(false)
		table.SetRowSeparator(" ")
		table.SetColumnSeparator(" ")
		table.SetCenterSeparator(" ")

		for name, provider := range *providers {
			desc, err := loginTypeDescription(provider.AuthenticationType, provider)
			if err != nil {
				return err
			}
			table.Append([]string{name, desc})
		}
		table.Render()
	}

	return nil
}

func getProviders(baseURL string) (*map[string]loginProvider, error) {
	client := httpclient.New(baseURL)
	response, err := client.Get("/acs/api/v1/auth/providers")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var resp map[string]loginProvider
	err = json.NewDecoder(response.Body).Decode(&resp)
	return &resp, err
}

func loginTypeDescription(loginType string, provider loginProvider) (string, error) {
	switch loginType {
	case LoginTypeDCOSUidPassword:
		return "Log in using a standard DC/OS user account (username and password)", nil
	case LoginTypeDCOSUidServiceKey:
		return "Log in using a DC/OS service user account (username and private key)", nil
	case LoginTypeDCOSUidPasswordLDAP:
		return "Log in in using an LDAP user account (username and password)", nil
	case LoginTypeSAMLSpInitiated:
		return fmt.Sprintf("Log in using SAML 2.0 (%s)", provider.Description), nil
	case LoginTypeOIDCImplicitFlow:
		return fmt.Sprintf("Log in using OpenID Connect(%s)", provider.Description), nil
	case LoginTypeOIDCAuthCodeFlow:
		return fmt.Sprintf("Log in using OpenID Connect(%s)", provider.Description), nil
	default:
		return "", fmt.Errorf("unknown login provider %s", loginType)
	}
}

type loginProvider struct {
	AuthenticationType string                  `json:"authentication-type"`
	ClientMethod       string                  `json:"client-method"`
	Config             loginListProviderConfig `json:"config"`
	Description        string                  `json:"description"`
}

type loginListProviderConfig struct {
	StartFlowURL string `json:"start_flow_url"`
}
