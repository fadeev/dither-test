package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/virgo-project/dither/x/dither/types"
)

// Keeper ...
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
	CoinKeeper bank.Keeper
}

// NewKeeper ...
func NewKeeper(coinKeeper bank.Keeper, cdc *codec.Codec, key sdk.StoreKey) Keeper {
	keeper := Keeper{
		CoinKeeper: coinKeeper,
		storeKey:   key,
		cdc:        cdc,
	}
	return keeper
}

// Logger ...
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetPost ...
func (k Keeper) SetPost(ctx sdk.Context, post types.Post) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(post)
	key := []byte(types.PostPrefix + post.ID)
	store.Set(key, bz)
}

// GetPost ...
func (k Keeper) GetPost(ctx sdk.Context, postID string) (types.Post, error) {
	store := ctx.KVStore(k.storeKey)
	var post types.Post
	byteKey := []byte(types.PostPrefix + postID)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &post)
	if err != nil {
		return post, err
	}
	return post, nil
}

// GetPostsIterator ...
func (k Keeper) GetPostsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte(types.PostPrefix))
}

// SetChannel ...
func (k Keeper) SetChannel(ctx sdk.Context, channel types.Channel) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(channel)
	key := []byte(types.ChannelPrefix + channel.ID)
	store.Set(key, bz)
}
