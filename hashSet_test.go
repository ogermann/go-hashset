package hashSet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testComparable struct {
	id int
	sub int
}

func (t *testComparable) Equals(obj interface{}) bool {
	cast, ok := obj.(*testComparable)
	if !ok {
		return false
	}
	return t.id == cast.id && t.sub == cast.sub
}

func (t *testComparable) HashCode() int {
	return t.id
}

func Test_HashSet_Add(t *testing.T) {
	item1 := &testComparable{1,0}
	set := NewHashSet()
	assert.Equal(t,nil, set.Add(item1))
	controlItem, _ := set.(*hashSet).set[item1.HashCode()]
	assert.Equal(t, item1, controlItem.data)
}

func Test_HashSet_AddCollidingHashCodes(t *testing.T) {
	item1 := &testComparable{1,0}
	item2 := &testComparable{1,1}
	set := NewHashSet()
	assert.Equal(t, nil, set.Add(item1))
	assert.Equal(t, nil, set.Add(item2))
	controlItem, _ := set.(*hashSet).set[item1.HashCode()]
	assert.Equal(t, item1, controlItem.data)
	assert.Equal(t, item2, controlItem.next.data)
}

func Test_HashSet_AddDuplicate(t *testing.T) {
	item1 := &testComparable{1,0}
	set := NewHashSet(item1)
	assert.EqualError(t, set.Add(&testComparable{1,0}), ErrEntryExistsAlready.Error())
	controlItem, _ := set.(*hashSet).set[item1.HashCode()]
	assert.Equal(t, item1, controlItem.data)
}

func Test_HashSet_Remove(t *testing.T) {
	set := NewHashSet(&testComparable{1,0})
	assert.Equal(t,nil, set.Remove(&testComparable{1,0}))
	_, ok := set.(*hashSet).set[1]
	assert.False(t, ok)
}

func Test_HashSet_RemoveCollidingHashCodes(t *testing.T) {
	item1 := &testComparable{1,0}
	item2 := &testComparable{1,1}
	set := NewHashSet(item1, item2)
	assert.Equal(t,nil, set.Remove(item1))
	assert.False(t, set.Contains(item1))
	assert.True(t, set.Contains(item2))
}

func Test_HashSet_Contains(t *testing.T) {
	set := NewHashSet(&testComparable{1,0})
	assert.True(t, set.Contains(&testComparable{1,0}))
}

func Test_HashSet_ContainsNot(t *testing.T) {
	set := NewHashSet(&testComparable{1,0})
	assert.False(t, set.Contains(&testComparable{1,1}))
}

func Test_HashSet_CollidingHashCodes(t *testing.T) {
	item1 := &testComparable{1,0}
	item2 := &testComparable{1,1}
	item3 := &testComparable{1,2}
	set := NewHashSet(item1, item2)
	assert.True(t, set.Contains(item1))
	assert.True(t, set.Contains(item2))
	assert.False(t, set.Contains(item3))
}

func Test_HashSet_IsEmptyTrue(t *testing.T) {
	set := NewHashSet()
	assert.True(t, set.IsEmpty())
}

func Test_HashSet_IsEmptyFalse(t *testing.T) {
	item1 := &testComparable{1,0}
	set := NewHashSet(item1)
	assert.False(t, set.IsEmpty())
}

func Test_HashSet_ToSlice(t *testing.T) {
	item1 := &testComparable{1,0}
	set := NewHashSet(item1)
	slice := set.ToSlice()
	assert.Equal(t, 1, len(slice))
	assert.True(t, slice[0].Equals(item1))
}