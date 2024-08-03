package comparator

import (
	"bytes"
	"errors"
	"reflect"
	"time"
)

/**
 *
 * @Author AiTao
 * @Date 2023/10/10 14:01
 * @Url
 **/

const (
	invalid = iota
	equal   = 1 << (iota - 1)
	less
	greater
)

var (
	nilValueError      = errors.New("comparator: the parameter has a nil value")
	typeNotMathError   = errors.New("comparator: type mismatch")
	valueNotMatchError = errors.New("comparator: value mismatch")
	invalidError       = errors.New("comparator: unable to establish a comparative relationship")
)

func Compare(a, b interface{}) int {
	r, _ := compareValue(a, b, false)
	return r
}

func compareValue(a, b interface{}, mark bool) (r int, e error) {
	if a == nil || b == nil {
		if a == b {
			return equal, nil
		} else if a == nil {
			return less, nilValueError
		} else {
			return greater, nilValueError
		}
	}
	ta, tb := reflect.TypeOf(a), reflect.TypeOf(b)
	if ta != tb {
		return invalid, typeNotMathError // 类型不一致
	}
	switch ta.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128,
		reflect.Bool,
		reflect.String:
		return comparePrimitiveValue(a, b)
	case reflect.Pointer, reflect.Struct, reflect.Array, reflect.Slice, reflect.Map:
		return reflectCompareValue(a, b, reflect.ValueOf(a), reflect.ValueOf(b), mark)
	default:
		if reflect.DeepEqual(a, b) {
			return equal, nil
		}
	}
	return invalid, invalidError
}

func reflectCompareValue(a, b interface{}, va, vb reflect.Value, rmark bool) (r int, e error) {
	if !va.IsValid() || !vb.IsValid() {
		if o1, o2 := va.IsValid(), vb.IsValid(); o1 == o2 {
			return equal, nil
		} else if o1 {
			return less, nilValueError
		} else {
			return greater, nilValueError
		}
	}
	ta, tb := va.Type(), vb.Type()
	if ta != tb {
		return invalid, typeNotMathError // 类型不一致
	}
	switch ta.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128,
		reflect.Bool,
		reflect.String:
		return reflectComparePrimitiveValue(va, vb)
	case reflect.Pointer:
		return comparePointer(a, b, va, vb)
	case reflect.Struct:
		return compareStruct(a, b, va, vb, rmark)
	case reflect.Array:
		return reflectCompareSliceValue(a, b, reflect.ValueOf(a), reflect.ValueOf(b))
	case reflect.Slice:
		if va.UnsafePointer() == vb.UnsafePointer() {
			return equal, nil
		}
		if elemtyp := ta.Elem(); isPrimitive(elemtyp.Kind()) || elemtyp.String() == "interface {}" {
			return compareSliceValue(a, b, va, vb, rmark)
		}
		return reflectCompareSliceValue(a, b, va, vb)
	case reflect.Map:
		if va.UnsafePointer() == vb.UnsafePointer() {
			return equal, nil
		}
		if keytyp := ta.Key(); isPrimitive(keytyp.Kind()) || keytyp.String() == "interface {}" {
			return compareMapValue(a, b, va, vb, rmark)
		}
		return compareMap(a, b, va, vb)
	default:
		var x, y interface{}
		if !rmark {
			x, y = a, b
		} else {
			x, y = va.Interface(), vb.Interface()
		}
		if reflect.DeepEqual(x, y) {
			return equal, nil
		}
	}
	return invalid, invalidError
}

