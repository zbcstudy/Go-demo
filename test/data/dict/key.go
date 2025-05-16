package dict

import "hash/fnv"

type Key struct {
	ID   uint64
	Name string
}

func IsValidKeyType(t any) bool {
	switch t.(type) {
	case float32, float64:
		return true
	case int, int8, int16, int32, int64:
		return true
	case uint, uint8, uint16, uint32, uint64:
		return true
	case string:
		return true
	case Stringer:
		return true
	}
	return false
}

func MakeKey(value any) *Key {
	var name string

	if !IsValidKeyType(value) {
		return nil
	}
	name = toString(value)

	if name == "" {
		return nil
	}

	h := fnv.New64a()
	_, err := h.Write([]byte(name))
	if err != nil {
		return nil
	}
	return &Key{
		ID:   h.Sum64(),
		Name: name,
	}
}
