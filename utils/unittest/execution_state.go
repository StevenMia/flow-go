package unittest

import (
	"encoding/hex"

	"github.com/dapperlabs/flow-go/model/flow"
)

// Used below with random root key
//privateKey := flow.AccountPrivateKey{
//	PrivateKey: rootKey,
//	SignAlgo:   crypto.ECDSAP256,
//	HashAlgo:   hash.SHA2_256,
//}

const RootAccountPrivateKeyHex = "e3a08ae3d0461cfed6d6f49bfc25fa899351c39d1bd21fdba8c87595b6c49bb4cc430201"

// Pre-calculated state commitment with root account with the above private key
const GenesisStateCommitmentHex = "03674d6f4d362edd3407f44962b6b2d5901a3121452a5804f26031d500458c79"

var GenesisStateCommitment flow.StateCommitment
var RootAccountPrivateKey flow.AccountPrivateKey
var RootAccountPublicKey flow.AccountPublicKey

func init() {
	var err error
	GenesisStateCommitment, err = hex.DecodeString(GenesisStateCommitmentHex)
	if err != nil {
		panic("error while hex decoding hardcoded state commitment")
	}

	rootAccountPrivateKeyBytes, err := hex.DecodeString(RootAccountPrivateKeyHex)
	if err != nil {
		panic("error while hex decoding hardcoded root key")
	}

	RootAccountPrivateKey, err = flow.DecodeAccountPrivateKey(rootAccountPrivateKeyBytes)
	if err != nil {
		panic("error while decoding hardcoded root key bytes")
	}

	// Cannot import virtual machine, due to circular dependency. Just use the value of
	// virtualmachine.AccountKeyWeightThreshold here
	RootAccountPublicKey = RootAccountPrivateKey.PublicKey(1000)
}