func comparePrimitiveValue(a, b interface{}) (r int, e error) {
	switch v1 := a.(type) {
	case string:
		if v2, ok := b.(string); !ok {
			return invalid, typeNotMathError
		} else if v1 == v2 {
			return equal, nil
		} else if v1 < v2 {
			return less, nil
		} else {
			return greater, nil
		}
	case bool:
		if v2, ok := b.(bool); !ok {
			return invalid, typeNotMathError
		} else if v1 == v2 {
			return equal, nil
		} else if v2 {
			return less, nil
		} else {
			return greater, nil
		}
	case int:
		if v2, ok := b.(int); !ok {
			return invalid, typeNotMathError
		} else if v1 == v2 {
			return equal, nil
		} else if v1 < v2 {
			return less, nil
		} else {
			return greater, nil
		}
	case int8:
		if v2, ok := b.(int8); !ok {
			return invalid, typeNotMathError
		} else if v1 == v2 {
			return equal, nil
		} else if v1 < v2 {
			return less, nil
		} else {
			return greater, nil
		}
	case int16:
		if v2, ok := b.(int16); !ok {
			return invalid, typeNotMathError
		} else if v1 == v2 {
			return equal, nil
		} else if v1 < v2 {
			return less, nil
		} else {
			return greater, nil
		}
	case int32:
		if v2, ok := b.(int32); !ok {
			return invalid, typeNotMathError
		} else if v1 == v2 {
			return equal, nil
		} else if v1 < v2 {
			return less, nil
		} else {
			return greater, nil
		}
	case int64:
		if v2, ok := b.(int64); !ok {
			return invalid, typeNotMathError
		} else if v1 == v2 {
			return equal, nil
		} else if v1 < v2 {
			return less, nil
		} else {
			return greater, nil
		}
	case uint:
		if v2, ok := b.(uint); !ok {
			return invalid, typeNotMathError
		} else if v1 == v2 {
			return equal, nil
		} else if v1 < v2 {
			return less, nil
		} else {
			return greater, nil
		}
	case uint8:
		if v2, ok := b.(uint8); !ok {
			return invalid, typeNotMathError
		} else if v1 == v2 {
			return equal, nil
		} else if v1 < v2 {
			return less, nil
		} else {
			return greater, nil
		}
	case uint16:
		if v2, ok := b.(uint16); !ok {
			return invalid, typeNotMathError
		} else if v1 == v2 {
			return equal, nil
		} else if v1 < v2 {
			return less, nil
		} else {
			return greater, nil
		}
	case uint32:
		if v2, ok := b.(uint32); !ok {
			return invalid, typeNotMathError
		} else if v1 == v2 {
			return equal, nil
		} else if v1 < v2 {
			return less, nil
		} else {
			return greater, nil
		}
	case uint64:
		if v2, ok := b.(uint64); !ok {
			return invalid, typeNotMathError
		} else if v1 == v2 {
			return equal, nil
		} else if v1 < v2 {
			return less, nil
		} else {
			return greater, nil
		}
	case float32:
		if v2, ok := b.(float32); !ok {
			return invalid, typeNotMathError
		} else if v1 == v2 {
			return equal, nil
		} else if v1 < v2 {
			return less, nil
		} else {
			return greater, nil
		}
	case float64:
		if v2, ok := b.(float64); !ok {
			return invalid, typeNotMathError
		} else if v1 == v2 {
			return equal, nil
		} else if v1 < v2 {
			return less, nil
		} else {
			return greater, nil
		}
	case complex64:
		v2, ok := b.(complex64)
		if !ok {
			return invalid, typeNotMathError
		}
		if v1r, v1i, v2r, v2i := real(v1), imag(v1), real(v2), imag(v2); v1 == v2 {
			return equal, nil
		} else if v1r < v2r || (v1r == v2r && v1i < v2i) {
			return less, nil
		} else {
			return greater, nil
		}
	case complex128:
		v2, ok := b.(complex128)
		if !ok {
			return invalid, typeNotMathError
		}
		if v1r, v1i, v2r, v2i := real(v1), imag(v1), real(v2), imag(v2); v1 == v2 {
			return equal, nil
		} else if v1r < v2r || (v1r == v2r && v1i < v2i) {
			return less, nil
		} else {
			return greater, nil
		}
	default:
		// 处理由 type 关键字创建的底层类型是基础类型的新类型
		if r, e = reflectComparePrimitiveValue(reflect.ValueOf(a), reflect.ValueOf(b)); r != invalid && e != nil {
			return r, e
		}
		return
	}
}

