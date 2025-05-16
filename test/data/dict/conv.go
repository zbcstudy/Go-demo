package dict

import (
	"reflect"
	"strconv"
)

type Stringer interface {
	String() string
}

var stringerType = reflect.TypeOf((*Stringer)(nil)).Elem()

func toFloat64(x any) float64 {
	if v, ok := x.(float32); ok {
		return float64(v)
	}
	return x.(float64)
}

func toUint64(x any) uint64 {
	switch v := x.(type) {
	case uint:
		return uint64(v)
	case uint8:
		return uint64(v)
	case uint16:
		return uint64(v)
	case uint32:
		return uint64(v)
	case uint64:
		return v
	}
	return x.(uint64)
}

func toInt64(x any) int64 {
	switch v := x.(type) {
	case int:
		return int64(v)
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	}
	return x.(int64)
}

func toString(x any) string {
	var s string
	switch v := x.(type) {
	case float32, float64:
		s = strconv.FormatFloat(toFloat64(v), 'f', -1, 64)
	case int, int8, int16, int32, int64:
		s = strconv.FormatInt(toInt64(v), 10)
	case uint, uint8, uint16, uint32, uint64:
		s = strconv.FormatUint(toUint64(v), 10)
	case string:
		s = v
	case Stringer:
		s = v.String()
	}
	return s
}

type Item struct {
	Key   any
	Value any
}

func toIterable(i any) <-chan Item {
	ci := make(chan Item)
	go func() {
		defer close(ci)

		if item, ok := i.(Item); ok {
			ci <- item
			return
		}

		t := reflect.TypeOf(i)
		if t == nil {
			return
		}

		switch v := reflect.ValueOf(i); t.Kind() {
		case reflect.Map:
			if !isKeyType(t.Key()) {
				break
			}
			for iter := v.MapRange(); iter.Next(); {
				ci <- Item{iter.Key().Interface(), iter.Value().Interface()}
			}
		case reflect.Chan:
		L:
			for {
				x, ok := v.Recv()
				if !ok {
					break L
				}
				ci <- Item{Value: x.Interface()}
			}
		case reflect.Array, reflect.Slice:
			for i := 0; i < v.Len(); i++ {
				ci <- Item{i, v.Index(i).Interface()}
			}
		default:
			ci <- Item{nil, v.Interface()}
		}
	}()
	return ci
}

func isKeyType(t reflect.Type) bool {
	kind := t.Kind()
	return (kind > reflect.Bool && kind < reflect.Uintptr) ||
		kind == reflect.Float64 || kind == reflect.Float32 ||
		kind == reflect.String ||
		t.Implements(stringerType)
}
