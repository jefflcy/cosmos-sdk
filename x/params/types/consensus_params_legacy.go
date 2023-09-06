package types

import (
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
)

// Deprecated.
func ConsensusParamsKeyTable() KeyTable {
	return NewKeyTable(
		NewParamSetPair(
			baseapp.ParamStoreKeyBlockParams, abci.BlockParams{}, baseapp.ValidateBlockParams,
		),
		NewParamSetPair(
			baseapp.ParamStoreKeyEvidenceParams, tmproto.EvidenceParams{}, baseapp.ValidateEvidenceParams,
		),
		NewParamSetPair(
			baseapp.ParamStoreKeyValidatorParams, tmproto.ValidatorParams{}, baseapp.ValidateValidatorParams,
		),
	)
}
