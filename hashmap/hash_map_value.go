package hashmap

import "crypto/sha1"

type HashMapValue struct {
	hash  [sha1.Size]byte
	value string
}

func newHashMapValue(hash [sha1.Size]byte, value string) *HashMapValue {
	return &HashMapValue{
		hash:  hash,
		value: value,
	}
}
