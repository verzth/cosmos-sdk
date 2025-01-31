package server

import (
	"github.com/verzth/cosmos-sdk/log"
	cmtlog "github.com/cometbft/cometbft/libs/log"
)

var _ cmtlog.Logger = (*CometZeroLogWrapper)(nil)

// CometZeroLogWrapper provides a wrapper around a zerolog.Logger instance.
// It implements CometBFT's Logger interface.
type CometZeroLogWrapper struct {
	log.Logger
}

// With returns a new wrapped logger with additional context provided by a set
// of key/value tuples. The number of tuples must be even and the key of the
// tuple must be a string.
func (cmt CometZeroLogWrapper) With(keyVals ...interface{}) cmtlog.Logger {
	logger := cmt.Logger.With(keyVals...)
	return CometZeroLogWrapper{logger}
}
