package login

import (
	"encoding/json"
	"fmt"

	"github.com/dcos/dcos-cli/pkg/httpclient"
	"github.com/sirupsen/logrus"
)

// Client is able to detect available login providers and login to DC/OS.
type Client struct {
	http   *httpclient.Client
	logger *logrus.Logger
}

// NewClient creates a new login client.
func NewClient(baseClient *httpclient.Client, logger *logrus.Logger) *Client {
	return &Client{
		http:   baseClient,
		logger: logger,
	}
}

// Providers returns the supported login providers for a given DC/OS cluster.
func (c *Client) Providers() (Providers, error) {
	resp, err := c.http.Get("/acs/api/v1/auth/providers")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	providers := Providers{}
	if resp.StatusCode == 200 {
		err := json.NewDecoder(resp.Body).Decode(&providers)
		if err != nil {
			return nil, err
		}
		return providers, nil
	}

	c.logger.Info("Falling back to the WWW-Authenticate challenge.")

	// This is for DC/OS EE 1.7/1.8 and DC/OS Open Source.
	authHeader, err := c.challengeAuth()
	if err != nil {
		return nil, err
	}

	switch authHeader {
	case "oauthjwt":
		// DC/OS Open Source
		provider := defaultOIDCImplicitFlowProvider()
		providers[provider.ID] = provider
	case "acsjwt":
		// DC/OS EE 1.7/1.8
		provider := defaultDCOSUIDPasswordProvider()
		providers[provider.ID] = provider
	default:
		return nil, fmt.Errorf("unsupported WWW-Authenticate challenge '%s'", authHeader)
	}
	return providers, nil
}

// challengeAuth sends an unauthenticated HTTP request to a DC/OS well-known resource.
// It then expects a 401 response with a WWW-Authenticate header containing the login method to use.
// This method is used to determine which login provider is available when the /acs/api/v1/auth/providers
// endpoint is not present.
func (c *Client) challengeAuth() (string, error) {
	req, err := c.http.NewRequest("HEAD", "/pkgpanda/active.buildinfo.full.json", nil)
	if err != nil {
		return "", err
	}
	req.Header.Del("Authorization")

	resp, err := c.http.Do(req)
	if err != nil {
		return "", err
	}
	resp.Body.Close()

	if resp.StatusCode != 401 {
		return "", fmt.Errorf("expected status code 401, got %d", resp.StatusCode)
	}
	return resp.Header.Get("WWW-Authenticate"), nil
}
