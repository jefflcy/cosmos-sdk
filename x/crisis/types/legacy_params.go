package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

<<<<<<< HEAD:x/crisis/types/legacy_params.go
// ParamStoreKeyConstantFee is the constant fee parameter
=======
// ParamStoreKeyConstantFee is the key for constant fee parameter
>>>>>>> v0.46.13-patch:x/crisis/types/params.go
var ParamStoreKeyConstantFee = []byte("ConstantFee")

// Deprecated: Type declaration for parameters
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable(
		paramtypes.NewParamSetPair(ParamStoreKeyConstantFee, sdk.Coin{}, validateConstantFee),
	)
}

func validateConstantFee(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if !v.IsValid() {
		return fmt.Errorf("invalid constant fee: %s", v)
	}

	return nil
}
