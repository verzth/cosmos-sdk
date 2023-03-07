package simapp

import (
	"time"

	runtimev1alpha1 "github.com/verzth/cosmos-sdk/api/cosmos/app/runtime/v1alpha1"
	appv1alpha1 "github.com/verzth/cosmos-sdk/api/cosmos/app/v1alpha1"
	authmodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/auth/module/v1"
	authzmodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/authz/module/v1"
	bankmodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/bank/module/v1"
	capabilitymodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/capability/module/v1"
	consensusmodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/consensus/module/v1"
	crisismodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/crisis/module/v1"
	distrmodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/distribution/module/v1"
	evidencemodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/evidence/module/v1"
	feegrantmodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/feegrant/module/v1"
	genutilmodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/genutil/module/v1"
	govmodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/gov/module/v1"
	groupmodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/group/module/v1"
	mintmodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/mint/module/v1"
	nftmodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/nft/module/v1"
	paramsmodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/params/module/v1"
	slashingmodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/slashing/module/v1"
	stakingmodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/staking/module/v1"
	txconfigv1 "github.com/verzth/cosmos-sdk/api/cosmos/tx/config/v1"
	upgrademodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/upgrade/module/v1"
	vestingmodulev1 "github.com/verzth/cosmos-sdk/api/cosmos/vesting/module/v1"
	"github.com/verzth/cosmos-sdk/core/appconfig"
	evidencetypes "github.com/verzth/cosmos-sdk/x/evidence/types"
	"github.com/verzth/cosmos-sdk/x/nft"
	upgradetypes "github.com/verzth/cosmos-sdk/x/upgrade/types"
	"google.golang.org/protobuf/types/known/durationpb"

	"github.com/verzth/cosmos-sdk/x/feegrant"

	"github.com/verzth/cosmos-sdk/runtime"
	authtypes "github.com/verzth/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/verzth/cosmos-sdk/x/auth/vesting/types"
	"github.com/verzth/cosmos-sdk/x/authz"
	banktypes "github.com/verzth/cosmos-sdk/x/bank/types"
	capabilitytypes "github.com/verzth/cosmos-sdk/x/capability/types"
	consensustypes "github.com/verzth/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/verzth/cosmos-sdk/x/crisis/types"
	distrtypes "github.com/verzth/cosmos-sdk/x/distribution/types"
	genutiltypes "github.com/verzth/cosmos-sdk/x/genutil/types"
	govtypes "github.com/verzth/cosmos-sdk/x/gov/types"
	"github.com/verzth/cosmos-sdk/x/group"
	minttypes "github.com/verzth/cosmos-sdk/x/mint/types"
	paramstypes "github.com/verzth/cosmos-sdk/x/params/types"
	slashingtypes "github.com/verzth/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/verzth/cosmos-sdk/x/staking/types"
)

