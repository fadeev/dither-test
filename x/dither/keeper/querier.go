package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/virgo-project/dither/x/dither/types"
)

// NewQuerier ...
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryListPosts:
			return listPosts(ctx, path[1:], req, k)
		case types.QueryListChannels:
			return listChannels(ctx, k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown query endpoint")
		}
	}
}

// RemovePrefixFromHash removes the prefix from the key
func RemovePrefixFromHash(key []byte, prefix []byte) (hash []byte) {
	hash = key[len(prefix):]
	return hash
}

func listPosts(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var postList types.Posts
	store := ctx.KVStore(k.storeKey)
	channelID := path[0]
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.PostPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var post types.Post
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &post)
		if channelID != "" {
			if channelID == post.ChannelID {
				postList = append(postList, post)
			}
		} else {
			postList = append(postList, post)
		}
	}
	res := codec.MustMarshalJSONIndent(k.cdc, postList)
	return res, nil
}

func listChannels(ctx sdk.Context, k Keeper) ([]byte, error) {
	var channelList types.Channels
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.ChannelPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var channel types.Channel
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &channel)
		channelList = append(channelList, channel)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, channelList)
	return res, nil
}
