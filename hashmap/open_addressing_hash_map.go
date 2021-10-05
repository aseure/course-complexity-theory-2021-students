package hashmap

import (
	"math"
)

type OpenAddressingHashMap struct {
	buckets [uint32(math.MaxUint16) + 1]*HashMapValue
}

func NewOpenAddressingHashMap() *OpenAddressingHashMap {
	return &OpenAddressingHashMap{}
}

func (h *OpenAddressingHashMap) Add(key, value string) {
	hash, i := hashFunction(key)
	for h.buckets[i] != nil {
		i++
	}
	h.buckets[i] = newHashMapValue(hash, value)
}

func (h *OpenAddressingHashMap) Get(key string) (string, bool) {
	hash, i := hashFunction(key)
	for h.buckets[i] != nil && h.buckets[i].hash != hash {
		i++
	}
	if h.buckets[i] == nil {
		return "", false
	}
	return h.buckets[i].value, true
}

func (h *OpenAddressingHashMap) Remove(key string) (string, bool) {
	hash, i := hashFunction(key)
	for h.buckets[i] != nil && h.buckets[i].hash != hash {
		i++
	}
	if h.buckets[i] == nil {
		return "", false
	}
	value := h.buckets[i].value
	h.buckets[i] = nil
	return value, true
}
