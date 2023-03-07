module github.com/verzth/cosmos-sdk/simapp

go 1.20

require (
	github.com/verzth/cosmos-sdk/api v0.3.1
	github.com/verzth/cosmos-sdk/client/v2 v2.0.0-20230220152935-67f04e629623
	github.com/verzth/cosmos-sdk/core v0.6.0
	github.com/verzth/cosmos-sdk/depinject v1.0.0-alpha.3
	github.com/verzth/cosmos-sdk/log v0.0.0-20230306220716-5e55f56d39d5
	github.com/verzth/cosmos-sdk/math v1.0.0-beta.6.0.20230216172121-959ce49135e4
	github.com/verzth/cosmos-sdk/store v0.0.0-20230227103508-bbe7f8a11b44
	github.com/verzth/cosmos-sdk/tools/confix v0.0.0-20230120150717-4f6f6c00021f
	github.com/verzth/cosmos-sdk/tools/rosetta v0.2.0
	github.com/verzth/cosmos-sdk/x/evidence v0.1.0
	github.com/verzth/cosmos-sdk/x/feegrant v0.0.0-20230117113717-50e7c4a4ceff
	github.com/verzth/cosmos-sdk/x/nft v0.0.0-20230113085233-fae3332d62fc
	github.com/verzth/cosmos-sdk/x/upgrade v0.0.0-20230127052425-54c8e1568335
	github.com/cometbft/cometbft v0.37.0
	github.com/cosmos/cosmos-db v1.0.0-rc.1
	// this version is not used as it is always replaced by the latest Cosmos SDK version
	github.com/verzth/cosmos-sdk v0.48.0
	github.com/cosmos/gogoproto v1.4.6
	github.com/golang/mock v1.6.0
	github.com/spf13/cast v1.5.0
	github.com/spf13/cobra v1.6.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.15.0
	github.com/stretchr/testify v1.8.2
	google.golang.org/protobuf v1.28.2-0.20230222093303-bc1253ad3743
)

// Here are the short-lived replace from the SimApp
// Replace here are pending PRs, or version to be tagged
replace (
	// TODO tag all extracted modules after SDK refactor
	github.com/verzth/cosmos-sdk/tools/confix => ../tools/confix
	github.com/verzth/cosmos-sdk/tools/rosetta => ../tools/rosetta
	github.com/verzth/cosmos-sdk/x/evidence => ../x/evidence
	github.com/verzth/cosmos-sdk/x/feegrant => ../x/feegrant
	github.com/verzth/cosmos-sdk/x/nft => ../x/nft
	github.com/verzth/cosmos-sdk/x/upgrade => ../x/upgrade
)

// Below are the long-lived replace of the SimApp
replace (
	// use cosmos fork of keyring
	github.com/99designs/keyring => github.com/cosmos/keyring v1.2.0
	// Simapp always use the latest version of the cosmos-sdk
	github.com/verzth/cosmos-sdk => ../.
	// Fix upstream GHSA-h395-qcrw-5vmq vulnerability.
	// TODO Remove it: https://github.com/verzth/cosmos-sdk/issues/10409
	github.com/gin-gonic/gin => github.com/gin-gonic/gin v1.8.1
	// Downgraded to avoid bugs in following commits which caused simulations to fail.
	github.com/syndtr/goleveldb => github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7
)
