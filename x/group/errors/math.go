package errors

import "github.com/verzth/cosmos-sdk/errors"

// mathCodespace is the codespace for all errors defined in math package
const mathCodespace = "math"

// ErrInvalidDecString defines an error for an invalid decimal string
var ErrInvalidDecString = errors.Register(mathCodespace, 10, "invalid decimal string")
