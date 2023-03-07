go 1.20

module github.com/verzth/cosmos-sdk

require (
	github.com/verzth/cosmos-sdk/api v0.3.1
	github.com/verzth/cosmos-sdk/collections v0.0.0-20230214153846-b6c6e4e99177
	github.com/verzth/cosmos-sdk/core v0.6.0
	github.com/verzth/cosmos-sdk/depinject v1.0.0-alpha.3
	github.com/verzth/cosmos-sdk/errors v1.0.0-beta.7
	github.com/verzth/cosmos-sdk/log v0.0.0-20230306220716-5e55f56d39d5
	github.com/verzth/cosmos-sdk/math v1.0.0-beta.6.0.20230216172121-959ce49135e4
	github.com/verzth/cosmos-sdk/store v0.0.0-20230227103508-bbe7f8a11b44
	github.com/verzth/cosmos-sdk/x/tx v0.2.2
	github.com/99designs/keyring v1.2.1
	github.com/armon/go-metrics v0.4.1
	github.com/bgentry/speakeasy v0.1.1-0.20220910012023-760eaf8b6816
	github.com/chzyer/readline v1.5.1
	github.com/cockroachdb/apd/v2 v2.0.2
	github.com/cockroachdb/errors v1.9.1
	github.com/cometbft/cometbft v0.37.0
	github.com/cosmos/btcutil v1.0.5
	github.com/cosmos/cosmos-db v1.0.0-rc.1
	github.com/cosmos/cosmos-proto v1.0.0-beta.2
	github.com/verzth/cosmos-sdk/db v1.0.0-beta.1.0.20220726092710-f848e4300a8a
	github.com/cosmos/go-bip39 v1.0.0
	github.com/cosmos/gogogateway v1.2.0
	github.com/cosmos/gogoproto v1.4.6
	github.com/cosmos/ledger-cosmos-go v0.13.0
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.1.0
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.2
	github.com/google/gofuzz v1.2.0
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/hashicorp/golang-lru v0.5.5-0.20210104140557-80c98217689d
	github.com/hdevalence/ed25519consensus v0.1.0
	github.com/huandu/skiplist v1.2.0
	github.com/improbable-eng/grpc-web v0.15.0
	github.com/jhump/protoreflect v1.15.1
	github.com/magiconair/properties v1.8.7
	github.com/manifoldco/promptui v0.9.0
	github.com/mattn/go-isatty v0.0.17
	github.com/prometheus/client_golang v1.14.0
	github.com/prometheus/common v0.41.0
	github.com/rs/zerolog v1.29.0
	github.com/spf13/cast v1.5.0
	github.com/spf13/cobra v1.6.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.15.0
	github.com/stretchr/testify v1.8.2
	github.com/tendermint/go-amino v0.16.0
	golang.org/x/crypto v0.7.0
	golang.org/x/exp v0.0.0-20230224173230-c95f2b4c22f2
	golang.org/x/sync v0.1.0
	google.golang.org/genproto v0.0.0-20230202175211-008b39050e57
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.28.2-0.20230222093303-bc1253ad3743
	gotest.tools/v3 v3.4.0
	pgregory.net/rapid v0.5.5
	sigs.k8s.io/yaml v1.3.0
)

// Below are the long-lived replace of the Cosmos SDK
replace (
	// use cosmos fork of keyring
	github.com/99designs/keyring => github.com/cosmos/keyring v1.2.0
	// dgrijalva/jwt-go is deprecated and doesn't receive security updates.
	// TODO: remove it: https://github.com/verzth/cosmos-sdk/issues/13134
	github.com/dgrijalva/jwt-go => github.com/golang-jwt/jwt/v4 v4.4.2
	// Fix upstream GHSA-h395-qcrw-5vmq vulnerability.
	// TODO Remove it: https://github.com/verzth/cosmos-sdk/issues/10409
	github.com/gin-gonic/gin => github.com/gin-gonic/gin v1.8.1
	// Downgraded to avoid bugs in following commits which caused simulations to fail.
	github.com/syndtr/goleveldb => github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7
)

retract (
	// subject to a bug in the group module and gov module migration
	[v0.46.5, v0.46.6]
	// subject to the dragonberry vulnerability
	// and/or the bank coin metadata migration issue
	[v0.46.0, v0.46.4]
	// subject to the dragonberry vulnerability
	[v0.45.0, v0.45.8]
	// do not use
	v0.43.0
)
