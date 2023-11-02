package keeper

import (
	"loan/x/loan/types"
)

var _ types.QueryServer = Keeper{}
