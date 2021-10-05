package hashmap

import (
	"crypto/sha1"
	"math"
)

type ChainedHashMap struct {
	buckets [uint32(math.MaxUint16) + 1][]*HashMapValue
}

func NewChainedHashMap() *ChainedHashMap {
	return &ChainedHashMap{}
}

func (h *ChainedHashMap) Add(key, value string) {
	hash, i := hashFunction(key)
	if v := findHashMapValue(h.buckets[i], hash); v == nil {
		h.buckets[i] = append(h.buckets[i], newHashMapValue(hash, value))
	}
}

func (h *ChainedHashMap) Get(key string) (string, bool) {
	hash, i := hashFunction(key)
	v := findHashMapValue(h.buckets[i], hash)
	if v == nil {
		return "", false
	}
	return v.value, true
}

func (h *ChainedHashMap) Remove(key string) (string, bool) {
	hash, i := hashFunction(key)
	var v *HashMapValue
	h.buckets[i], v = removeValue(h.buckets[i], hash)
	if v == nil {
		return "", false
	}
	return v.value, true
}

func removeValue(values []*HashMapValue, hash [sha1.Size]byte) ([]*HashMapValue, *HashMapValue) {
	var index int
	var found *HashMapValue = nil

	for i, v := range values {
		if v.hash == hash {
			index = i
			found = v
			break
		}
	}

	if found == nil {
		return values, nil
	}

	return append(values[:index], values[index+1:]...), found
}

func findHashMapValue(values []*HashMapValue, hash [sha1.Size]byte) *HashMapValue {
	for _, v := range values {
		if v.hash == hash {
			return v
		}
	}
	return nil
}
