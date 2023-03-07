module github.com/verzth/cosmos-sdk/store

go 1.20

require (
	github.com/verzth/cosmos-sdk/errors v1.0.0-beta.7
	github.com/verzth/cosmos-sdk/log v0.0.0-20230306220716-5e55f56d39d5
	github.com/verzth/cosmos-sdk/math v1.0.0-beta.6.0.20230216172121-959ce49135e4
	github.com/armon/go-metrics v0.4.1
	github.com/cometbft/cometbft v0.37.0
	github.com/confio/ics23/go v0.9.0
	github.com/cosmos/cosmos-db v1.0.0-rc.1
	github.com/cosmos/gogoproto v1.4.6
	github.com/cosmos/iavl v0.21.0-alpha.1
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/hashicorp/golang-lru v0.5.5-0.20210104140557-80c98217689d
	github.com/spf13/cast v1.5.0
	github.com/stretchr/testify v1.8.2
	github.com/tidwall/btree v1.6.0
	golang.org/x/exp v0.0.0-20230224173230-c95f2b4c22f2
	google.golang.org/genproto v0.0.0-20230202175211-008b39050e57 // indirect
	google.golang.org/grpc v1.53.0 // indirect
	google.golang.org/protobuf v1.28.2-0.20230222093303-bc1253ad3743 // indirect
	gotest.tools/v3 v3.4.0
)

// Below are the long-lived replace for store.
// Fix upstream GHSA-h395-qcrw-5vmq vulnerability.
// TODO Remove it: https://github.com/verzth/cosmos-sdk/issues/10409
replace github.com/gin-gonic/gin => github.com/gin-gonic/gin v1.8.1
