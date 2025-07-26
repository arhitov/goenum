package goenum

import (
	"database/sql/driver"
	"fmt"
	"reflect"
)

type ValueMapper interface {
	ValueMap() map[string]any
}

type Validatable interface {
	Valid() bool
}

type ValueSortable interface {
	Sorted() []string
}

func Find[T any](value any) T {
	v, _ := scan[T](value)
	return v
}

func List[T any]() []T {
	var (
		eZero    T
		lst      []T = []T{}
		valueMap map[string]any
	)

	if tt, ok := any(eZero).(ValueMapper); ok {
		valueMap = tt.ValueMap()
	} else {
		panic(fmt.Errorf("type %T does not implement ValueMapper", eZero))
	}

	if sorted, ok := any(eZero).(ValueSortable); ok {
		for _, k := range sorted.Sorted() {
			lst = append(lst, valueMap[k].(T))
		}
	} else {
		for _, v := range valueMap {
			lst = append(lst, v.(T))
		}
	}
	return lst
}

type enumSql interface {
	Valid() bool
	Key() string
}

func enumSqlValue(e enumSql) (driver.Value, error) {
	if !e.Valid() {
		return nil, fmt.Errorf("invalid type: %T[%v]", e, e)
	}
	return e.Key(), nil
}

// newValueMap создает экземпляр VMap безопасным способом
func newMap[VMap any]() VMap {
	var v VMap
	// Если V является указателем (например, *UserStateValues),
	// то инициализируем его
	if reflect.TypeOf(v).Kind() == reflect.Ptr {
		return reflect.New(reflect.TypeOf(v).Elem()).Interface().(VMap)
	}
	return v
}

func scan[T any](value any) (T, error) {
	var (
		eZero    T
		valueMap map[string]any
	)
	if tt, ok := any(eZero).(ValueMapper); ok {
		valueMap = tt.ValueMap()
	} else {
		return eZero, fmt.Errorf("type %T does not implement ValueMapper", eZero)
	}

	return scanValueMapper[T](value, valueMap)
}

func scanValueMapper[T any](
	value any,
	valueMap map[string]any,
) (T, error) {
	var eZero T

	// Обрабатываем как строку или []uint8
	var strValue string
	switch val := value.(type) {
	case string:
		strValue = val
	case []byte:
		strValue = string(val)
	default:
		return eZero, fmt.Errorf("unsupported type for %T: %T[%s]", eZero, value, value)
	}

	if t, ok := valueMap[strValue]; ok {
		if tv, ok := t.(Validatable); ok {
			if !tv.Valid() {
				return eZero, fmt.Errorf("invalid value for %T: %v", eZero, strValue)
			}
		}
		return t.(T), nil
	} else {
		return eZero, fmt.Errorf("cannot value for %T: %v", eZero, strValue)
	}
}
