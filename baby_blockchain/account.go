package core

type AccountKind int

const (
	AccountKindIssuer AccountKind = iota
	AccountKindUser
)

type Account struct {
	// In general, it is recommended to use a different DID for each application,
	// rather than using the same DID across multiple applications.
	// pros this way: Security, Privacy, and Portability.
	DID  *DID
	Kind AccountKind
	// We can have different proofs for one DID. For example, we can have a proof
	// about age and another proof about membership of a club.
	proofs []*Proof
}

// Proofs get list of proofs.
func (a *Account) Proofs() []*Proof {
	return a.proofs
}

// AddProof add a new proof to the account.
func (a *Account) AddProof(proof *Proof) {
	a.proofs = append(a.proofs, proof)
}

// NewAccount creates a new account.
func NewAccount(did *DID, kind AccountKind) *Account {
	return &Account{
		DID:  did,
		Kind: kind,
	}
}
