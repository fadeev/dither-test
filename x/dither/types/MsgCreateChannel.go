package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgCreateChannel
// ------------------------------------------------------------------------------
var _ sdk.Msg = &MsgCreateChannel{}

// MsgCreateChannel ...
type MsgCreateChannel struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Name    string         `json:"name" yaml:"name"`
}

// NewMsgCreateChannel ...
func NewMsgCreateChannel(creator sdk.AccAddress, name string) MsgCreateChannel {
	return MsgCreateChannel{
		Creator: creator,
		Name:    name,
	}
}

// CreateChannelConst ...
const CreateChannelConst = "CreateChannel"

// Route ...
func (msg MsgCreateChannel) Route() string { return RouterKey }

// Type ...
func (msg MsgCreateChannel) Type() string { return CreateChannelConst }

// GetSigners ...
func (msg MsgCreateChannel) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

// GetSignBytes ...
func (msg MsgCreateChannel) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic ...
func (msg MsgCreateChannel) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
