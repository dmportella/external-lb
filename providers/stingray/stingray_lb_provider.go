package stingray

import (
	"fmt"
	"os"

	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/dmportella/stingray"
	"github.com/rancher/external-lb/model"
	"github.com/rancher/external-lb/providers"
)

const (
	providerName = "Brocade Stingray"
	providerSlug = "brocade_stingray"
)

// Provider The brocade stingray provider implementation for rancher.
type Provider struct {
	client *stingray.Client
}

func init() {
	providers.RegisterProvider(providerSlug, new(Provider))
}

// GetName returns the name of the provider.
func (provider *StingrayIPProvider) GetName() string {
	return providerName
}

// Init initializes the provider.
func (provider *StingrayIPProvider) Init() error {
	stingrayEndpoint := os.Getenv("STINGRAY_ENDPOINT")
	if len(stingrayEndpoint) == 0 {
		return fmt.Errorf("STINGRAY_ENDPOINT is not set")
	}

	stingrayUsername := os.Getenv("STINGRAY_USERNAME")
	if len(stingrayUsername) == 0 {
		return fmt.Errorf("STINGRAY_USERNAME is not set")
	}

	stingrayPassword := os.Getenv("STINGRAY_PASSWORD")
	if len(stingrayPassword) == 0 {
		return fmt.Errorf("STINGRAY_PASSWORD is not set")
	}

	stingrayInsecureSkipVerify := strings.ToLower(os.Getenv("STINGRAY_INSECURESKIPVERIFY")) == "true"

	logrus.Debugf("Initializing %s provider with host: %s, admin: %s, pwd-length: %d, insecureSkipVerify: %t",
		provider.GetName(), stingrayEndpoint, stingrayUsername, len(stingrayPassword), stingrayInsecureSkipVerify)

	p.client, err := stingray.NewClient(stingray.Config{
        UserAgent: userAgent,
        HTTPEndpoint: stingrayEndpoint,
        InsecureSkipVerify: stingrayInsecureSkipVerify,
        UserName: stingrayUsername,
        Password: stingrayPassword,
    })

	if err := p.HealthCheck(); err != nil {
		return err
	}

	logrus.Infof("Configured %s provider using host %s", provider.GetName(), f5_host)
	return nil
}

// HealthCheck checks the connection to the provider.
func (provider *StingrayIPProvider) HealthCheck() error {
    if _, err := p.GetActionList("/"); err != nil {
		return fmt.Errorf("Could not connect to %s host '%s': %v", provider.GetName(), stingrayEndpoint, err)
	}
	return nil
}

// AddLBConfig adds a new endpoint configuration. It may
// return the FQDN for the endpoint if supported by the provider.
func (provider *StingrayIPProvider) AddLBConfig(config model.LBConfig) (fqdn string, err error) {
	return "", nil
}

// UpdateLBConfig updates the endpoint configuration. It may
// return the FQDN for the endpoint if supported by the provider.
func (provider *StingrayIPProvider) UpdateLBConfig(config model.LBConfig) (fqdn string, err error) {
	return "", nil
}

// RemoveLBConfig removes the specified endpoint configuration.
func (provider *StingrayIPProvider) RemoveLBConfig(config model.LBConfig) error {
	return nil
}

// GetLBConfigs returns all endpoint configurations.
func (provider *StingrayIPProvider) GetLBConfigs() ([]model.LBConfig, error) {
	return nil, nil
}
