//go:build !integration
// +build !integration

package encrypt_test

import (
	"testing"

	"github.com/Trousseau-io/trousseau-tsh/pkg/config"
	"github.com/Trousseau-io/trousseau-tsh/pkg/encrypt"
	"github.com/stretchr/testify/assert"
)

type testConfig struct{}

func (t *testConfig) GetProvider() string {
	return "vault"
}
func (t *testConfig) GetVaultConfig() config.VaultConfig {
	return config.VaultConfig{
		Address: "http://localhost:9200",
		Token:   "test",
	}
}
func TestCreatingVaultClient(t *testing.T) {
	cfg := testConfig{}
	_, err := encrypt.NewService(&cfg)
	assert.NoError(t, err)

}
