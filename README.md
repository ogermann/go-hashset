#go-hashSet
A minimalistic hashSet implementation in Go.

Every entry must fulfill a specific Comparable interface:
```
type Comparable interface {
	HashCode() int
	Equals(obj interface{}) bool
}
```
Example:
```
type testComparable struct {
	id int
}

func (t *testComparable) Equals(obj interface{}) bool {
	cast, ok := obj.(*testComparable)
	if !ok {
		return false
	}
	return t.id == cast.id
}

func (t *testComparable) HashCode() int {
	return t.id
}

func main() {
	hashSet := NewHashSet()
	hashSet.Add(&testComparable{1})
	println(hashSet.Contains(&testComparable{1})) // true
}
```