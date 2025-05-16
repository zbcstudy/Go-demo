package dict

type Dict struct {
	size, version int
	keys          []*Key
	values        map[uint64]any
}

func (d *Dict) Version() int {
	return d.version
}

func (d *Dict) Len() int {
	return d.size
}

func New(args ...any) *Dict {
	d := &Dict{values: map[uint64]any{}}
	return d
}

// Set inserts a new item into the dict. If a value matching the key already exists,
// its value is replaced, otherwise a new item is added.
func (d *Dict) Set(key, value any) *Dict {
	if d == nil {
		d = New()
	}

	mk := MakeKey(key)
	if mk == nil {
		return d
	}
	if _, ok := d.values[mk.ID]; ok {
		d.values[mk.ID] = value
		return d
	}

	d.keys = append(d.keys, mk)
	d.values[mk.ID] = value
	d.size++
	d.version++
	return d
}

func (d *Dict) Update(args ...any) bool {
	if args == nil {
		return false
	}

	v := d.version

	for i := range args {
		if other, ok := args[i].(*Dict); ok {
			for item := range other.Items() {
				d.Set(item.Key, item.Value)
			}
			continue
		}
		for item := range toIterable(args[i]) {
			if item.Key == nil {
				item.Key = d.size
			}
			d.Set(item.Key, item.Value)
		}
	}
	return v != d.version
}

func (d *Dict) Items() <-chan Item {
	ci := make(chan Item)
	go func() {
		defer close(ci)
		for _, key := range d.keys {
			ci <- Item{Key: key.Name, Value: key.ID}
		}
	}()
	return ci
}
