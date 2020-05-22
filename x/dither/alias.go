package dither

import (
	"github.com/virgo-project/dither/x/dither/keeper"
	"github.com/virgo-project/dither/x/dither/types"
)

const (
	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	DefaultParamspace = types.DefaultParamspace
	QuerierRoute      = types.QuerierRoute
)

var (
	NewKeeper           = keeper.NewKeeper
	NewQuerier          = keeper.NewQuerier
	RegisterCodec       = types.RegisterCodec
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis

	ModuleCdc           = types.ModuleCdc
	NewMsgCreatePost    = types.NewMsgCreatePost
	NewMsgLikePost      = types.NewMsgLikePost
	NewMsgCreateChannel = types.NewMsgCreateChannel
)

type (
	Keeper           = keeper.Keeper
	GenesisState     = types.GenesisState
	Params           = types.Params
	MsgCreatePost    = types.MsgCreatePost
	MsgLikePost      = types.MsgLikePost
	MsgCreateChannel = types.MsgCreateChannel
)
