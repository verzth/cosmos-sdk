package client

import (
	govclient "github.com/verzth/cosmos-sdk/x/gov/client"
	"github.com/verzth/cosmos-sdk/x/params/client/cli"
)

// ProposalHandler is the param change proposal handler.
var ProposalHandler = govclient.NewProposalHandler(cli.NewSubmitParamChangeProposalTxCmd)
