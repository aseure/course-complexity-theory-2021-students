package hashmap

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashMap(t *testing.T) {
	for _, c := range []struct {
		description string
		h           HashMap
	}{
		{"chained hash map", NewChainedHashMap()},
	} {
		c.h.Add("rob", "pike")

		v, ok := c.h.Get("rob")
		require.True(t, ok)
		require.Equal(t, "pike", v)

		v, ok = c.h.Remove("rob")
		require.True(t, ok)
		require.Equal(t, "pike", v)

		_, ok = c.h.Get("rob")
		require.False(t, ok)
	}

	sum := sha1.Sum([]byte("2021 S2"))
	fmt.Println(sum)
	fmt.Println(hex.EncodeToString(sum[:]))
}
