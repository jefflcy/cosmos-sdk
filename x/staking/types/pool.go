package types

<<<<<<< HEAD
import (
	"cosmossdk.io/math"
)
=======
import "cosmossdk.io/math"
>>>>>>> v0.46.13-patch

// names used as root for pool module accounts:
//
// - NotBondedPool -> "not_bonded_tokens_pool"
//
// - BondedPool -> "bonded_tokens_pool"
const (
	NotBondedPoolName = "not_bonded_tokens_pool"
	BondedPoolName    = "bonded_tokens_pool"
)

// NewPool creates a new Pool instance used for queries
func NewPool(notBonded, bonded math.Int) Pool {
	return Pool{
		NotBondedTokens: notBonded,
		BondedTokens:    bonded,
	}
}
