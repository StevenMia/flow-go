package model

import (
	"github.com/dapperlabs/flow-go/crypto"
	"github.com/dapperlabs/flow-go/model/flow"
)

type Vote struct {
	View     uint64
	BlockID  flow.Identifier
	SignerID flow.Identifier
	SigData  []byte
}

func (uv *Vote) ID() flow.Identifier {
	return flow.MakeID(uv)
}

// VoteFromFlow turns the vote parameters into a vote struct.
func VoteFromFlow(signerID flow.Identifier, blockID flow.Identifier, view uint64, sig crypto.Signature) *Vote {
	vote := Vote{
		View:     view,
		BlockID:  blockID,
		SignerID: signerID,
		SigData:  sig,
	}
	return &vote
}
