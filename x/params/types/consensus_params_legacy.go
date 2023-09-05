package types

import (
<<<<<<< HEAD:x/params/types/consensus_params_legacy.go
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
=======
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
>>>>>>> v0.46.13-patch:x/params/types/consensus_params.go

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
