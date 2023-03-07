//go:build e2e
// +build e2e

package group

import (
	"testing"

	"github.com/verzth/cosmos-sdk/simapp"
	"github.com/verzth/cosmos-sdk/testutil/network"

	"github.com/stretchr/testify/suite"
)

func TestE2ETestSuite(t *testing.T) {
	cfg := network.DefaultConfig(simapp.NewTestNetworkFixture)
	cfg.NumValidators = 2
	suite.Run(t, NewE2ETestSuite(cfg))
}
