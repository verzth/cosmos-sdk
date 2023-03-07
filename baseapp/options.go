package baseapp

import (
	"fmt"
	"io"

	dbm "github.com/cosmos/cosmos-db"
	"github.com/spf13/cast"
	"github.com/verzth/cosmos-sdk/store/metrics"
	pruningtypes "github.com/verzth/cosmos-sdk/store/pruning/types"
	"github.com/verzth/cosmos-sdk/store/snapshots"
	snapshottypes "github.com/verzth/cosmos-sdk/store/snapshots/types"
	"github.com/verzth/cosmos-sdk/store/streaming"
	storetypes "github.com/verzth/cosmos-sdk/store/types"

	"github.com/verzth/cosmos-sdk/client/flags"
	"github.com/verzth/cosmos-sdk/codec/types"
	servertypes "github.com/verzth/cosmos-sdk/server/types"
	sdk "github.com/verzth/cosmos-sdk/types"
	"github.com/verzth/cosmos-sdk/types/mempool"
)

// File for storing in-package BaseApp optional functions,
// for options that need access to non-exported fields of the BaseApp

// SetPruning sets a pruning option on the multistore associated with the app
func SetPruning(opts pruningtypes.PruningOptions) func(*BaseApp) {
	return func(bapp *BaseApp) { bapp.cms.SetPruning(opts) }
}

// SetMinGasPrices returns an option that sets the minimum gas prices on the app.
func SetMinGasPrices(gasPricesStr string) func(*BaseApp) {
	gasPrices, err := sdk.ParseDecCoins(gasPricesStr)
	if err != nil {
		panic(fmt.Sprintf("invalid minimum gas prices: %v", err))
	}

	return func(bapp *BaseApp) { bapp.setMinGasPrices(gasPrices) }
}

// SetHaltHeight returns a BaseApp option function that sets the halt block height.
func SetHaltHeight(blockHeight uint64) func(*BaseApp) {
	return func(bapp *BaseApp) { bapp.setHaltHeight(blockHeight) }
}

// SetHaltTime returns a BaseApp option function that sets the halt block time.
func SetHaltTime(haltTime uint64) func(*BaseApp) {
	return func(bapp *BaseApp) { bapp.setHaltTime(haltTime) }
}

// SetMinRetainBlocks returns a BaseApp option function that sets the minimum
// block retention height value when determining which heights to prune during
// ABCI Commit.
func SetMinRetainBlocks(minRetainBlocks uint64) func(*BaseApp) {
	return func(bapp *BaseApp) { bapp.setMinRetainBlocks(minRetainBlocks) }
}

// SetTrace will turn on or off trace flag
func SetTrace(trace bool) func(*BaseApp) {
	return func(app *BaseApp) { app.setTrace(trace) }
}

// SetIndexEvents provides a BaseApp option function that sets the events to index.
func SetIndexEvents(ie []string) func(*BaseApp) {
	return func(app *BaseApp) { app.setIndexEvents(ie) }
}

// SetIAVLCacheSize provides a BaseApp option function that sets the size of IAVL cache.
func SetIAVLCacheSize(size int) func(*BaseApp) {
	return func(bapp *BaseApp) { bapp.cms.SetIAVLCacheSize(size) }
}

// SetIAVLDisableFastNode enables(false)/disables(true) fast node usage from the IAVL store.
func SetIAVLDisableFastNode(disable bool) func(*BaseApp) {
	return func(bapp *BaseApp) { bapp.cms.SetIAVLDisableFastNode(disable) }
}

// SetIAVLLazyLoading enables/disables lazy loading of the IAVL store.
func SetIAVLLazyLoading(lazyLoading bool) func(*BaseApp) {
	return func(bapp *BaseApp) { bapp.cms.SetLazyLoading(lazyLoading) }
}

