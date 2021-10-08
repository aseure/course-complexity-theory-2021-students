package hashmap

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestHashMap(t *testing.T) {
	type KeyValuePair struct {
		key   string
		value string
	}

	var pairs []KeyValuePair
	for i := 0; i < 10000; i++ {
		pairs = append(pairs, KeyValuePair{time.Now().String(), strconv.Itoa(i)})
	}

	for _, c := range []struct {
		description string
		h           HashMap
	}{
		{"chained hash map", NewChainedHashMap()},
		{"open addressing hash map", NewOpenAddressingHashMap()},
	} {
		t.Run(c.description, func(t *testing.T) {
			for _, pair := range pairs {
				c.h.Add(pair.key, pair.value)

				v, ok := c.h.Get(pair.key)
				require.True(t, ok, "should insert %s in %s successfully", pair.key, c.h)
				require.Equal(t, pair.value, v, "wrong inserted value for key %s in %s", pair.key, c.h)
			}

			time.Sleep(1 * time.Millisecond)

			for _, pair := range pairs {
				v, ok := c.h.Remove(pair.key)
				require.True(t, ok, "should remove %s from %s successfully", pair.key, c.h)
				require.Equal(t, pair.value, v, "wrong removed value for key %s from %s", pair.key, c.h)

				_, ok = c.h.Get(pair.key)
				require.False(t, ok, "should not find removed key from %s", c.h)
			}
		})

	}
}
