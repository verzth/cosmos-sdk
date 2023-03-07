package testutil

import (
	"github.com/verzth/cosmos-sdk/core/appconfig"
	_ "github.com/verzth/cosmos-sdk/x/auth"           // import as blank for app wiring
	_ "github.com/verzth/cosmos-sdk/x/auth/tx/config" // import as blank for app wiring
	authtypes "github.com/verzth/cosmos-sdk/x/auth/types"
	_ "github.com/verzth/cosmos-sdk/x/bank" // import as blank for app wiring
	banktypes "github.com/verzth/cosmos-sdk/x/bank/types"
	_ "github.com/verzth/cosmos-sdk/x/consensus" // import as blank for app wiring
	consensustypes "github.com/verzth/cosmos-sdk/x/consensus/types"
	_ "github.com/verzth/cosmos-sdk/x/evidence" // import as blank for app wiring
	evidencetypes "github.com/verzth/cosmos-sdk/x/evidence/types"
	_ "github.com/verzth/cosmos-sdk/x/genutil" // import as blank for app wiring
	genutiltypes "github.com/verzth/cosmos-sdk/x/genutil/types"
	minttypes "github.com/verzth/cosmos-sdk/x/mint/types"
	_ "github.com/verzth/cosmos-sdk/x/params" // import as blank for app wiring
	paramstypes "github.com/verzth/cosmos-sdk/x/params/types"
	_ "github.com/verzth/cosmos-sdk/x/slashing" // import as blank for app wiring
	slashingtypes "github.com/verzth/cosmos-sdk/x/slashing/types"
	_ "github.com/verzth/cosmos-sdk/x/staking" // import as blank for app wiring
	stakingtypes "github.com/verzth/cosmos-sdk/x/staking/types"

	runtimev1alpha1 "github.com/verzth/cosmos-sdk/api/cosmos/app/runtime/v1alpha1"
	appv1alpha1 "github.com/verzth/cosmos-sdk/api/cosmos/app/v1alpha1"
	authmodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/auth/module/v1"
	bankmodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/bank/module/v1"
	consensusmodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/consensus/module/v1"
	evidencemodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/evidence/module/v1"
	genutilmodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/genutil/module/v1"
	paramsmodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/params/module/v1"
	slashingmodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/slashing/module/v1"
	stakingmodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/staking/module/v1"
	txconfigv1 "github.com/verzth/cosmos-sdk/api/cosmos/tx/config/v1"
)

var AppConfig = appconfig.Compose(&appv1alpha1.Config{
	Modules: []*appv1alpha1.ModuleConfig{
		{
			Name: "runtime",
			Config: appconfig.WrapAny(&runtimev1alpha1.Module{
				AppName: "EvidenceApp",
				BeginBlockers: []string{
					slashingtypes.ModuleName,
					evidencetypes.ModuleName,
					stakingtypes.ModuleName,
					genutiltypes.ModuleName,
				},
				EndBlockers: []string{
					stakingtypes.ModuleName,
					genutiltypes.ModuleName,
				},
				InitGenesis: []string{
					authtypes.ModuleName,
					banktypes.ModuleName,
					stakingtypes.ModuleName,
					slashingtypes.ModuleName,
					genutiltypes.ModuleName,
					evidencetypes.ModuleName,
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
			Name:   slashingtypes.ModuleName,
			Config: appconfig.WrapAny(&slashingmodulev1.Module{}),
		},
		{
			Name:   paramstypes.ModuleName,
			Config: appconfig.WrapAny(&paramsmodulev1.Module{}),
		},
		{
			Name:   consensustypes.ModuleName,
			Config: appconfig.WrapAny(&consensusmodulev1.Module{}),
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
			Name:   evidencetypes.ModuleName,
			Config: appconfig.WrapAny(&evidencemodulev1.Module{}),
		},
	},
})
