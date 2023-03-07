package types

import "github.com/verzth/cosmos-sdk/errors"

// x/staking module sentinel errors
//
// TODO: Many of these errors are redundant. They should be removed and replaced
// by sdkerrors.ErrInvalidRequest.
//
// REF: https://github.com/verzth/cosmos-sdk/issues/5450
var (
	ErrEmptyValidatorAddr              = errors.Register(ModuleName, 2, "empty validator address")
	ErrNoValidatorFound                = errors.Register(ModuleName, 3, "validator does not exist")
	ErrValidatorOwnerExists            = errors.Register(ModuleName, 4, "validator already exist for this operator address; must use new validator operator address")
	ErrValidatorPubKeyExists           = errors.Register(ModuleName, 5, "validator already exist for this pubkey; must use new validator pubkey")
	ErrValidatorPubKeyTypeNotSupported = errors.Register(ModuleName, 6, "validator pubkey type is not supported")
	ErrValidatorJailed                 = errors.Register(ModuleName, 7, "validator for this address is currently jailed")
	ErrBadRemoveValidator              = errors.Register(ModuleName, 8, "failed to remove validator")
	ErrCommissionNegative              = errors.Register(ModuleName, 9, "commission must be positive")
	ErrCommissionHuge                  = errors.Register(ModuleName, 10, "commission cannot be more than 100%")
	ErrCommissionGTMaxRate             = errors.Register(ModuleName, 11, "commission cannot be more than the max rate")
	ErrCommissionUpdateTime            = errors.Register(ModuleName, 12, "commission cannot be changed more than once in 24h")
	ErrCommissionChangeRateNegative    = errors.Register(ModuleName, 13, "commission change rate must be positive")
	ErrCommissionChangeRateGTMaxRate   = errors.Register(ModuleName, 14, "commission change rate cannot be more than the max rate")
	ErrCommissionGTMaxChangeRate       = errors.Register(ModuleName, 15, "commission cannot be changed more than max change rate")
	ErrSelfDelegationBelowMinimum      = errors.Register(ModuleName, 16, "validator's self delegation must be greater than their minimum self delegation")
	ErrMinSelfDelegationDecreased      = errors.Register(ModuleName, 17, "minimum self delegation cannot be decrease")
	ErrEmptyDelegatorAddr              = errors.Register(ModuleName, 18, "empty delegator address")
	ErrNoDelegation                    = errors.Register(ModuleName, 19, "no delegation for (address, validator) tuple")
	ErrBadDelegatorAddr                = errors.Register(ModuleName, 20, "delegator does not exist with address")
	ErrNoDelegatorForAddress           = errors.Register(ModuleName, 21, "delegator does not contain delegation")
	ErrInsufficientShares              = errors.Register(ModuleName, 22, "insufficient delegation shares")
	ErrDelegationValidatorEmpty        = errors.Register(ModuleName, 23, "cannot delegate to an empty validator")
	ErrNotEnoughDelegationShares       = errors.Register(ModuleName, 24, "not enough delegation shares")
	ErrNotMature                       = errors.Register(ModuleName, 25, "entry not mature")
	ErrNoUnbondingDelegation           = errors.Register(ModuleName, 26, "no unbonding delegation found")
	ErrMaxUnbondingDelegationEntries   = errors.Register(ModuleName, 27, "too many unbonding delegation entries for (delegator, validator) tuple")
	ErrNoRedelegation                  = errors.Register(ModuleName, 28, "no redelegation found")
	ErrSelfRedelegation                = errors.Register(ModuleName, 29, "cannot redelegate to the same validator")
	ErrTinyRedelegationAmount          = errors.Register(ModuleName, 30, "too few tokens to redelegate (truncates to zero tokens)")
	ErrBadRedelegationDst              = errors.Register(ModuleName, 31, "redelegation destination validator not found")
	ErrTransitiveRedelegation          = errors.Register(ModuleName, 32, "redelegation to this validator already in progress; first redelegation to this validator must complete before next redelegation")
	ErrMaxRedelegationEntries          = errors.Register(ModuleName, 33, "too many redelegation entries for (delegator, src-validator, dst-validator) tuple")
	ErrDelegatorShareExRateInvalid     = errors.Register(ModuleName, 34, "cannot delegate to validators with invalid (zero) ex-rate")
	ErrBothShareMsgsGiven              = errors.Register(ModuleName, 35, "both shares amount and shares percent provided")
	ErrNeitherShareMsgsGiven           = errors.Register(ModuleName, 36, "neither shares amount nor shares percent provided")
	ErrInvalidHistoricalInfo           = errors.Register(ModuleName, 37, "invalid historical info")
	ErrNoHistoricalInfo                = errors.Register(ModuleName, 38, "no historical info found")
	ErrEmptyValidatorPubKey            = errors.Register(ModuleName, 39, "empty validator public key")
	ErrCommissionLTMinRate             = errors.Register(ModuleName, 40, "commission cannot be less than min rate")
	ErrUnbondingNotFound               = errors.Register(ModuleName, 41, "unbonding operation not found")
	ErrUnbondingOnHoldRefCountNegative = errors.Register(ModuleName, 42, "cannot un-hold unbonding operation that is not on hold")
)
