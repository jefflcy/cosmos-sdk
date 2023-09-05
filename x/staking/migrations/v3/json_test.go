<<<<<<< HEAD:x/staking/migrations/v3/json_test.go
package v3_test
=======
package v046_test
>>>>>>> v0.46.13-patch:x/staking/migrations/v046/json_test.go

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/client"
<<<<<<< HEAD:x/staking/migrations/v3/json_test.go
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	v3 "github.com/cosmos/cosmos-sdk/x/staking/migrations/v3"
=======
	"github.com/cosmos/cosmos-sdk/simapp"
	v046 "github.com/cosmos/cosmos-sdk/x/staking/migrations/v046"
>>>>>>> v0.46.13-patch:x/staking/migrations/v046/json_test.go
	"github.com/cosmos/cosmos-sdk/x/staking/types"
)

func TestMigrateJSON(t *testing.T) {
<<<<<<< HEAD:x/staking/migrations/v3/json_test.go
	encodingConfig := moduletestutil.MakeTestEncodingConfig()
=======
	encodingConfig := simapp.MakeTestEncodingConfig()
>>>>>>> v0.46.13-patch:x/staking/migrations/v046/json_test.go
	clientCtx := client.Context{}.
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithCodec(encodingConfig.Codec)

	oldState := types.DefaultGenesisState()

<<<<<<< HEAD:x/staking/migrations/v3/json_test.go
	newState, err := v3.MigrateJSON(*oldState)
=======
	newState, err := v046.MigrateJSON(*oldState)
>>>>>>> v0.46.13-patch:x/staking/migrations/v046/json_test.go
	require.NoError(t, err)

	bz, err := clientCtx.Codec.MarshalJSON(&newState)
	require.NoError(t, err)

	// Indent the JSON bz correctly.
	var jsonObj map[string]interface{}
	err = json.Unmarshal(bz, &jsonObj)
	require.NoError(t, err)
	indentedBz, err := json.MarshalIndent(jsonObj, "", "\t")
	require.NoError(t, err)

	// Make sure about new param MinCommissionRate.
	expected := `{
	"delegations": [],
	"exported": false,
	"last_total_power": "0",
	"last_validator_powers": [],
	"params": {
		"bond_denom": "stake",
		"historical_entries": 10000,
		"max_entries": 7,
		"max_validators": 100,
		"min_commission_rate": "0.000000000000000000",
		"unbonding_time": "1814400s"
	},
	"redelegations": [],
	"unbonding_delegations": [],
	"validators": []
}`

	require.Equal(t, expected, string(indentedBz))
}
