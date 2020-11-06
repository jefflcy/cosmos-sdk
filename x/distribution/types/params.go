package types

import (
	"fmt"

	"cosmossdk.io/math"
	"sigs.k8s.io/yaml"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultParams returns default distribution parameters
func DefaultParams() Params {
	return Params{
		CommunityTax:            sdk.NewDecWithPrec(2, 2), // 2%
		BaseProposerReward:      sdk.NewDecWithPrec(1, 2), // 1% deprecated
		BonusProposerReward:     sdk.NewDecWithPrec(4, 2), // 4% deprecated
		LiquidityProviderReward: sdk.ZeroDec(),            // 0%
		WithdrawAddrEnabled:     true,
	}
}

func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// ValidateBasic performs basic validation on distribution parameters.
func (p Params) ValidateBasic() error {
	if p.CommunityTax.IsNegative() {
		return fmt.Errorf(
			"community tax should be positive: %s", p.CommunityTax,
		)
	}
	if p.LiquidityProviderReward.IsNegative() {
		return fmt.Errorf(
			"liquidity provider reward should be positive: %s", p.CommunityTax,
		)
	}

	return nil
}

func validateCommunityTax(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("community tax must be not nil")
	}
	if v.IsNegative() {
		return fmt.Errorf("community tax must be positive: %s", v)
	}
	if v.GT(math.LegacyOneDec()) {
		return fmt.Errorf("community tax too large: %s", v)
	}

	return nil
}

func validateLiquidityProviderReward(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("liquidity provider reward must be not nil")
	}
	if v.IsNegative() {
		return fmt.Errorf("liquidity provider reward must be positive: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("liquidity provider reward too large: %s", v)
	}

	return nil
}

func validateWithdrawAddrEnabled(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}
