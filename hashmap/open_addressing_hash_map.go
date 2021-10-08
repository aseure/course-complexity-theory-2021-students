package hashmap

import (
	"crypto/sha1"
	"math"
)

type OpenAddressingHashMap struct {
	buckets [uint32(math.MaxUint16) + 1]*OpenAddressingHashMapValue
}

type OpenAddressingHashMapValue struct {
	hash    [sha1.Size]byte
	key     string
	value   string
	deleted bool
}

func newHashMapValue(hash [sha1.Size]byte, key, value string) *OpenAddressingHashMapValue {
	return &OpenAddressingHashMapValue{
		hash:  hash,
		key:   key,
		value: value,
	}
}

func NewOpenAddressingHashMap() *OpenAddressingHashMap {
	return &OpenAddressingHashMap{}
}

func (h *OpenAddressingHashMap) Add(key, value string) {
	hash, i := hashFunction(key)
	for h.buckets[i] != nil && (!h.buckets[i].deleted || h.buckets[i].hash != hash || h.buckets[i].key != key) {
		i++
	}
	if h.buckets[i] == nil || h.buckets[i].deleted {
		h.buckets[i] = newHashMapValue(hash, key, value)
	}
}

func (h *OpenAddressingHashMap) Get(key string) (string, bool) {
	hash, i := hashFunction(key)
	for h.buckets[i] != nil && (h.buckets[i].hash != hash || h.buckets[i].key != key || h.buckets[i].deleted) {
		i++
	}
	if h.buckets[i] == nil {
		return "", false
	}
	return h.buckets[i].value, true
}

func (h *OpenAddressingHashMap) Remove(key string) (string, bool) {
	hash, i := hashFunction(key)
	for h.buckets[i] != nil && (h.buckets[i].hash != hash || h.buckets[i].key != key || h.buckets[i].deleted) {
		i++
	}
	if h.buckets[i] == nil {
		return "", false
	}
	value := h.buckets[i].value
	h.buckets[i].deleted = true
	return value, true
}