func reflectComparePrimitiveValue(va, vb reflect.Value) (int, error) {
	switch va.Kind() {
	case reflect.Bool:
		if vb.Kind() != reflect.Bool {
			return invalid, typeNotMathError
		} else if x, y := va.Bool(), vb.Bool(); x == y {
			return equal, nil
		} else if !x && y {
			return less, nil
		} else {
			return greater, nil
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if k := vb.Kind(); k != reflect.Int && k != reflect.Int8 && k != reflect.Int16 && k != reflect.Int32 && k != reflect.Int64 {
			return invalid, typeNotMathError
		} else if x, y := va.Int(), vb.Int(); x == y {
			return equal, nil
		} else if x < y {
			return less, nil
		} else {
			return greater, nil
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if k := vb.Kind(); k != reflect.Uint8 && k != reflect.Uint16 && k != reflect.Uint32 && k != reflect.Uint64 && k != reflect.Uintptr {
			return invalid, typeNotMathError
		} else if x, y := va.Uint(), vb.Uint(); x == y {
			return equal, nil
		} else if x < y {
			return less, nil
		} else {
			return greater, nil
		}
	case reflect.Float32, reflect.Float64:
		if k := vb.Kind(); k != reflect.Float32 && k != reflect.Float64 {
			return invalid, typeNotMathError
		} else if x, y := va.Float(), vb.Float(); x == y {
			return equal, nil
		} else if x < y {
			return less, nil
		} else {
			return greater, nil
		}
	case reflect.Complex64, reflect.Complex128:
		if k := vb.Kind(); k != reflect.Complex64 && k != reflect.Complex128 {
			return invalid, typeNotMathError
		} else if x, y := va.Complex(), vb.Complex(); x == y {
			return equal, nil
		} else if real(x) < real(y) || (real(x) == real(y) && imag(x) < imag(y)) { // 先比较实部，再比较虚部
			return less, nil
		} else {
			return greater, nil
		}
	case reflect.String:
		if vb.Kind() != reflect.String {
			return invalid, typeNotMathError
		} else if x, y := va.String(), vb.String(); x == y {
			return equal, nil
		} else if x < y {
			return less, nil
		} else {
			return greater, nil
		}
	}
	return invalid, invalidError
}

func comparePointer(a, b interface{}, va, vb reflect.Value) (int, error) {
	// 解析多级指针
	for x, y := va.Elem(), vb.Elem(); va.Kind() == reflect.Pointer; va, vb = x, y {
	}
	if !va.IsValid() || !vb.IsValid() {
		if o1, o2 := va.IsValid(), vb.IsValid(); o1 == o2 {
			return equal, nil
		} else if o1 {
			return less, nilValueError
		} else {
			return greater, nilValueError
		}
	}
	return reflectCompareValue(a, b, va, vb, true)
}

func compareStruct(a, b interface{}, va, vb reflect.Value, mark bool) (r int, e error) {
	var v1, v2 interface{}
	if !mark {
		v1, v2 = a, b
	} else {
		v1, v2 = va.Interface(), vb.Interface()
	}
	if t1, o1 := v1.(time.Time); o1 {
		if t2, o2 := v2.(time.Time); o2 {
			if x, y := t1.UnixNano(), t2.UnixNano(); x == y {
				return equal, nil
			} else if x < y {
				return less, nil
			} else {
				return greater, nil
			}
		}
		return invalid, typeNotMathError // 类型不一致
	}
	if c1, o1 := v1.(Iface); o1 {
		if c2, o2 := v2.(Iface); o2 {
			if ret := c1.CompareTo(c2); ret == 0 {
				return equal, nil
			} else if ret < 0 {
				return less, nil
			} else {
				return greater, nil
			}
		}
		return invalid, typeNotMathError // 类型不一致
	}
	// 按字段声明的顺序比较字段值的大小
	if x, y := va.NumField(), vb.NumField(); x == y {
		for i := 0; i < x; i++ {
			if r, e = reflectCompareValue(a, b, va.Field(i), vb.Field(i), true); r != equal {
				return r, e
			}
		}
		return equal, nil
	} else if x < y {
		return less, nil
	} else {
		return greater, nil
	}
}

func compareMap(a, b interface{}, va, vb reflect.Value) (r int, e error) {
	if x, y := va.Len(), vb.Len(); x == y {
		for _, k := range va.MapKeys() {
			v1 := va.MapIndex(k)
			v2 := vb.MapIndex(k)
			if !v1.IsValid() || !v2.IsValid() {
				return invalid, valueNotMatchError
			}
			if r, e = reflectCompareValue(a, b, v1, v2, true); r != equal {
				return r, e
			}
		}
		return equal, nil
	} else if x < y {
		return less, nil
	} else {
		return greater, nil
	}
}

func sliceCompareT[T comparable](s1, s2 []T) (r int, e error) {
	if x, y := len(s1), len(s2); x == y {
		for i := 0; i < x; i++ {
			if r, e = comparePrimitiveValue(s1[i], s2[i]); r != equal {
				return r, e
			}
		}
		return equal, e
	} else if x < y {
		return less, nil
	} else {
		return greater, nil
	}
}

func sliceCompareAny(s1, s2 []interface{}) (r int, e error) {
	if x, y := len(s1), len(s2); x == y {
		for i := 0; i < x; i++ {
			if asPrimitive(s1[i]) {
				if r, e = comparePrimitiveValue(s1[i], s2[i]); r != equal {
					return r, e
				}
			} else {
				v1, v2 := reflect.ValueOf(s1[i]), reflect.ValueOf(s2[i])
				if r, e = reflectCompareValue(s1[i], s2[i], v1, v2, false); r != equal {
					return r, e
				}
			}
		}
		return equal, e
	} else if x < y {
		return less, nil
	} else {
		return greater, nil
	}
}

func mapCompareT[K comparable, V interface{}](m1, m2 map[K]V) (r int, e error) {
	if x, y := len(m1), len(m2); x == y {
		for k, v1 := range m1 {
			if v2, exists := m2[k]; exists {
				if asPrimitive(v1) && asPrimitive(v2) {
					if r, e = comparePrimitiveValue(v1, v2); r != equal {
						return r, e
					}
				} else {
					if r, e = compareValue(v1, v2, false); r != equal {
						return r, e
					}
				}
			} else {
				return invalid, valueNotMatchError
			}
		}
		return equal, nil
	} else if x < y {
		return less, nil
	} else {
		return greater, nil
	}
}

func compareSliceValue(a, b interface{}, va, vb reflect.Value, mark bool) (r int, e error) {
	var x, y interface{}
	if !mark {
		x, y = a, b
	} else {
		x, y = va.Interface(), vb.Interface()
	}
	switch v1 := x.(type) {
	case []byte:
		if v2, ok := y.([]byte); !ok {
			return invalid, valueNotMatchError
		} else if r := bytes.Compare(v1, v2); r == 0 {
			return equal, nil
		} else if r < 0 {
			return less, nil
		} else {
			return greater, nil
		}
	case []string:
		v2, ok := y.([]string)
		if !ok {
			return invalid, valueNotMatchError
		}
		return sliceCompareT(v1, v2)
	case []bool:
		v2, ok := y.([]bool)
		if !ok {
			return invalid, valueNotMatchError
		}
		return sliceCompareT(v1, v2)
	case []int:
		v2, ok := y.([]int)
		if !ok {
			return invalid, valueNotMatchError
		}
		return sliceCompareT(v1, v2)
	case []int8:
		v2, ok := y.([]int8)
		if !ok {
			return invalid, valueNotMatchError
		}
		return sliceCompareT(v1, v2)
	case []int16:
		v2, ok := y.([]int16)
		if !ok {
			return invalid, valueNotMatchError
		}
		return sliceCompareT(v1, v2)
	case []int32:
		v2, ok := y.([]int32)
		if !ok {
			return invalid, valueNotMatchError
		}
		return sliceCompareT(v1, v2)
	case []int64:
		v2, ok := y.([]int64)
		if !ok {
			return invalid, valueNotMatchError
		}
		return sliceCompareT(v1, v2)
	case []uint:
		v2, ok := y.([]uint)
		if !ok {
			return invalid, valueNotMatchError
		}
		return sliceCompareT(v1, v2)
	case []uint16:
		v2, ok := y.([]uint16)
		if !ok {
			return invalid, valueNotMatchError
		}
		return sliceCompareT(v1, v2)
	case []uint32:
		v2, ok := y.([]uint32)
		if !ok {
			return invalid, valueNotMatchError
		}
		return sliceCompareT(v1, v2)
	case []uint64:
		v2, ok := y.([]uint64)
		if !ok {
			return invalid, valueNotMatchError
		}
		return sliceCompareT(v1, v2)
	case []float32:
		v2, ok := y.([]float32)
		if !ok {
			return invalid, valueNotMatchError
		}
		return sliceCompareT(v1, v2)
	case []float64:
		v2, ok := y.([]float64)
		if !ok {
			return invalid, valueNotMatchError
		}
		return sliceCompareT(v1, v2)
	case []complex64:
		v2, ok := y.([]complex64)
		if !ok {
			return invalid, valueNotMatchError
		}
		return sliceCompareT(v1, v2)
	case []complex128:
		v2, ok := y.([]complex128)
		if !ok {
			return invalid, valueNotMatchError
		}
		return sliceCompareT(v1, v2)
	case []interface{}:
		v2, ok := y.([]interface{})
		if !ok {
			return invalid, valueNotMatchError
		}
		return sliceCompareAny(v1, v2)
	}
	return invalid, nil
}

func reflectCompareSliceValue(a, b interface{}, va, vb reflect.Value) (r int, e error) {
	if x, y := va.Len(), vb.Len(); x == y {
		for i := 0; i < x; i++ {
			if r, e = reflectCompareValue(a, b, va.Index(i), vb.Index(i), true); r != equal {
				return r, e
			}
		}
		return equal, e
	} else if x < y {
		return less, nil
	} else {
		return greater, nil
	}
}

func isPrimitive(k reflect.Kind) bool {
	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128,
		reflect.Bool,
		reflect.String:
		return true
	default:
		return false
	}
}

func asPrimitive(v interface{}) bool {
	switch v.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64,
		complex64, complex128,
		string, bool:
		return true
	default:
		return false
	}
}

