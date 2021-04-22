package hashSet

import (
	"fmt"
	"sync"
)

// A HashSet is a data structure holding any type fulfilling the Comparable interface.
// Every Comparable struct can exists only once in the HashSet.
// The performance of this data structure is highly influenced by the Comparable#HashCode() method.
type HashSet interface {
	// Add adds a new Comparable to the HashSet, as long as there's no equal Comparable
	// present in the HashSet. An equal Comparable means a comparable that returns true at their
	// Comparable#Equals(obj interface{}) method.
	// An ErrEntryExistsAlready is being returned if there's already another Comparable existing
	// in the HashSet.
	Add(entry Comparable) error
	// Remove removes a Comparable from the HashSet if its Equals method return true.
	// An ErrEntryDoesNotExist is being returned if no Comparable can be found.
	Remove(entry Comparable) error
	// Contains returns a boolean which is true if a Comparable with a truthy Equals method exist or
	// false if it doesn't.
	Contains(entry Comparable) bool
	// IsEmpty returns a boolean determining if the HashSet is empty or not
	IsEmpty() bool
	// Size returns an int determining the size of the HashSet
	Size() int
	// ToSlice returns a slice containing all entries of the HashSet
	ToSlice() []Comparable
}

var ErrEntryDoesNotExist = fmt.Errorf("entry does not exist in set")
var ErrEntryExistsAlready = fmt.Errorf("entry does already exist in set")

type linkedEntry struct {
	data Comparable
	next *linkedEntry
}

type hashSet struct {
	sync.RWMutex
	set map[int]*linkedEntry
}

func NewHashSet(entries ...Comparable) HashSet {
	set := &hashSet{
		set: make(map[int]*linkedEntry),
	}
	for _, entry := range entries {
		_ = set.Add(entry)
	}
	return set
}

func (set *hashSet) Add(entry Comparable) error {
	hashCode := entry.HashCode()
	set.Lock()
	defer set.Unlock()
	existingEntry, ok := set.set[hashCode]
	if !ok {
		set.set[hashCode] = &linkedEntry{data: entry}
		return nil
	}

	for {
		if existingEntry.data.Equals(entry) {
			return ErrEntryExistsAlready
		}
		if existingEntry.next == nil {
			existingEntry.next = &linkedEntry{data: entry}
			return nil
		}
		existingEntry = existingEntry.next
	}
}

func (set *hashSet) Remove(entry Comparable) error {
	hashCode := entry.HashCode()
	set.Lock()
	defer set.Unlock()
	existingEntry, ok := set.set[hashCode]
	if !ok {
		return ErrEntryDoesNotExist
	}

	if existingEntry.data.Equals(entry) {
		if existingEntry.next == nil {
			delete(set.set, hashCode)
			return nil
		} else {
			set.set[hashCode] = existingEntry.next
			return nil
		}
	}

	for {
		if existingEntry.next.data.Equals(entry) {
			existingEntry.next = existingEntry.next.next
			return nil
		}
		if existingEntry.next != nil {
			existingEntry = existingEntry.next
		} else {
			break
		}
	}

	return ErrEntryDoesNotExist
}

func (set *hashSet) Contains(entry Comparable) bool {
	hashCode := entry.HashCode()
	set.RLock()
	defer set.RUnlock()
	existingEntry, ok := set.set[hashCode]
	if !ok {
		return false
	}

	for {
		if existingEntry.data.Equals(entry) {
			return true
		}
		if existingEntry.next == nil {
			return false
		}
		existingEntry = existingEntry.next
	}
}

func (set *hashSet) IsEmpty() bool {
	set.RLock()
	defer set.RUnlock()
	return len(set.set) <= 0
}

func (set *hashSet) Size() int {
	set.RLock()
	defer set.RUnlock()
	size := 0
	for _, entry := range set.set {
		for {
			size++
			if entry.next == nil {
				break
			}
			entry = entry.next
		}
	}
	return size
}

func (set *hashSet) ToSlice() []Comparable {
	var all []Comparable
	set.RLock()
	defer set.RUnlock()
	for _, entry := range set.set {
		for {
			all = append(all, entry.data)
			if entry.next == nil {
				break
			}
			entry = entry.next
		}
	}
	return all
}