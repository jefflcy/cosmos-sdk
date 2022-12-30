package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// NewAccountWithAddress implements AccountKeeperI.
func (ak AccountKeeper) NewAccountWithAddress(ctx sdk.Context, addr sdk.AccAddress) types.AccountI {
	acc := ak.proto()
	err := acc.SetAddress(addr)
	if err != nil {
		panic(err)
	}

	return ak.NewAccount(ctx, acc)
}

// NewAccount sets the next account number to a given account interface
func (ak AccountKeeper) NewAccount(ctx sdk.Context, acc types.AccountI) types.AccountI {
	if err := acc.SetAccountNumber(ak.GetNextAccountNumber(ctx)); err != nil {
		panic(err)
	}

	return acc
}

// AccountExists implements AccountKeeperI.
// Check if an account exists in the store, includes check for mapping.
func (ak AccountKeeper) AccountExists(ctx sdk.Context, addr sdk.AccAddress) bool {
	if addr == nil {
		return false
	}
	store := ctx.KVStore(ak.key)
	if !store.Has(types.AddressStoreKey(addr)) {
		cosmosAddr := ak.GetCorrespondingCosmosAddressIfExists(ctx, addr)
		if cosmosAddr == nil {
			return false
		}
		return store.Has(types.AddressStoreKey(cosmosAddr))
	}
	return true
}

// HasAccount implements AccountKeeperI.
// Check if an account exists in the store based on address directly, doesn't check for mapping.
func (ak AccountKeeper) HasAccount(ctx sdk.Context, addr sdk.AccAddress) bool {
	store := ctx.KVStore(ak.key)
	return store.Has(types.AddressStoreKey(addr))
}

// HasAccountAddressByID checks account address exists by id.
func (ak AccountKeeper) HasAccountAddressByID(ctx sdk.Context, id uint64) bool {
	store := ctx.KVStore(ak.key)
	return store.Has(types.AccountNumberStoreKey(id))
}

// GetAccount implements AccountKeeperI.
func (ak AccountKeeper) GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI {
	if addr == nil {
		return nil
	}
	store := ctx.KVStore(ak.key)
	bz := store.Get(types.AddressStoreKey(addr))
	if bz == nil {
		cosmosAddr := ak.GetCorrespondingCosmosAddressIfExists(ctx, addr)
		if cosmosAddr == nil {
			return nil
		}
		accBz := store.Get(types.AddressStoreKey(cosmosAddr))
		return ak.decodeAccount(accBz)
	}

	return ak.decodeAccount(bz)
}

// GetAccountAddressById returns account address by id.
func (ak AccountKeeper) GetAccountAddressByID(ctx sdk.Context, id uint64) string {
	store := ctx.KVStore(ak.key)
	bz := store.Get(types.AccountNumberStoreKey(id))
	if bz == nil {
		return ""
	}
	return sdk.AccAddress(bz).String()
}

// GetAllAccounts returns all accounts in the accountKeeper.
func (ak AccountKeeper) GetAllAccounts(ctx sdk.Context) (accounts []types.AccountI) {
	ak.IterateAccounts(ctx, func(acc types.AccountI) (stop bool) {
		accounts = append(accounts, acc)
		return false
	})

	return accounts
}

// SetAccount implements AccountKeeperI.
func (ak AccountKeeper) SetAccount(ctx sdk.Context, acc types.AccountI) {
	addr := acc.GetAddress()
	store := ctx.KVStore(ak.key)

	bz, err := ak.MarshalAccount(acc)
	if err != nil {
		panic(err)
	}

	store.Set(types.AddressStoreKey(addr), bz)
	store.Set(types.AccountNumberStoreKey(acc.GetAccountNumber()), addr.Bytes())
}

// RemoveAccount removes an account for the account mapper store.
// NOTE: this will cause supply invariant violation if called
func (ak AccountKeeper) RemoveAccount(ctx sdk.Context, acc types.AccountI) {
	addr := acc.GetAddress()
	store := ctx.KVStore(ak.key)
	store.Delete(types.AddressStoreKey(addr))
	store.Delete(types.AccountNumberStoreKey(acc.GetAccountNumber()))
}

// IterateAccounts iterates over all the stored accounts and performs a callback function.
// Stops iteration when callback returns true.
func (ak AccountKeeper) IterateAccounts(ctx sdk.Context, cb func(account types.AccountI) (stop bool)) {
	store := ctx.KVStore(ak.key)
	iterator := sdk.KVStorePrefixIterator(store, types.AddressStoreKeyPrefix)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		account := ak.decodeAccount(iterator.Value())

		if cb(account) {
			break
		}
	}
}

func (ak AccountKeeper) GetCorrespondingEthAddressIfExists(ctx sdk.Context, cosmosAddr sdk.AccAddress) (correspondingEthAddr sdk.AccAddress) {
	mapping := ak.Store(ctx, types.CosmosAddressToEthAddressKey)
	return mapping.Get(cosmosAddr)
}

func (ak AccountKeeper) GetCorrespondingCosmosAddressIfExists(ctx sdk.Context, ethAddr sdk.AccAddress) (correspondingCosmosAddr sdk.AccAddress) {
	mapping := ak.Store(ctx, types.EthAddressToCosmosAddressKey)
	return mapping.Get(ethAddr)

}

func (ak AccountKeeper) SetCorrespondingAddresses(ctx sdk.Context, cosmosAddr sdk.AccAddress, ethAddr sdk.AccAddress) {
	ak.AddToEthToCosmosAddressMap(ctx, ethAddr, cosmosAddr)
	ak.AddToCosmosToEthAddressMap(ctx, cosmosAddr, ethAddr)

}

func (ak AccountKeeper) AddToCosmosToEthAddressMap(ctx sdk.Context, cosmosAddr sdk.AccAddress, ethAddr sdk.AccAddress) {
	cosmosAddrToEthAddrMapping := ak.Store(ctx, types.CosmosAddressToEthAddressKey)
	cosmosAddrToEthAddrMapping.Set(cosmosAddr, ethAddr)
}

func (ak AccountKeeper) AddToEthToCosmosAddressMap(ctx sdk.Context, ethAddr sdk.AccAddress, cosmosAddr sdk.AccAddress) {
	ethAddrToCosmosAddrMapping := ak.Store(ctx, types.EthAddressToCosmosAddressKey)
	ethAddrToCosmosAddrMapping.Set(ethAddr, cosmosAddr)
}

func (ak AccountKeeper) IterateEthToCosmosAddressMapping(ctx sdk.Context, cb func(ethAddress, cosmosAddress sdk.AccAddress) bool) {
	store := ctx.KVStore(ak.key)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.EthAddressToCosmosAddressKey))

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		addressKey := make([]byte, len(iterator.Key())-len(types.KeyPrefix(types.EthAddressToCosmosAddressKey)))
		copy(addressKey, iterator.Key()[len(types.KeyPrefix(types.EthAddressToCosmosAddressKey)):])
		if cb(addressKey, iterator.Value()) {
			break
		}
	}

}
func (ak AccountKeeper) IterateCosmosToEthAddressMapping(ctx sdk.Context, cb func(cosmosAddress, ethAddress sdk.AccAddress) bool) {
	store := ctx.KVStore(ak.key)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.CosmosAddressToEthAddressKey))

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		addressKey := make([]byte, len(iterator.Key())-len(types.KeyPrefix(types.CosmosAddressToEthAddressKey)))
		copy(addressKey, iterator.Key()[len(types.KeyPrefix(types.CosmosAddressToEthAddressKey)):])
		if cb(addressKey, iterator.Value()) {
			break
		}
	}
}
