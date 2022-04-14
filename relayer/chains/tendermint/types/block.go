package types

import (
	tmtypes "github.com/tendermint/tendermint/types"
)

func (b *CanonicalBlockID) BlockID() *tmtypes.BlockID {
	return &tmtypes.BlockID{
		Hash: b.Hash,
		PartSetHeader: tmtypes.PartSetHeader{
			Total: b.PartSetHeader.Total,
			Hash:  b.PartSetHeader.Hash,
		},
	}
}

func (c *Commit) Commit() *tmtypes.Commit {
	sigs := make([]tmtypes.CommitSig, len(c.Signatures))
	for i, sig := range c.Signatures {
		sigs[i] = *sig.CommitSig()
	}

	return tmtypes.NewCommit(
		c.Height,
		c.Round,
		*c.BlockId.BlockID(),
		sigs,
	)
}

func (c *CommitSig) CommitSig() *tmtypes.CommitSig {
	cs := &tmtypes.CommitSig{
		BlockIDFlag:      tmtypes.BlockIDFlag(c.BlockIdFlag),
		ValidatorAddress: c.ValidatorAddress,
		Signature:        c.Signature,
	}
	if c.Timestamp != nil {
		cs.Timestamp = c.Timestamp.Time()
	}
	return cs
}

func NewCanonicalBlockIDFromTm(id *tmtypes.BlockID) *CanonicalBlockID {
	return &CanonicalBlockID{
		Hash:          id.Hash,
		PartSetHeader: NewPartSetHeaderFromTm(&id.PartSetHeader),
	}
}

func NewCommitFromTm(c *tmtypes.Commit) *Commit {
	sigs := make([]*CommitSig, len(c.Signatures))
	for i, sig := range c.Signatures {
		sigs[i] = NewCommitFromTmSig(&sig)
	}
	return &Commit{
		Height:     c.Height,
		Round:      c.Round,
		BlockId:    NewCanonicalBlockIDFromTm(&c.BlockID),
		Signatures: sigs,
	}
}

func NewCommitFromTmSig(cs *tmtypes.CommitSig) *CommitSig {
	return &CommitSig{
		BlockIdFlag:      BlockIDFlag(cs.BlockIDFlag),
		ValidatorAddress: cs.ValidatorAddress.Bytes(),
		Timestamp:        NewTimestampFromTime(cs.Timestamp),
		Signature:        cs.Signature,
	}
}

func NewPartSetHeaderFromTm(h *tmtypes.PartSetHeader) *CanonicalPartSetHeader {
	return &CanonicalPartSetHeader{
		Total: h.Total,
		Hash:  h.Hash,
	}
}
