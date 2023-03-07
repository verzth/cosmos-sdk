module github.com/verzth/cosmos-sdk/tests

go 1.20

require (
	github.com/verzth/cosmos-sdk/api v0.3.1
	github.com/verzth/cosmos-sdk/depinject v1.0.0-alpha.3
	github.com/verzth/cosmos-sdk/errors v1.0.0-beta.7
	github.com/verzth/cosmos-sdk/log v0.0.0-20230306220716-5e55f56d39d5
	github.com/verzth/cosmos-sdk/math v1.0.0-beta.6.0.20230216172121-959ce49135e4
	github.com/verzth/cosmos-sdk/simapp v0.0.0-00010101000000-000000000000
	github.com/verzth/cosmos-sdk/store v0.0.0-20230227103508-bbe7f8a11b44
	github.com/verzth/cosmos-sdk/x/evidence v0.1.0
	github.com/verzth/cosmos-sdk/x/feegrant v0.0.0-20230117113717-50e7c4a4ceff
	github.com/verzth/cosmos-sdk/x/nft v0.0.0-20230113085233-fae3332d62fc
	github.com/verzth/cosmos-sdk/x/tx v0.2.2
	github.com/verzth/cosmos-sdk/x/upgrade v0.0.0-20230127052425-54c8e1568335
	github.com/cometbft/cometbft v0.37.0
	github.com/cosmos/cosmos-db v1.0.0-rc.1
	github.com/cosmos/cosmos-proto v1.0.0-beta.2
	// this version is not used as it is always replaced by the latest Cosmos SDK version
	github.com/verzth/cosmos-sdk v0.48.0
	github.com/cosmos/gogoproto v1.4.6
	github.com/golang/mock v1.6.0
	github.com/google/uuid v1.3.0
	github.com/spf13/cobra v1.6.1
	github.com/stretchr/testify v1.8.2
	google.golang.org/protobuf v1.28.2-0.20230222093303-bc1253ad3743
	gotest.tools/v3 v3.4.0
	pgregory.net/rapid v0.5.5
)

// Here are the short-lived replace for tests
// Replace here are pending PRs, or version to be tagged.
// It must be in sync with SimApp temporary replaces
replace (
	// TODO tag all extracted modules after SDK refactor
	github.com/verzth/cosmos-sdk/client/v2 => ../client/v2
	github.com/verzth/cosmos-sdk/x/evidence => ../x/evidence
	github.com/verzth/cosmos-sdk/x/feegrant => ../x/feegrant
	github.com/verzth/cosmos-sdk/x/nft => ../x/nft
	github.com/verzth/cosmos-sdk/x/upgrade => ../x/upgrade
)

// Below are the long-lived replace for tests.
replace (
	// We always want to test against the latest version of the simapp.
	github.com/verzth/cosmos-sdk/simapp => ../simapp
	github.com/99designs/keyring => github.com/cosmos/keyring v1.2.0
	// We always want to test against the latest version of the SDK.
	github.com/verzth/cosmos-sdk => ../.
	// Fix upstream GHSA-h395-qcrw-5vmq vulnerability.
	// TODO Remove it: https://github.com/verzth/cosmos-sdk/issues/10409
	github.com/gin-gonic/gin => github.com/gin-gonic/gin v1.8.1
)
