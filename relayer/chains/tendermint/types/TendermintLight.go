package types

import (
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
)

func NewHeightFromTm(h clienttypes.Height) *Height {
	return &Height{
		RevisionNumber: h.RevisionNumber,
		RevisionHeight: h.RevisionHeight,
	}
}
