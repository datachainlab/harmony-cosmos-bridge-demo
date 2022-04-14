package types

import (
	//crypto "github.com/tendermint/tendermint/proto/tendermint/crypto"
	crypto "github.com/tendermint/tendermint/crypto/ed25519"
	tmtypes "github.com/tendermint/tendermint/types"
)

func (v *Validator) Validator() *tmtypes.Validator {
	pubkey := crypto.PubKey(v.PubKey)
	return &tmtypes.Validator{
		Address:     pubkey.Address(),
		PubKey:      pubkey,
		VotingPower: v.VotingPower,
		// TODO
		ProposerPriority: 0,
	}
}

func (v *ValidatorSet) ValidatorSet() *tmtypes.ValidatorSet {
	if v.Validators == nil {
		return nil
	}
	vals := make([]*tmtypes.Validator, len(v.Validators))
	for i, val := range v.Validators {
		vals[i] = val.Validator()
	}
	vset := tmtypes.NewValidatorSet(vals)
	vset.GetProposer()
	return vset
}

// TODO tmproto or tmtypes
func NewValidatorFromTm(v *tmtypes.Validator) *Validator {
	return &Validator{
		PubKey:      v.PubKey.Bytes(),
		VotingPower: v.VotingPower,
	}
}

func NewValidatorSetFromTm(vs *tmtypes.ValidatorSet) *ValidatorSet {
	vals := make([]*Validator, len(vs.Validators))
	for i, v := range vs.Validators {
		vals[i] = NewValidatorFromTm(v)
	}
	return &ValidatorSet{
		Validators:       vals,
		TotalVotingPower: vs.TotalVotingPower(),
	}
}
