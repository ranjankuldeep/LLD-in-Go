package trie

import (
	"bytes"
	"fmt"
	"testing"
)

func byteSliceEqual(a, b []Bytes) bool {
	if len(a) != len(b) {
		return false
	}

	result := true
	i, n := 0, len(a)

	for i < n {
		j := 0
		equal := false
		for j < n && !equal {
			if bytes.Equal(a[i], b[j]) {
				equal = true
			}
			j++
		}

		if !equal {
			result = false
		}
		i++
	}
	return result
}

func TestNewTrie(t *testing.T) {
	trie := NewTrie()
	if trie.root == nil {
		t.Errorf("trie root is invalid, expected %v (got %v)", &trieNode{}, trie.root)
	}

	if trie.size != 1 {
		t.Errorf("trie size is invalid, expected %v (got %v)", 1, trie.size)
	}
}

func TestTrieSize(t *testing.T) {
	trie := NewTrie()
	if trie.size != 1 {
		t.Errorf("trie size is invalid, expected %v (got %v)", 1, trie.size)
	}
}

func TestTrieInsert(t *testing.T) {
	trie := NewTrie()
	size := 1

	kvPairs := map[string][]byte{
		"baby":  []byte("milk"),
		"sugar": []byte("daddy"),
		"lock":  []byte("locker"),
	}
	for key, value := range kvPairs {
		trie.insert([]byte(key), value)
		size++
	}

	if trie.size != size {
		t.Errorf("trie size is invalid, expected %v (got %v)", size, trie.size)
	}
}

func TestTrieSearch(t *testing.T) {
	trie := NewTrie()

	kvPairs := map[string][]byte{
		"baby":  []byte("milk"),
		"sugar": []byte("daddy"),
		"lock":  []byte("locker"),
	}
	for key, value := range kvPairs {
		trie.insert([]byte(key), value)
	}

	// test for correct key and values pair
	for key, value := range kvPairs {
		trieValue, ok := trie.search([]byte(key))
		if !ok {
			t.Errorf("unable to find key %v, expected %v (got %v)", key, true, ok)
		}

		if !bytes.Equal(trieValue, value) {
			t.Errorf("trie value is invalid, expected %v (got %v)", value, trieValue)
		}
	}

	// test for no value for keys
	invalidKey := "idontexist"
	v, ok := trie.search([]byte(invalidKey))
	if ok {
		t.Errorf("invalid key %v should not exist, expected %v (got %v)", invalidKey, nil, v)
	}

	if len(v) != 0 {
		t.Errorf("invalid value for key %v, expected %v (got %v)", invalidKey, []byte{}, v)
	}
}

func TestGetAllValues(t *testing.T) {
	trie := NewTrie()

	kvPairs := map[string][]byte{
		"baby":  []byte("milk"),
		"sugar": []byte("daddy"),
		"lock":  []byte("locker"),
	}
	for key, value := range kvPairs {
		trie.insert([]byte(key), value)
	}

	testCases := []map[string]interface{}{
		map[string]interface{}{
			"expectedLen": 3,
			"expectedValues": []Bytes{
				[]byte("milk"),
				[]byte("daddy"),
				[]byte("locker"),
			},
		},
	}

	for _, tc := range testCases {
		trieValues := trie.GetAllValues()
		if len(trieValues) != tc["expectedLen"].(int) {
			t.Errorf("invalid length of valies returned, expected %v, got %v", tc["expectedLen"].(int), len(trieValues))
		}

		if !byteSliceEqual(trieValues, tc["expectedValues"].([]Bytes)) {
			t.Errorf("invalid values, expected %v, got %v", tc["expectedValues"].([]Bytes), trieValues)
		}
	}
}

func TestGetAllKeys(t *testing.T) {
	trie := NewTrie()

	kvPairs := map[string][]byte{
		"baby":  []byte("milk"),
		"sugar": []byte("daddy"),
		"lock":  []byte("locker"),
	}
	for key, value := range kvPairs {
		trie.insert([]byte(key), value)
	}

	testCases := []map[string]interface{}{
		map[string]interface{}{
			"expectedLen": 3,
			"expectedKeys": []Bytes{
				[]byte("baby"),
				[]byte("sugar"),
				[]byte("lock"),
			},
		},
	}

	for _, tc := range testCases {
		trieKeys := trie.GetAllKeys()
		fmt.Println(trieKeys)
		if len(trieKeys) != tc["expectedLen"].(int) {
			t.Errorf("invalid length of keys returned, expected %v, got %v", tc["expectedLen"].(int), len(trieKeys))
		}

		if !byteSliceEqual(trieKeys, tc["expectedKeys"].([]Bytes)) {
			t.Errorf("invalid keys, expected %v, got %v", tc["expectedKeys"].([]Bytes), trieKeys)
		}
	}
}

