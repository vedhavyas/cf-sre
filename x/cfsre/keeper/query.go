package keeper

import (
	"github.com/vedhavyas/cf-sre/x/cfsre/types"
)

var _ types.QueryServer = Keeper{}
