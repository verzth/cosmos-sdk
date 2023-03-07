package store

import (
	dbm "github.com/cosmos/cosmos-db"

	"github.com/verzth/cosmos-sdk/log"
	"github.com/verzth/cosmos-sdk/store/cache"
	"github.com/verzth/cosmos-sdk/store/metrics"
	"github.com/verzth/cosmos-sdk/store/rootmulti"
	"github.com/verzth/cosmos-sdk/store/types"
)

func NewCommitMultiStore(db dbm.DB, logger log.Logger, metricGatherer metrics.StoreMetrics) types.CommitMultiStore {
	return rootmulti.NewStore(db, logger, metricGatherer)
}

func NewCommitKVStoreCacheManager() types.MultiStorePersistentCache {
	return cache.NewCommitKVStoreCacheManager(cache.DefaultCommitKVStoreCacheSize)
}
