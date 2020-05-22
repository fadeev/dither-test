package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgCreatePost
// ------------------------------------------------------------------------------
var _ sdk.Msg = &MsgCreatePost{}

// MsgCreatePost ...
type MsgCreatePost struct {
	Creator   sdk.AccAddress   `json:"creator" yaml:"creator"`
	Body      string           `json:"body" yaml:"body"`
	Likes     []sdk.AccAddress `json:"likes" yaml:"likes"`
	ChannelID string           `json:"channel_id" yaml:"channel_id"`
}

// NewMsgCreatePost ...
func NewMsgCreatePost(creator sdk.AccAddress, body string, channelID string) MsgCreatePost {
	return MsgCreatePost{
		Creator:   creator,
		Body:      body,
		ChannelID: channelID,
	}
}

// CreatePostConst ...
const CreatePostConst = "CreatePost"

// Route ...
func (msg MsgCreatePost) Route() string { return RouterKey }

// Type ...
func (msg MsgCreatePost) Type() string { return CreatePostConst }

// GetSigners ...
func (msg MsgCreatePost) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

// GetSignBytes ...
func (msg MsgCreatePost) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic ...
func (msg MsgCreatePost) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
