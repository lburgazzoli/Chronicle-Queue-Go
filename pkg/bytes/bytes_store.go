package bytes

// **************************************
//
//
//
// **************************************

// New --
func New() Store {
	return Store{}
}

// **************************************
//
//
//
// **************************************

// Store is an immutable reference to some bytes with fixed extents.
type Store struct {
}
