package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgLikePost
// ------------------------------------------------------------------------------
var _ sdk.Msg = &MsgLikePost{}

// MsgLikePost ...
type MsgLikePost struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	PostID  string         `json:"post_id" yaml:"post_id"`
}

// NewMsgLikePost ...
func NewMsgLikePost(creator sdk.AccAddress, postID string) MsgLikePost {
	return MsgLikePost{
		Creator: creator,
		PostID:  postID,
	}
}

// LikePostConst ...
const LikePostConst = "LikePost"

// Route ...
func (msg MsgLikePost) Route() string { return RouterKey }

// Type ...
func (msg MsgLikePost) Type() string { return LikePostConst }

// GetSigners ...
func (msg MsgLikePost) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

// GetSignBytes ...
func (msg MsgLikePost) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic ...
func (msg MsgLikePost) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
