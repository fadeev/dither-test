package dither

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
	"github.com/virgo-project/dither/x/dither/types"
)

// NewHandler creates an sdk.Handler for all the dither type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgCreatePost:
			return handleMsgCreatePost(ctx, k, msg)
		case MsgLikePost:
			return handleMsgLikePost(ctx, k, msg)
		case MsgCreateChannel:
			return handleMsgCreateChannel(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// handleMsgCreatePost ...
func handleMsgCreatePost(ctx sdk.Context, k Keeper, msg MsgCreatePost) (*sdk.Result, error) {
	var post = types.Post{
		Creator:   msg.Creator,
		Body:      msg.Body,
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
		Likes:     []sdk.AccAddress{},
		ChannelID: msg.ChannelID,
	}
	k.SetPost(ctx, post)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeCreatePost),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator.String()),
			sdk.NewAttribute(types.AttributeBody, msg.Body),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

// handleMsgLikePost ...
func handleMsgLikePost(ctx sdk.Context, k Keeper, msg MsgLikePost) (*sdk.Result, error) {
	var post types.Post
	post, _ = k.GetPost(ctx, msg.PostID)
	post.Likes = append(post.Likes, msg.Creator)
	k.SetPost(ctx, post)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeCreatePost),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator.String()),
			// sdk.NewAttribute(sdk.AttributeKeyPostID, msg.PostID.String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

// handleMsgCreateChannel ...
func handleMsgCreateChannel(ctx sdk.Context, k Keeper, msg MsgCreateChannel) (*sdk.Result, error) {
	var channel = types.Channel{
		Creator: msg.Creator,
		Name:    msg.Name,
		ID:      uuid.New().String(),
	}
	k.SetChannel(ctx, channel)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeCreateChannel),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator.String()),
			sdk.NewAttribute(types.AttributeName, msg.Name),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
