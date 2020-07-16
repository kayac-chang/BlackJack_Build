package lobby

const (

	// RejectLatter indicates that a lobby should reject the player
	// who had already logged-in with same ID.
	RejectLatter policy = iota

	// RemoveFormer indicates that a lobby should kick out the player,
	// replacing it with the newer one who logs-in with same ID.
	RemoveFormer
)

// A policy implements the Policy interface, preventing external
// packages from customizing Policies.
type policy int

func (p policy) c() policy {
	return p
}

// A Policy defines the behavior of a lobby.
type Policy interface {
	c() policy
}
