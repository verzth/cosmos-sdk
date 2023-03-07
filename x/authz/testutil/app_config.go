package testutil

import (
	_ "github.com/verzth/cosmos-sdk/x/auth"           // import as blank for app wiring
	_ "github.com/verzth/cosmos-sdk/x/auth/tx/config" // import as blank for app wiring
	_ "github.com/verzth/cosmos-sdk/x/authz/module"   // import as blank for app wiring
	_ "github.com/verzth/cosmos-sdk/x/bank"           // import as blank for app wiring
	_ "github.com/verzth/cosmos-sdk/x/consensus"      // import as blank for app wiring
	_ "github.com/verzth/cosmos-sdk/x/genutil"        // import as blank for app wiring
	_ "github.com/verzth/cosmos-sdk/x/gov"            // import as blank for app wiring
	_ "github.com/verzth/cosmos-sdk/x/mint"           // import as blank for app wiring
	_ "github.com/verzth/cosmos-sdk/x/params"         // import as blank for app wiring
	_ "github.com/verzth/cosmos-sdk/x/staking"        // import as blank for app wiring

	txconfigv1 "github.com/verzth/cosmos-sdk/api/cosmos/tx/config/v1"
	"github.com/verzth/cosmos-sdk/core/appconfig"
	authtypes "github.com/verzth/cosmos-sdk/x/auth/types"
	"github.com/verzth/cosmos-sdk/x/authz"
	banktypes "github.com/verzth/cosmos-sdk/x/bank/types"
	consensustypes "github.com/verzth/cosmos-sdk/x/consensus/types"
	genutiltypes "github.com/verzth/cosmos-sdk/x/genutil/types"
	govtypes "github.com/verzth/cosmos-sdk/x/gov/types"
	minttypes "github.com/verzth/cosmos-sdk/x/mint/types"
	paramstypes "github.com/verzth/cosmos-sdk/x/params/types"
	stakingtypes "github.com/verzth/cosmos-sdk/x/staking/types"

	runtimev1alpha1 "github.com/verzth/cosmos-sdk/api/cosmos/app/runtime/v1alpha1"
	appv1alpha1 "github.com/verzth/cosmos-sdk/api/cosmos/app/v1alpha1"
	authmodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/auth/module/v1"
	authzmodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/authz/module/v1"
	bankmodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/bank/module/v1"
	consensusmodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/consensus/module/v1"
	genutilmodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/genutil/module/v1"
	mintmodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/mint/module/v1"
	paramsmodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/params/module/v1"
	stakingmodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/staking/module/v1"
)

var AppConfig = appconfig.Compose(&appv1alpha1.Config{
	Modules: []*appv1alpha1.ModuleConfig{
		{
			Name: "runtime",
			Config: appconfig.WrapAny(&runtimev1alpha1.Module{
				AppName: "AuthzApp",
				BeginBlockers: []string{
					minttypes.ModuleName,
					stakingtypes.ModuleName,
					genutiltypes.ModuleName,
					authz.ModuleName,
				},
				EndBlockers: []string{
					stakingtypes.ModuleName,
					genutiltypes.ModuleName,
				},
				InitGenesis: []string{
					authtypes.ModuleName,
					banktypes.ModuleName,
					stakingtypes.ModuleName,
					minttypes.ModuleName,
					genutiltypes.ModuleName,
					authz.ModuleName,
					paramstypes.ModuleName,
					consensustypes.ModuleName,
				},
			}),
		},
		{
			Name: authtypes.ModuleName,
			Config: appconfig.WrapAny(&authmodulev1.Module{
				Bech32Prefix: "cosmos",
				ModuleAccountPermissions: []*authmodulev1.ModuleAccountPermission{
					{Account: authtypes.FeeCollectorName},
					{Account: minttypes.ModuleName, Permissions: []string{authtypes.Minter}},
					{Account: stakingtypes.BondedPoolName, Permissions: []string{authtypes.Burner, stakingtypes.ModuleName}},
					{Account: stakingtypes.NotBondedPoolName, Permissions: []string{authtypes.Burner, stakingtypes.ModuleName}},
					{Account: govtypes.ModuleName, Permissions: []string{authtypes.Burner}},
				},
			}),
		},
		{
			Name:   banktypes.ModuleName,
			Config: appconfig.WrapAny(&bankmodulev1.Module{}),
		},
		{
			Name:   stakingtypes.ModuleName,
			Config: appconfig.WrapAny(&stakingmodulev1.Module{}),
		},
		{
			Name:   paramstypes.ModuleName,
			Config: appconfig.WrapAny(&paramsmodulev1.Module{}),
		},
		{
			Name:   "tx",
			Config: appconfig.WrapAny(&txconfigv1.Config{}),
		},
		{
			Name:   genutiltypes.ModuleName,
			Config: appconfig.WrapAny(&genutilmodulev1.Module{}),
		},
		{
			Name:   consensustypes.ModuleName,
			Config: appconfig.WrapAny(&consensusmodulev1.Module{}),
		},
		{
			Name:   authz.ModuleName,
			Config: appconfig.WrapAny(&authzmodulev1.Module{}),
		},
		{
			Name:   minttypes.ModuleName,
			Config: appconfig.WrapAny(&mintmodulev1.Module{}),
		},
	},
})
