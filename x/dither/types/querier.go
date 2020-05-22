package types

import (
	"strings"
)

// nolint
const (
	QueryListPosts    = "list-posts"
	QueryListChannels = "list-channels"
	QueryLikePost     = "like"
)

// QueryResPosts ...
type QueryResPosts []string

// implement fmt.Stringer
func (n QueryResPosts) String() string {
	return strings.Join(n[:], "\n")
}