// SetInterBlockCache provides a BaseApp option function that sets the
// inter-block cache.
func SetInterBlockCache(cache storetypes.MultiStorePersistentCache) func(*BaseApp) {
	return func(app *BaseApp) { app.setInterBlockCache(cache) }
}

// SetSnapshot sets the snapshot store.
func SetSnapshot(snapshotStore *snapshots.Store, opts snapshottypes.SnapshotOptions) func(*BaseApp) {
	return func(app *BaseApp) { app.SetSnapshot(snapshotStore, opts) }
}

// SetMempool sets the mempool on BaseApp.
func SetMempool(mempool mempool.Mempool) func(*BaseApp) {
	return func(app *BaseApp) { app.SetMempool(mempool) }
}

func (app *BaseApp) SetName(name string) {
	if app.sealed {
		panic("SetName() on sealed BaseApp")
	}

	app.name = name
}

// SetParamStore sets a parameter store on the BaseApp.
func (app *BaseApp) SetParamStore(ps ParamStore) {
	if app.sealed {
		panic("SetParamStore() on sealed BaseApp")
	}

	app.paramStore = ps
}

// SetVersion sets the application's version string.
func (app *BaseApp) SetVersion(v string) {
	if app.sealed {
		panic("SetVersion() on sealed BaseApp")
	}
	app.version = v
}

// SetProtocolVersion sets the application's protocol version
func (app *BaseApp) SetProtocolVersion(v uint64) {
	app.appVersion = v
}

func (app *BaseApp) SetDB(db dbm.DB) {
	if app.sealed {
		panic("SetDB() on sealed BaseApp")
	}

	app.db = db
}

func (app *BaseApp) SetCMS(cms storetypes.CommitMultiStore) {
	if app.sealed {
		panic("SetEndBlocker() on sealed BaseApp")
	}

	app.cms = cms
}

func (app *BaseApp) SetInitChainer(initChainer sdk.InitChainer) {
	if app.sealed {
		panic("SetInitChainer() on sealed BaseApp")
	}

	app.initChainer = initChainer
}

func (app *BaseApp) SetBeginBlocker(beginBlocker sdk.BeginBlocker) {
	if app.sealed {
		panic("SetBeginBlocker() on sealed BaseApp")
	}

	app.beginBlocker = beginBlocker
}

func (app *BaseApp) SetEndBlocker(endBlocker sdk.EndBlocker) {
	if app.sealed {
		panic("SetEndBlocker() on sealed BaseApp")
	}

	app.endBlocker = endBlocker
}

func (app *BaseApp) SetAnteHandler(ah sdk.AnteHandler) {
	if app.sealed {
		panic("SetAnteHandler() on sealed BaseApp")
	}

	app.anteHandler = ah
}

func (app *BaseApp) SetPostHandler(ph sdk.PostHandler) {
	if app.sealed {
		panic("SetPostHandler() on sealed BaseApp")
	}

	app.postHandler = ph
}

func (app *BaseApp) SetAddrPeerFilter(pf sdk.PeerFilter) {
	if app.sealed {
		panic("SetAddrPeerFilter() on sealed BaseApp")
	}

	app.addrPeerFilter = pf
}

func (app *BaseApp) SetIDPeerFilter(pf sdk.PeerFilter) {
	if app.sealed {
		panic("SetIDPeerFilter() on sealed BaseApp")
	}

	app.idPeerFilter = pf
}

func (app *BaseApp) SetFauxMerkleMode() {
	if app.sealed {
		panic("SetFauxMerkleMode() on sealed BaseApp")
	}

	app.fauxMerkleMode = true
}

// SetCommitMultiStoreTracer sets the store tracer on the BaseApp's underlying
// CommitMultiStore.
func (app *BaseApp) SetCommitMultiStoreTracer(w io.Writer) {
	app.cms.SetTracer(w)
}