func TestGetPrefixKeys(t *testing.T) {
	trie := NewTrie()
	prefix := make([]byte, 0, trie.size)
	trieKeys := trie.GetPrefixKeys(prefix)

	if len(trieKeys) != 0 {
		t.Errorf("invalid length of keys returned. expected: %v (got %v)", 0, len(trieKeys))
	}

	kvPairs := map[string]Bytes{
		"baby":  Bytes{1, 2, 3, 4},
		"bad":   Bytes{2, 1, 4, 6},
		"badly": Bytes{4, 6, 1, 1},
		"bank":  Bytes{7, 7, 4, 4},
		"box":   Bytes{8, 1, 1, 9},
		"dad":   Bytes{9, 0, 1, 1},
		"dance": Bytes{6, 4, 2, 1},
		"zip":   Bytes{0, 0, 1, 2},
	}

	for k, v := range kvPairs {
		trie.insert(Bytes(k), v)
	}

	testCases := []map[string]interface{}{
		map[string]interface{}{
			"prefix":      "z",
			"expectedLen": 1,
			"expectedKeys": []Bytes{
				Bytes("zip"),
			},
		},
		map[string]interface{}{
			"prefix":      "ba",
			"expectedLen": 4,
			"expectedKeys": []Bytes{
				Bytes("baby"),
				Bytes("bad"),
				Bytes("badly"),
				Bytes("bank"),
			},
		},
		map[string]interface{}{
			"prefix":      "bad",
			"expectedLen": 2,
			"expectedKeys": []Bytes{
				Bytes("bad"),
				Bytes("badly"),
			},
		},
	}

	for _, tc := range testCases {
		prefix = Bytes(tc["prefix"].(string))
		trieKeys = trie.GetPrefixKeys(prefix)

		if len(trieKeys) != tc["expectedLen"].(int) {
			t.Errorf("invalid length of keys returned. expected: %v (got %v)",
				tc["expectedLen"].(int),
				len(trieKeys),
			)
		}

		if !byteSliceEqual(trieKeys, tc["expectedKeys"].([]Bytes)) {
			t.Errorf("missing key from expected list of keys for prefix: %v. expected: %v (got %v)",
				prefix,
				tc["expectedKeys"].([]Bytes),
				trieKeys,
			)
		}
	}
}

func TestGetPrefixValues(t *testing.T) {
	trie := NewTrie()

	prefix := []byte{}
	trieValues := trie.GetPrefixValues(prefix)

	if len(trieValues) != 0 {
		t.Errorf("invalid length of keys returned. expected: %v (got %v)", 0, len(trieValues))
	}

	kvPairs := map[string]Bytes{
		"baby":  Bytes{1, 2, 3, 4},
		"bad":   Bytes{2, 1, 4, 6},
		"badly": Bytes{4, 6, 1, 1},
		"bank":  Bytes{7, 7, 4, 4},
		"box":   Bytes{8, 1, 1, 9},
		"dad":   Bytes{9, 0, 1, 1},
		"dance": Bytes{6, 4, 2, 1},
		"zip":   Bytes{0, 0, 1, 2},
	}

	for k, v := range kvPairs {
		trie.insert(Bytes(k), v)
	}

	testCases := []map[string]interface{}{
		map[string]interface{}{
			"prefix":      "z",
			"expectedLen": 1,
			"expectedValues": []Bytes{
				Bytes{0, 0, 1, 2},
			},
		},
		map[string]interface{}{
			"prefix":      "ba",
			"expectedLen": 4,
			"expectedValues": []Bytes{
				Bytes{1, 2, 3, 4},
				Bytes{2, 1, 4, 6},
				Bytes{4, 6, 1, 1},
				Bytes{7, 7, 4, 4},
			},
		},
		map[string]interface{}{
			"prefix":      "bad",
			"expectedLen": 2,
			"expectedValues": []Bytes{
				Bytes{2, 1, 4, 6},
				Bytes{4, 6, 1, 1},
			},
		},
	}

	for _, tc := range testCases {
		prefix = Bytes(tc["prefix"].(string))
		trieValues = trie.GetPrefixValues(prefix)

		if len(trieValues) != tc["expectedLen"].(int) {
			t.Errorf("invalid length of values returned. expected: %v (got %v)",
				tc["expectedLen"].(int),
				len(trieValues),
			)
		}

		if !byteSliceEqual(trieValues, tc["expectedValues"].([]Bytes)) {
			t.Errorf("missing value from expected list of values for prefix: %v. expected: %v (got %v)",
				string(prefix),
				tc["expectedValues"].([]Bytes),
				trieValues,
			)
		}
	}
}
