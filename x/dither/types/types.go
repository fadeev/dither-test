package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Post ...
type Post struct {
	Creator   sdk.AccAddress   `json:"creator" yaml:"creator"`
	Body      string           `json:"body" yaml:"body"`
	ID        string           `json:"id" yaml:"id"`
	Likes     []sdk.AccAddress `json:"likes" yaml:"likes"`
	ChannelID string           `json:"channel_id" yaml:"channel_id"`
	CreatedAt time.Time        `json:"created_at" yaml:"created_at"`
}

// Posts ...
type Posts []Post

// Channel ...
type Channel struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Name    string         `json:"name" yaml:"name"`
	ID      string         `json:"id" yaml:"id"`
}

// Channels ...
type Channels []Channel
