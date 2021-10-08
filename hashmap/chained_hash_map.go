package hashmap

import (
	"crypto/sha1"
	"math"
)

type ChainedHashMap struct {
	buckets [uint32(math.MaxUint16) + 1][]*ChainedHashMapValue
}

type ChainedHashMapValue struct {
	hash  [sha1.Size]byte
	key   string
	value string
}

func newChainedHashMapValue(hash [sha1.Size]byte, key, value string) *ChainedHashMapValue {
	return &ChainedHashMapValue{
		hash:  hash,
		key:   key,
		value: value,
	}
}

func NewChainedHashMap() *ChainedHashMap {
	return &ChainedHashMap{}
}

func (h *ChainedHashMap) Add(key, value string) {
	hash, i := hashFunction(key)
	if v := findChainedHashMapValue(h.buckets[i], hash); v == nil {
		h.buckets[i] = append(h.buckets[i], newChainedHashMapValue(hash, key, value))
	}
}

func (h *ChainedHashMap) Get(key string) (string, bool) {
	hash, i := hashFunction(key)
	v := findChainedHashMapValue(h.buckets[i], hash)
	if v == nil {
		return "", false
	}
	return v.value, true
}

func (h *ChainedHashMap) Remove(key string) (string, bool) {
	hash, i := hashFunction(key)
	var v *ChainedHashMapValue
	h.buckets[i], v = removeChainedHashMapValue(h.buckets[i], hash)
	if v == nil {
		return "", false
	}
	return v.value, true
}

func removeChainedHashMapValue(values []*ChainedHashMapValue, hash [sha1.Size]byte) ([]*ChainedHashMapValue, *ChainedHashMapValue) {
	var index int
	var found *ChainedHashMapValue = nil

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

func findChainedHashMapValue(values []*ChainedHashMapValue, hash [sha1.Size]byte) *ChainedHashMapValue {
	for _, v := range values {
		if v.hash == hash {
			return v
		}
	}
	return nil
}