// SetStoreLoader allows us to customize the rootMultiStore initialization.
func (app *BaseApp) SetStoreLoader(loader StoreLoader) {
	if app.sealed {
		panic("SetStoreLoader() on sealed BaseApp")
	}

	app.storeLoader = loader
}

// SetSnapshot sets the snapshot store and options.
func (app *BaseApp) SetSnapshot(snapshotStore *snapshots.Store, opts snapshottypes.SnapshotOptions) {
	if app.sealed {
		panic("SetSnapshot() on sealed BaseApp")
	}
	if snapshotStore == nil {
		app.snapshotManager = nil
		return
	}
	app.cms.SetSnapshotInterval(opts.Interval)
	app.snapshotManager = snapshots.NewManager(snapshotStore, opts, app.cms, nil, app.logger)
}

// SetInterfaceRegistry sets the InterfaceRegistry.
func (app *BaseApp) SetInterfaceRegistry(registry types.InterfaceRegistry) {
	app.interfaceRegistry = registry
	app.grpcQueryRouter.SetInterfaceRegistry(registry)
	app.msgServiceRouter.SetInterfaceRegistry(registry)
}

// SetStreamingService is used to set a streaming service into the BaseApp hooks and load the listeners into the multistore
func (app *BaseApp) SetStreamingService(
	appOpts servertypes.AppOptions,
	appCodec storetypes.Codec,
	keys map[string]*storetypes.KVStoreKey,
) error {
	homePath := cast.ToString(appOpts.Get(flags.FlagHome))
	streamers, _, err := streaming.LoadStreamingServices(appOpts, appCodec, app.logger, keys, homePath)
	if err != nil {
		return err
	}
	// add the listeners for each StoreKey
	for _, streamer := range streamers {
		for key, lis := range streamer.Listeners() {
			app.cms.AddListeners(key, lis)
		}
		// register the StreamingService within the BaseApp
		// BaseApp will pass BeginBlock, DeliverTx, and EndBlock requests and responses to the streaming services to update their ABCI context
		app.abciListeners = append(app.abciListeners, streamer)
	}
	return nil
}

// SetTxDecoder sets the TxDecoder if it wasn't provided in the BaseApp constructor.
func (app *BaseApp) SetTxDecoder(txDecoder sdk.TxDecoder) {
	app.txDecoder = txDecoder
}

// SetTxEncoder sets the TxEncoder if it wasn't provided in the BaseApp constructor.
func (app *BaseApp) SetTxEncoder(txEncoder sdk.TxEncoder) {
	app.txEncoder = txEncoder
}

// SetQueryMultiStore set a alternative MultiStore implementation to support grpc query service.
//
// Ref: https://github.com/verzth/cosmos-sdk/issues/13317
func (app *BaseApp) SetQueryMultiStore(ms storetypes.MultiStore) {
	app.qms = ms
}

// SetMempool sets the mempool for the BaseApp and is required for the app to start up.
func (app *BaseApp) SetMempool(mempool mempool.Mempool) {
	if app.sealed {
		panic("SetMempool() on sealed BaseApp")
	}
	app.mempool = mempool
}

// SetProcessProposal sets the process proposal function for the BaseApp.
func (app *BaseApp) SetProcessProposal(handler sdk.ProcessProposalHandler) {
	if app.sealed {
		panic("SetProcessProposal() on sealed BaseApp")
	}
	app.processProposal = handler
}

// SetPrepareProposal sets the prepare proposal function for the BaseApp.
func (app *BaseApp) SetPrepareProposal(handler sdk.PrepareProposalHandler) {
	if app.sealed {
		panic("SetPrepareProposal() on sealed BaseApp")
	}

	app.prepareProposal = handler
}

// SetStoreMetrics sets the prepare proposal function for the BaseApp.
func (app *BaseApp) SetStoreMetrics(gatherer metrics.StoreMetrics) {
	if app.sealed {
		panic("SetStoreMetrics() on sealed BaseApp")
	}

	app.cms.SetMetrics(gatherer)
}
