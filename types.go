package hashSet

// Comparable this interface is used to efficiently compare structs. As HashCode are just integer
// values they can clash, if they do the Equals method is used.
type Comparable interface {
	// HashCode returns a hashCode as int. The hashCode has to be identical for every invocation
	// with the same parameters
	HashCode() int
	// Equals compares an another to this struct and returns a boolean telling if these structs are
	// equal or not
	Equals(obj interface{}) bool
}
