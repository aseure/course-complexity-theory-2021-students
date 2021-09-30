package hashmap

import (
	"crypto/sha1"
	"encoding/binary"
)

type ChainedHashMap struct {
	buckets [65536][]ChainedHashMapValue
}

type ChainedHashMapValue struct {
	hash  [sha1.Size]byte
	value string
}

func NewChainedHashMap() *ChainedHashMap {
	return &ChainedHashMap{}
}

func (h *ChainedHashMap) Add(key, value string) {
	hash, i := hashKey(key)
	if v := searchValue(h.buckets[i], hash); v == nil {
		h.buckets[i] = append(h.buckets[i], ChainedHashMapValue{hash, value})
	}
}

func (h *ChainedHashMap) Get(key string) (string, bool) {
	hash, i := hashKey(key)
	v := searchValue(h.buckets[i], hash)
	if v == nil {
		return "", false
	}
	return v.value, true
}

func (h *ChainedHashMap) Remove(key string) (string, bool) {
	hash, i := hashKey(key)
	var v *ChainedHashMapValue
	h.buckets[i], v = removeValue(h.buckets[i], hash)
	if v == nil {
		return "", false
	}
	return v.value, true
}

func hashKey(key string) ([sha1.Size]byte, uint16) {
	hash := sha1.Sum([]byte(key))
	index := binary.BigEndian.Uint16(hash[sha1.Size-2:])
	return hash, index
}

func removeValue(values []ChainedHashMapValue, hash [sha1.Size]byte) ([]ChainedHashMapValue, *ChainedHashMapValue) {
	var index int
	var found *ChainedHashMapValue = nil

	for i, v := range values {
		if v.hash == hash {
			index = i
			found = &v
			break
		}
	}

	if found == nil {
		return values, nil
	}

	return append(values[:index], values[index+1:]...), found
}

func searchValue(values []ChainedHashMapValue, hash [sha1.Size]byte) *ChainedHashMapValue {
	for _, v := range values {
		if v.hash == hash {
			return &v
		}
	}
	return nil
}
