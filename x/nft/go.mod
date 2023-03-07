module github.com/verzth/cosmos-sdk/x/nft

go 1.20

require (
	github.com/verzth/cosmos-sdk/api v0.3.1
	github.com/verzth/cosmos-sdk/core v0.6.0
	github.com/verzth/cosmos-sdk/depinject v1.0.0-alpha.3
	github.com/verzth/cosmos-sdk/errors v1.0.0-beta.7
	github.com/verzth/cosmos-sdk/math v1.0.0-beta.6.0.20230216172121-959ce49135e4
	github.com/verzth/cosmos-sdk/store v0.0.0-20230227103508-bbe7f8a11b44
	github.com/cometbft/cometbft v0.37.0
	github.com/cosmos/cosmos-proto v1.0.0-beta.2
	github.com/verzth/cosmos-sdk v0.46.0-beta2.0.20230306220716-5e55f56d39d5
	github.com/cosmos/gogoproto v1.4.6
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/spf13/cobra v1.6.1
	github.com/stretchr/testify v1.8.2
	google.golang.org/genproto v0.0.0-20230202175211-008b39050e57
	google.golang.org/grpc v1.53.0
)

// Below are the long-lived replace of the Cosmos SDK
// Fix upstream GHSA-h395-qcrw-5vmq vulnerability.
// TODO Remove it: https://github.com/verzth/cosmos-sdk/issues/10409
replace github.com/gin-gonic/gin => github.com/gin-gonic/gin v1.8.1
