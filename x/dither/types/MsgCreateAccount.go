package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgCreateAccount
// ------------------------------------------------------------------------------
var _ sdk.Msg = &MsgCreateAccount{}

// MsgCreateAccount ...
type MsgCreateAccount struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Name    string         `json:"name" yaml:"name"`
}

// NewMsgCreateAccount ...
func NewMsgCreateAccount(creator sdk.AccAddress, name string) MsgCreateAccount {
	return MsgCreateAccount{
		Creator: creator,
		Name:    name,
	}
}

// CreateAccountConst ...
const CreateAccountConst = "CreateAccount"

// Route ...
func (msg MsgCreateAccount) Route() string { return RouterKey }

// Type ...
func (msg MsgCreateAccount) Type() string { return CreateAccountConst }

// GetSigners ...
func (msg MsgCreateAccount) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

// GetSignBytes ...
func (msg MsgCreateAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic ...
func (msg MsgCreateAccount) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
