module github.com/verzth/cosmos-sdk/tools/confix

go 1.20

require (
	github.com/verzth/cosmos-sdk v0.46.0-beta2.0.20230306220716-5e55f56d39d5
	github.com/creachadair/atomicfile v0.2.8
	github.com/creachadair/tomledit v0.0.24
	github.com/spf13/cobra v1.6.1
	github.com/spf13/viper v1.15.0
	golang.org/x/exp v0.0.0-20230224173230-c95f2b4c22f2
	gotest.tools/v3 v3.4.0
)

// Fix upstream GHSA-h395-qcrw-5vmq vulnerability.
// TODO Remove it: https://github.com/verzth/cosmos-sdk/issues/10409
// TODO investigate if we can outright delete this dependency, otherwise go install won't work :(
replace github.com/gin-gonic/gin => github.com/gin-gonic/gin v1.8.1
