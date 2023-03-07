//go:build e2e
// +build e2e

package testutil

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/verzth/cosmos-sdk/simapp"
	"github.com/verzth/cosmos-sdk/testutil/network"
)

func TestE2ETestSuite(t *testing.T) {
	cfg := network.DefaultConfig(simapp.NewTestNetworkFixture)
	cfg.NumValidators = 2
	suite.Run(t, NewE2ETestSuite(cfg))
}
