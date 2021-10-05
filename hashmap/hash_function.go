package hashmap

import (
	"crypto/sha1"
	"encoding/binary"
)

func hashFunction(key string) ([sha1.Size]byte, uint16) {
	hash := sha1.Sum([]byte(key))
	index := binary.BigEndian.Uint16(hash[sha1.Size-2:])
	return hash, index
}
