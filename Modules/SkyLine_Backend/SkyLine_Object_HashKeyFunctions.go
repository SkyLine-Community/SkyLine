package SkyLine

import (
	"hash/fnv"
	"strconv"
)

func (i *Integer) HashKey() HashKey {
	return HashKey{
		Type_Object: i.Type_Object(),
		Value:       uint64(i.Value),
	}
}

func (f *Float) HashKey() HashKey {
	s := strconv.FormatFloat(f.Value, 'f', -1, 64)
	h := fnv.New64a()
	h.Write([]byte(s))

	return HashKey{
		Type_Object: f.Type_Object(),
		Value:       h.Sum64(),
	}
}

func (b *Boolean_Object) HashKey() HashKey {
	key := HashKey{Type_Object: b.Type_Object()}
	if b.Value {
		key.Value = 1
	}
	return key
}

func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{
		Type_Object: s.Type_Object(),
		Value:       h.Sum64(),
	}
}

func (ui *Unsigned) HashKey() HashKey {
	return HashKey{
		Type_Object: ui.Type_Object(),
		Value:       ui.Value,
	}
}