var (
	// module account permissions
	moduleAccPerms = []*authmodulev1.ModuleAccountPermission{
		{Account: authtypes.FeeCollectorName},
		{Account: distrtypes.ModuleName},
		{Account: minttypes.ModuleName, Permissions: []string{authtypes.Minter}},
		{Account: stakingtypes.BondedPoolName, Permissions: []string{authtypes.Burner, stakingtypes.ModuleName}},
		{Account: stakingtypes.NotBondedPoolName, Permissions: []string{authtypes.Burner, stakingtypes.ModuleName}},
		{Account: govtypes.ModuleName, Permissions: []string{authtypes.Burner}},
		{Account: nft.ModuleName},
	}

	// blocked account addresses
	blockAccAddrs = []string{
		authtypes.FeeCollectorName,
		distrtypes.ModuleName,
		minttypes.ModuleName,
		stakingtypes.BondedPoolName,
		stakingtypes.NotBondedPoolName,
		nft.ModuleName,
		// We allow the following module accounts to receive funds:
		// govtypes.ModuleName
	}

	// application configuration (used by depinject)
	AppConfig = appconfig.Compose(&appv1alpha1.Config{
		Modules: []*appv1alpha1.ModuleConfig{
			{
				Name: runtime.ModuleName,
				Config: appconfig.WrapAny(&runtimev1alpha1.Module{
					AppName: "SimApp",
					// During begin block slashing happens after distr.BeginBlocker so that
					// there is nothing left over in the validator fee pool, so as to keep the
					// CanWithdrawInvariant invariant.
					// NOTE: staking module is required if HistoricalEntries param > 0
					// NOTE: capability module's beginblocker must come before any modules using capabilities (e.g. IBC)
					BeginBlockers: []string{
						upgradetypes.ModuleName,
						capabilitytypes.ModuleName,
						minttypes.ModuleName,
						distrtypes.ModuleName,
						slashingtypes.ModuleName,
						evidencetypes.ModuleName,
						stakingtypes.ModuleName,
						genutiltypes.ModuleName,
						authz.ModuleName,
					},
					EndBlockers: []string{
						crisistypes.ModuleName,
						govtypes.ModuleName,
						stakingtypes.ModuleName,
						genutiltypes.ModuleName,
						feegrant.ModuleName,
						group.ModuleName,
					},
					OverrideStoreKeys: []*runtimev1alpha1.StoreKeyConfig{
						{
							ModuleName: authtypes.ModuleName,
							KvStoreKey: "acc",
						},
					},
					// NOTE: The genutils module must occur after staking so that pools are
					// properly initialized with tokens from genesis accounts.
					// NOTE: The genutils module must also occur after auth so that it can access the params from auth.
					// NOTE: Capability module must occur first so that it can initialize any capabilities
					// so that other modules that want to create or claim capabilities afterwards in InitChain
					// can do so safely.
					InitGenesis: []string{
						capabilitytypes.ModuleName,
						authtypes.ModuleName,
						banktypes.ModuleName,
						distrtypes.ModuleName,
						stakingtypes.ModuleName,
						slashingtypes.ModuleName,
						govtypes.ModuleName,
						minttypes.ModuleName,
						crisistypes.ModuleName,
						genutiltypes.ModuleName,
						evidencetypes.ModuleName,
						authz.ModuleName,
						feegrant.ModuleName,
						nft.ModuleName,
						group.ModuleName,
						paramstypes.ModuleName,
						upgradetypes.ModuleName,
						vestingtypes.ModuleName,
						consensustypes.ModuleName,
					},
					// When ExportGenesis is not specified, the export genesis module order
					// is equal to the init genesis order
					// ExportGenesis: []string{},
					// Uncomment if you want to set a custom migration order here.
					// OrderMigrations: []string{},
				}),
			},
			{
				Name: authtypes.ModuleName,
				Config: appconfig.WrapAny(&authmodulev1.Module{
					Bech32Prefix:             "cosmos",
					ModuleAccountPermissions: moduleAccPerms,
					// By default modules authority is the governance module. This is configurable with the following:
					// Authority: "group", // A custom module authority can be set using a module name
					// Authority: "cosmos1cwwv22j5ca08ggdv9c2uky355k908694z577tv", // or a specific address
				}),
			},
			{
				Name:   vestingtypes.ModuleName,
				Config: appconfig.WrapAny(&vestingmodulev1.Module{}),
			},
			{
				Name: banktypes.ModuleName,
				Config: appconfig.WrapAny(&bankmodulev1.Module{
					BlockedModuleAccountsOverride: blockAccAddrs,
				}),
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
				Name:   "tx",
				Config: appconfig.WrapAny(&txconfigv1.Config{}),
			},
			{
				Name:   genutiltypes.ModuleName,
				Config: appconfig.WrapAny(&genutilmodulev1.Module{}),
			},
			{
				Name:   authz.ModuleName,
				Config: appconfig.WrapAny(&authzmodulev1.Module{}),
			},
			{
				Name:   upgradetypes.ModuleName,
				Config: appconfig.WrapAny(&upgrademodulev1.Module{}),
			},
			{
				Name:   distrtypes.ModuleName,
				Config: appconfig.WrapAny(&distrmodulev1.Module{}),
			},
			{
				Name: capabilitytypes.ModuleName,
				Config: appconfig.WrapAny(&capabilitymodulev1.Module{
					SealKeeper: true,
				}),
			},
			{
				Name:   evidencetypes.ModuleName,
				Config: appconfig.WrapAny(&evidencemodulev1.Module{}),
			},
			{
				Name:   minttypes.ModuleName,
				Config: appconfig.WrapAny(&mintmodulev1.Module{}),
			},
			{
				Name: group.ModuleName,
				Config: appconfig.WrapAny(&groupmodulev1.Module{
					MaxExecutionPeriod: durationpb.New(time.Second * 1209600),
					MaxMetadataLen:     255,
				}),
			},
			{
				Name:   nft.ModuleName,
				Config: appconfig.WrapAny(&nftmodulev1.Module{}),
			},
			{
				Name:   feegrant.ModuleName,
				Config: appconfig.WrapAny(&feegrantmodulev1.Module{}),
			},
			{
				Name:   govtypes.ModuleName,
				Config: appconfig.WrapAny(&govmodulev1.Module{}),
			},
			{
				Name:   crisistypes.ModuleName,
				Config: appconfig.WrapAny(&crisismodulev1.Module{}),
			},
			{
				Name:   consensustypes.ModuleName,
				Config: appconfig.WrapAny(&consensusmodulev1.Module{}),
			},
		},
	})
)