func compareMapValue(a, b interface{}, va, vb reflect.Value, mark bool) (int, error) {
	var x, y interface{}
	if !mark {
		x, y = a, b
	} else {
		x, y = va.Interface(), vb.Interface()
	}
	switch v := x.(type) {
	case map[string]string:
		return mapCompareT(v, y.(map[string]string))
	case map[string]bool:
		return mapCompareT(v, y.(map[string]bool))
	case map[string]int:
		return mapCompareT(v, y.(map[string]int))
	case map[string]int8:
		return mapCompareT(v, y.(map[string]int8))
	case map[string]int16:
		return mapCompareT(v, y.(map[string]int16))
	case map[string]int32:
		return mapCompareT(v, y.(map[string]int32))
	case map[string]int64:
		return mapCompareT(v, y.(map[string]int64))
	case map[string]uint:
		return mapCompareT(v, y.(map[string]uint))
	case map[string]uint8:
		return mapCompareT(v, y.(map[string]uint8))
	case map[string]uint16:
		return mapCompareT(v, y.(map[string]uint16))
	case map[string]uint32:
		return mapCompareT(v, y.(map[string]uint32))
	case map[string]uint64:
		return mapCompareT(v, y.(map[string]uint64))
	case map[string]float32:
		return mapCompareT(v, y.(map[string]float32))
	case map[string]float64:
		return mapCompareT(v, y.(map[string]float64))
	case map[string]complex64:
		return mapCompareT(v, y.(map[string]complex64))
	case map[string]complex128:
		return mapCompareT(v, y.(map[string]complex128))
	case map[string]interface{}:
		return mapCompareT(v, y.(map[string]interface{}))
	case map[bool]string:
		return mapCompareT(v, y.(map[bool]string))
	case map[bool]bool:
		return mapCompareT(v, y.(map[bool]bool))
	case map[bool]int:
		return mapCompareT(v, y.(map[bool]int))
	case map[bool]int8:
		return mapCompareT(v, y.(map[bool]int8))
	case map[bool]int16:
		return mapCompareT(v, y.(map[bool]int16))
	case map[bool]int32:
		return mapCompareT(v, y.(map[bool]int32))
	case map[bool]int64:
		return mapCompareT(v, y.(map[bool]int64))
	case map[bool]uint:
		return mapCompareT(v, y.(map[bool]uint))
	case map[bool]uint8:
		return mapCompareT(v, y.(map[bool]uint8))
	case map[bool]uint16:
		return mapCompareT(v, y.(map[bool]uint16))
	case map[bool]uint32:
		return mapCompareT(v, y.(map[bool]uint32))
	case map[bool]uint64:
		return mapCompareT(v, y.(map[bool]uint64))
	case map[bool]float32:
		return mapCompareT(v, y.(map[bool]float32))
	case map[bool]float64:
		return mapCompareT(v, y.(map[bool]float64))
	case map[bool]complex64:
		return mapCompareT(v, y.(map[bool]complex64))
	case map[bool]complex128:
		return mapCompareT(v, y.(map[bool]complex128))
	case map[bool]interface{}:
		return mapCompareT(v, y.(map[bool]interface{}))
	case map[int]string:
		return mapCompareT(v, y.(map[int]string))
	case map[int]bool:
		return mapCompareT(v, y.(map[int]bool))
	case map[int]int:
		return mapCompareT(v, y.(map[int]int))
	case map[int]int8:
		return mapCompareT(v, y.(map[int]int8))
	case map[int]int16:
		return mapCompareT(v, y.(map[int]int16))
	case map[int]int32:
		return mapCompareT(v, y.(map[int]int32))
	case map[int]int64:
		return mapCompareT(v, y.(map[int]int64))
	case map[int]uint:
		return mapCompareT(v, y.(map[int]uint))
	case map[int]uint8:
		return mapCompareT(v, y.(map[int]uint8))
	case map[int]uint16:
		return mapCompareT(v, y.(map[int]uint16))
	case map[int]uint32:
		return mapCompareT(v, y.(map[int]uint32))
	case map[int]uint64:
		return mapCompareT(v, y.(map[int]uint64))
	case map[int]float32:
		return mapCompareT(v, y.(map[int]float32))
	case map[int]float64:
		return mapCompareT(v, y.(map[int]float64))
	case map[int]complex64:
		return mapCompareT(v, y.(map[int]complex64))
	case map[int]complex128:
		return mapCompareT(v, y.(map[int]complex128))
	case map[int]interface{}:
		return mapCompareT(v, y.(map[int]interface{}))
	case map[int8]string:
		return mapCompareT(v, y.(map[int8]string))
	case map[int8]bool:
		return mapCompareT(v, y.(map[int8]bool))
	case map[int8]int:
		return mapCompareT(v, y.(map[int8]int))
	case map[int8]int8:
		return mapCompareT(v, y.(map[int8]int8))
	case map[int8]int16:
		return mapCompareT(v, y.(map[int8]int16))
	case map[int8]int32:
		return mapCompareT(v, y.(map[int8]int32))
	case map[int8]int64:
		return mapCompareT(v, y.(map[int8]int64))
	case map[int8]uint:
		return mapCompareT(v, y.(map[int8]uint))
	case map[int8]uint8:
		return mapCompareT(v, y.(map[int8]uint8))
	case map[int8]uint16:
		return mapCompareT(v, y.(map[int8]uint16))
	case map[int8]uint32:
		return mapCompareT(v, y.(map[int8]uint32))
	case map[int8]uint64:
		return mapCompareT(v, y.(map[int8]uint64))
	case map[int8]float32:
		return mapCompareT(v, y.(map[int8]float32))
	case map[int8]float64:
		return mapCompareT(v, y.(map[int8]float64))
	case map[int8]complex64:
		return mapCompareT(v, y.(map[int8]complex64))
	case map[int8]complex128:
		return mapCompareT(v, y.(map[int8]complex128))
	case map[int8]interface{}:
		return mapCompareT(v, y.(map[int8]interface{}))
	case map[int16]string:
		return mapCompareT(v, y.(map[int16]string))
	case map[int16]bool:
		return mapCompareT(v, y.(map[int16]bool))
	case map[int16]int:
		return mapCompareT(v, y.(map[int16]int))
	case map[int16]int8:
		return mapCompareT(v, y.(map[int16]int8))
	case map[int16]int16:
		return mapCompareT(v, y.(map[int16]int16))
	case map[int16]int32:
		return mapCompareT(v, y.(map[int16]int32))
	case map[int16]int64:
		return mapCompareT(v, y.(map[int16]int64))
	case map[int16]uint:
		return mapCompareT(v, y.(map[int16]uint))
	case map[int16]uint8:
		return mapCompareT(v, y.(map[int16]uint8))
	case map[int16]uint16:
		return mapCompareT(v, y.(map[int16]uint16))
	case map[int16]uint32:
		return mapCompareT(v, y.(map[int16]uint32))
	case map[int16]uint64:
		return mapCompareT(v, y.(map[int16]uint64))
	case map[int16]float32:
		return mapCompareT(v, y.(map[int16]float32))
	case map[int16]float64:
		return mapCompareT(v, y.(map[int16]float64))
	case map[int16]complex64:
		return mapCompareT(v, y.(map[int16]complex64))
	case map[int16]complex128:
		return mapCompareT(v, y.(map[int16]complex128))
	case map[int16]interface{}:
		return mapCompareT(v, y.(map[int16]interface{}))
	case map[int32]string:
		return mapCompareT(v, y.(map[int32]string))
	case map[int32]bool:
		return mapCompareT(v, y.(map[int32]bool))
	case map[int32]int:
		return mapCompareT(v, y.(map[int32]int))
	case map[int32]int8:
		return mapCompareT(v, y.(map[int32]int8))
	case map[int32]int16:
		return mapCompareT(v, y.(map[int32]int16))
	case map[int32]int32:
		return mapCompareT(v, y.(map[int32]int32))
	case map[int32]int64:
		return mapCompareT(v, y.(map[int32]int64))
	case map[int32]uint:
		return mapCompareT(v, y.(map[int32]uint))
	case map[int32]uint8:
		return mapCompareT(v, y.(map[int32]uint8))
	case map[int32]uint16:
		return mapCompareT(v, y.(map[int32]uint16))
	case map[int32]uint32:
		return mapCompareT(v, y.(map[int32]uint32))
	case map[int32]uint64:
		return mapCompareT(v, y.(map[int32]uint64))
	case map[int32]float32:
		return mapCompareT(v, y.(map[int32]float32))
	case map[int32]float64:
		return mapCompareT(v, y.(map[int32]float64))
	case map[int32]complex64:
		return mapCompareT(v, y.(map[int32]complex64))
	case map[int32]complex128:
		return mapCompareT(v, y.(map[int32]complex128))
	case map[int32]interface{}:
		return mapCompareT(v, y.(map[int32]interface{}))
	case map[int64]string:
		return mapCompareT(v, y.(map[int64]string))
	case map[int64]bool:
		return mapCompareT(v, y.(map[int64]bool))
	case map[int64]int:
		return mapCompareT(v, y.(map[int64]int))
	case map[int64]int8:
		return mapCompareT(v, y.(map[int64]int8))
	case map[int64]int16:
		return mapCompareT(v, y.(map[int64]int16))
	case map[int64]int32:
		return mapCompareT(v, y.(map[int64]int32))
	case map[int64]int64:
		return mapCompareT(v, y.(map[int64]int64))
	case map[int64]uint:
		return mapCompareT(v, y.(map[int64]uint))
	case map[int64]uint8:
		return mapCompareT(v, y.(map[int64]uint8))
	case map[int64]uint16:
		return mapCompareT(v, y.(map[int64]uint16))
	case map[int64]uint32:
		return mapCompareT(v, y.(map[int64]uint32))
	case map[int64]uint64:
		return mapCompareT(v, y.(map[int64]uint64))
	case map[int64]float32:
		return mapCompareT(v, y.(map[int64]float32))
	case map[int64]float64:
		return mapCompareT(v, y.(map[int64]float64))
	case map[int64]complex64:
		return mapCompareT(v, y.(map[int64]complex64))
	case map[int64]complex128:
		return mapCompareT(v, y.(map[int64]complex128))
	case map[int64]interface{}:
		return mapCompareT(v, y.(map[int64]interface{}))
	case map[uint]string:
		return mapCompareT(v, y.(map[uint]string))
	case map[uint]bool:
		return mapCompareT(v, y.(map[uint]bool))
	case map[uint]int:
		return mapCompareT(v, y.(map[uint]int))
	case map[uint]int8:
		return mapCompareT(v, y.(map[uint]int8))
	case map[uint]int16:
		return mapCompareT(v, y.(map[uint]int16))
	case map[uint]int32:
		return mapCompareT(v, y.(map[uint]int32))
	case map[uint]int64:
		return mapCompareT(v, y.(map[uint]int64))
	case map[uint]uint:
		return mapCompareT(v, y.(map[uint]uint))
	case map[uint]uint8:
		return mapCompareT(v, y.(map[uint]uint8))
	case map[uint]uint16:
		return mapCompareT(v, y.(map[uint]uint16))
	case map[uint]uint32:
		return mapCompareT(v, y.(map[uint]uint32))
	case map[uint]uint64:
		return mapCompareT(v, y.(map[uint]uint64))
	case map[uint]float32:
		return mapCompareT(v, y.(map[uint]float32))
	case map[uint]float64:
		return mapCompareT(v, y.(map[uint]float64))
	case map[uint]complex64:
		return mapCompareT(v, y.(map[uint]complex64))
	case map[uint]complex128:
		return mapCompareT(v, y.(map[uint]complex128))
	case map[uint]interface{}:
		return mapCompareT(v, y.(map[uint]interface{}))
	case map[uint8]string:
		return mapCompareT(v, y.(map[uint8]string))
	case map[uint8]bool:
		return mapCompareT(v, y.(map[uint8]bool))
	case map[uint8]int:
		return mapCompareT(v, y.(map[uint8]int))
	case map[uint8]int8:
		return mapCompareT(v, y.(map[uint8]int8))
	case map[uint8]int16:
		return mapCompareT(v, y.(map[uint8]int16))
	case map[uint8]int32:
		return mapCompareT(v, y.(map[uint8]int32))
	case map[uint8]int64:
		return mapCompareT(v, y.(map[uint8]int64))
	case map[uint8]uint:
		return mapCompareT(v, y.(map[uint8]uint))
	case map[uint8]uint8:
		return mapCompareT(v, y.(map[uint8]uint8))
	case map[uint8]uint16:
		return mapCompareT(v, y.(map[uint8]uint16))
	case map[uint8]uint32:
		return mapCompareT(v, y.(map[uint8]uint32))
	case map[uint8]uint64:
		return mapCompareT(v, y.(map[uint8]uint64))
	case map[uint8]float32:
		return mapCompareT(v, y.(map[uint8]float32))
	case map[uint8]float64:
		return mapCompareT(v, y.(map[uint8]float64))
	case map[uint8]complex64:
		return mapCompareT(v, y.(map[uint8]complex64))
	case map[uint8]complex128:
		return mapCompareT(v, y.(map[uint8]complex128))
	case map[uint8]interface{}:
		return mapCompareT(v, y.(map[uint8]interface{}))
	case map[uint16]string:
		return mapCompareT(v, y.(map[uint16]string))
	case map[uint16]bool:
		return mapCompareT(v, y.(map[uint16]bool))
	case map[uint16]int:
		return mapCompareT(v, y.(map[uint16]int))
	case map[uint16]int8:
		return mapCompareT(v, y.(map[uint16]int8))
	case map[uint16]int16:
		return mapCompareT(v, y.(map[uint16]int16))
	case map[uint16]int32:
		return mapCompareT(v, y.(map[uint16]int32))
	case map[uint16]int64:
		return mapCompareT(v, y.(map[uint16]int64))
	case map[uint16]uint:
		return mapCompareT(v, y.(map[uint16]uint))
	case map[uint16]uint8:
		return mapCompareT(v, y.(map[uint16]uint8))
	case map[uint16]uint16:
		return mapCompareT(v, y.(map[uint16]uint16))
	case map[uint16]uint32:
		return mapCompareT(v, y.(map[uint16]uint32))
	case map[uint16]uint64:
		return mapCompareT(v, y.(map[uint16]uint64))
	case map[uint16]float32:
		return mapCompareT(v, y.(map[uint16]float32))
	case map[uint16]float64:
		return mapCompareT(v, y.(map[uint16]float64))
	case map[uint16]complex64:
		return mapCompareT(v, y.(map[uint16]complex64))
	case map[uint16]complex128:
		return mapCompareT(v, y.(map[uint16]complex128))
	case map[uint16]interface{}:
		return mapCompareT(v, y.(map[uint16]interface{}))
	case map[uint32]string:
		return mapCompareT(v, y.(map[uint32]string))
	case map[uint32]bool:
		return mapCompareT(v, y.(map[uint32]bool))
	case map[uint32]int:
		return mapCompareT(v, y.(map[uint32]int))
	case map[uint32]int8:
		return mapCompareT(v, y.(map[uint32]int8))
	case map[uint32]int16:
		return mapCompareT(v, y.(map[uint32]int16))
	case map[uint32]int32:
		return mapCompareT(v, y.(map[uint32]int32))
	case map[uint32]int64:
		return mapCompareT(v, y.(map[uint32]int64))
	case map[uint32]uint:
		return mapCompareT(v, y.(map[uint32]uint))
	case map[uint32]uint8:
		return mapCompareT(v, y.(map[uint32]uint8))
	case map[uint32]uint16:
		return mapCompareT(v, y.(map[uint32]uint16))
	case map[uint32]uint32:
		return mapCompareT(v, y.(map[uint32]uint32))
	case map[uint32]uint64:
		return mapCompareT(v, y.(map[uint32]uint64))
	case map[uint32]float32:
		return mapCompareT(v, y.(map[uint32]float32))
	case map[uint32]float64:
		return mapCompareT(v, y.(map[uint32]float64))
	case map[uint32]complex64:
		return mapCompareT(v, y.(map[uint32]complex64))
	case map[uint32]complex128:
		return mapCompareT(v, y.(map[uint32]complex128))
	case map[uint32]interface{}:
		return mapCompareT(v, y.(map[uint32]interface{}))
	case map[uint64]string:
		return mapCompareT(v, y.(map[uint64]string))
	case map[uint64]bool:
		return mapCompareT(v, y.(map[uint64]bool))
	case map[uint64]int:
		return mapCompareT(v, y.(map[uint64]int))
	case map[uint64]int8:
		return mapCompareT(v, y.(map[uint64]int8))
	case map[uint64]int16:
		return mapCompareT(v, y.(map[uint64]int16))
	case map[uint64]int32:
		return mapCompareT(v, y.(map[uint64]int32))
	case map[uint64]int64:
		return mapCompareT(v, y.(map[uint64]int64))
	case map[uint64]uint:
		return mapCompareT(v, y.(map[uint64]uint))
	case map[uint64]uint8:
		return mapCompareT(v, y.(map[uint64]uint8))
	case map[uint64]uint16:
		return mapCompareT(v, y.(map[uint64]uint16))
	case map[uint64]uint32:
		return mapCompareT(v, y.(map[uint64]uint32))
	case map[uint64]uint64:
		return mapCompareT(v, y.(map[uint64]uint64))
	case map[uint64]float32:
		return mapCompareT(v, y.(map[uint64]float32))
	case map[uint64]float64:
		return mapCompareT(v, y.(map[uint64]float64))
	case map[uint64]complex64:
		return mapCompareT(v, y.(map[uint64]complex64))
	case map[uint64]complex128:
		return mapCompareT(v, y.(map[uint64]complex128))
	case map[uint64]interface{}:
		return mapCompareT(v, y.(map[uint64]interface{}))
	case map[float32]string:
		return mapCompareT(v, y.(map[float32]string))
	case map[float32]bool:
		return mapCompareT(v, y.(map[float32]bool))
	case map[float32]int:
		return mapCompareT(v, y.(map[float32]int))
	case map[float32]int8:
		return mapCompareT(v, y.(map[float32]int8))
	case map[float32]int16:
		return mapCompareT(v, y.(map[float32]int16))
	case map[float32]int32:
		return mapCompareT(v, y.(map[float32]int32))
	case map[float32]int64:
		return mapCompareT(v, y.(map[float32]int64))
	case map[float32]uint:
		return mapCompareT(v, y.(map[float32]uint))
	case map[float32]uint8:
		return mapCompareT(v, y.(map[float32]uint8))
	case map[float32]uint16:
		return mapCompareT(v, y.(map[float32]uint16))
	case map[float32]uint32:
		return mapCompareT(v, y.(map[float32]uint32))
	case map[float32]uint64:
		return mapCompareT(v, y.(map[float32]uint64))
	case map[float32]float32:
		return mapCompareT(v, y.(map[float32]float32))
	case map[float32]float64:
		return mapCompareT(v, y.(map[float32]float64))
	case map[float32]complex64:
		return mapCompareT(v, y.(map[float32]complex64))
	case map[float32]complex128:
		return mapCompareT(v, y.(map[float32]complex128))
	case map[float32]interface{}:
		return mapCompareT(v, y.(map[float32]interface{}))
	case map[float64]string:
		return mapCompareT(v, y.(map[float64]string))
	case map[float64]bool:
		return mapCompareT(v, y.(map[float64]bool))
	case map[float64]int:
		return mapCompareT(v, y.(map[float64]int))
	case map[float64]int8:
		return mapCompareT(v, y.(map[float64]int8))
	case map[float64]int16:
		return mapCompareT(v, y.(map[float64]int16))
	case map[float64]int32:
		return mapCompareT(v, y.(map[float64]int32))
	case map[float64]int64:
		return mapCompareT(v, y.(map[float64]int64))
	case map[float64]uint:
		return mapCompareT(v, y.(map[float64]uint))
	case map[float64]uint8:
		return mapCompareT(v, y.(map[float64]uint8))
	case map[float64]uint16:
		return mapCompareT(v, y.(map[float64]uint16))
	case map[float64]uint32:
		return mapCompareT(v, y.(map[float64]uint32))
	case map[float64]uint64:
		return mapCompareT(v, y.(map[float64]uint64))
	case map[float64]float32:
		return mapCompareT(v, y.(map[float64]float32))
	case map[float64]float64:
		return mapCompareT(v, y.(map[float64]float64))
	case map[float64]complex64:
		return mapCompareT(v, y.(map[float64]complex64))
	case map[float64]complex128:
		return mapCompareT(v, y.(map[float64]complex128))
	case map[float64]interface{}:
		return mapCompareT(v, y.(map[float64]interface{}))
	case map[complex64]string:
		return mapCompareT(v, y.(map[complex64]string))
	case map[complex64]bool:
		return mapCompareT(v, y.(map[complex64]bool))
	case map[complex64]int:
		return mapCompareT(v, y.(map[complex64]int))
	case map[complex64]int8:
		return mapCompareT(v, y.(map[complex64]int8))
	case map[complex64]int16:
		return mapCompareT(v, y.(map[complex64]int16))
	case map[complex64]int32:
		return mapCompareT(v, y.(map[complex64]int32))
	case map[complex64]int64:
		return mapCompareT(v, y.(map[complex64]int64))
	case map[complex64]uint:
		return mapCompareT(v, y.(map[complex64]uint))
	case map[complex64]uint8:
		return mapCompareT(v, y.(map[complex64]uint8))
	case map[complex64]uint16:
		return mapCompareT(v, y.(map[complex64]uint16))
	case map[complex64]uint32:
		return mapCompareT(v, y.(map[complex64]uint32))
	case map[complex64]uint64:
		return mapCompareT(v, y.(map[complex64]uint64))
	case map[complex64]float32:
		return mapCompareT(v, y.(map[complex64]float32))
	case map[complex64]float64:
		return mapCompareT(v, y.(map[complex64]float64))
	case map[complex64]complex64:
		return mapCompareT(v, y.(map[complex64]complex64))
	case map[complex64]complex128:
		return mapCompareT(v, y.(map[complex64]complex128))
	case map[complex64]interface{}:
		return mapCompareT(v, y.(map[complex64]interface{}))
	case map[complex128]string:
		return mapCompareT(v, y.(map[complex128]string))
	case map[complex128]bool:
		return mapCompareT(v, y.(map[complex128]bool))
	case map[complex128]int:
		return mapCompareT(v, y.(map[complex128]int))
	case map[complex128]int8:
		return mapCompareT(v, y.(map[complex128]int8))
	case map[complex128]int16:
		return mapCompareT(v, y.(map[complex128]int16))
	case map[complex128]int32:
		return mapCompareT(v, y.(map[complex128]int32))
	case map[complex128]int64:
		return mapCompareT(v, y.(map[complex128]int64))
	case map[complex128]uint:
		return mapCompareT(v, y.(map[complex128]uint))
	case map[complex128]uint8:
		return mapCompareT(v, y.(map[complex128]uint8))
	case map[complex128]uint16:
		return mapCompareT(v, y.(map[complex128]uint16))
	case map[complex128]uint32:
		return mapCompareT(v, y.(map[complex128]uint32))
	case map[complex128]uint64:
		return mapCompareT(v, y.(map[complex128]uint64))
	case map[complex128]float32:
		return mapCompareT(v, y.(map[complex128]float32))
	case map[complex128]float64:
		return mapCompareT(v, y.(map[complex128]float64))
	case map[complex128]complex64:
		return mapCompareT(v, y.(map[complex128]complex64))
	case map[complex128]complex128:
		return mapCompareT(v, y.(map[complex128]complex128))
	case map[complex128]interface{}:
		return mapCompareT(v, y.(map[complex128]interface{}))
	case map[interface{}]string:
		return mapCompareT(v, y.(map[interface{}]string))
	case map[interface{}]bool:
		return mapCompareT(v, y.(map[interface{}]bool))
	case map[interface{}]int:
		return mapCompareT(v, y.(map[interface{}]int))
	case map[interface{}]int8:
		return mapCompareT(v, y.(map[interface{}]int8))
	case map[interface{}]int16:
		return mapCompareT(v, y.(map[interface{}]int16))
	case map[interface{}]int32:
		return mapCompareT(v, y.(map[interface{}]int32))
	case map[interface{}]int64:
		return mapCompareT(v, y.(map[interface{}]int64))
	case map[interface{}]uint:
		return mapCompareT(v, y.(map[interface{}]uint))
	case map[interface{}]uint8:
		return mapCompareT(v, y.(map[interface{}]uint8))
	case map[interface{}]uint16:
		return mapCompareT(v, y.(map[interface{}]uint16))
	case map[interface{}]uint32:
		return mapCompareT(v, y.(map[interface{}]uint32))
	case map[interface{}]uint64:
		return mapCompareT(v, y.(map[interface{}]uint64))
	case map[interface{}]float32:
		return mapCompareT(v, y.(map[interface{}]float32))
	case map[interface{}]float64:
		return mapCompareT(v, y.(map[interface{}]float64))
	case map[interface{}]complex64:
		return mapCompareT(v, y.(map[interface{}]complex64))
	case map[interface{}]complex128:
		return mapCompareT(v, y.(map[interface{}]complex128))
	case map[interface{}]interface{}:
		return mapCompareT(v, y.(map[interface{}]interface{}))

	}
	return invalid, nil
}
