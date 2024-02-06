package comparator

import (
	"time"
)

/**
 *
 * @Author AiTao
 * @Date 2023/6/27 19:53
 * @Url
 **/

// New 函数用于初始化 Comparable 实例, 该函数提供了对 int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,
// float32,float64,string,bool,time.Time 等类型数据的比较功能。 它可以根据入参的类型返回一个指定类型的比较器实例
//
// Example:
// New((int)0) 表示返回一个 int 类型的比较器
// New((uint)0) 表示返回一个 uint 类型的比较器
// New("") 	 表示返回一个 string 类型的比较器
func New(defaultValue any) Type {
	switch defaultValue.(type) {
	case string:
		return String
	case int:
		return Int
	case int8:
		return Int8
	case int16:
		return Int16
	case int32:
		return Int32
	case int64:
		return Int64
	case uint:
		return Uint
	case uint8:
		return Uint8
	case uint16:
		return Uint16
	case uint32:
		return Uint32
	case uint64:
		return Uint64
	case float32:
		return Float32
	case float64:
		return Float64
	case bool:
		return Bool
	case time.Time:
		return Time
	default:
		return Comparable
	}
}

// Reverse 返回一个逆序比较器, 如果 a > b, 返回-1; 如果 a = b, 返回0; 如果 a < b, 返回1.
//
// Example:
// Reverse(Int) 返回一个 int 类型的逆序比较器
// Reverse(Int8) 返回一个 int 类型的逆序比较器
// Reverse(New("")) 返回一个 string 类型的逆序比较器
func Reverse(compare Type) Type {
	return func(a, b any) int { return -compare(a, b) }
}

// Comparable 函数用于对 Comparable 类型数据进行类型断言, 并实现基础比较功能.
func Comparable(x, y any) int {
	a, b := x.(Iface), y.(Iface)
	return a.CompareTo(b)
}

// Int 函数用于对 int 类型数据进行类型断言, 并实现基础比较功能.
func Int(x, y any) int {
	a, b := x.(int), y.(int)
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

// Int8 函数用于对 int8 类型数据进行类型断言, 并实现基础比较功能.
func Int8(x, y any) int {
	a, b := x.(int8), y.(int8)
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

// Int16 函数用于对 int16 类型数据进行类型断言, 并实现基础比较功能.
func Int16(x, y any) int {
	a, b := x.(int16), y.(int16)
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

// Int32 函数用于对 int32 类型数据进行类型断言, 并实现基础比较功能.
func Int32(x, y any) int {
	a, b := x.(int32), y.(int32)
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

// Int64 函数用于对 int64 类型数据进行类型断言, 并实现基础比较功能.
func Int64(x, y any) int {
	a, b := x.(int64), y.(int64)
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

// Uint 函数用于对 uint 类型数据进行类型断言, 并实现基础比较功能.
func Uint(x, y any) int {
	a, b := x.(uint), y.(uint)
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

// Uint8 函数用于对 uint8 类型数据进行类型断言, 并实现基础比较功能.
func Uint8(x, y any) int {
	a, b := x.(uint8), y.(uint8)
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

// Uint16 函数用于对 uint16 类型数据进行类型断言, 并实现基础比较功能.
func Uint16(x, y any) int {
	a, b := x.(uint16), y.(uint16)
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

// Uint32 函数用于对 uint32 类型数据进行类型断言, 并实现基础比较功能.
func Uint32(x, y any) int {
	a, b := x.(uint32), y.(uint32)
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

// Uint64 函数用于对 uint64 类型数据进行类型断言, 并实现基础比较功能.
func Uint64(x, y any) int {
	a, b := x.(uint64), y.(uint64)
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

// Float32 函数用于对 float32 类型数据进行类型断言, 并实现基础比较功能.
func Float32(x, y any) int {
	a, b := x.(float32), y.(float32)
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

// Float64 函数用于对 float64 类型数据进行类型断言, 并实现基础比较功能.
func Float64(x, y any) int {
	a, b := x.(float64), y.(float64)
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

// String 函数用于对 string 类型数据进行类型断言, 并实现基础比较功能.
func String(x, y any) int {
	a, b := x.(string), y.(string)
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

// Time 函数用于对 time.Time 类型数据进行类型断言, 并实现基础比较功能.
func Time(x, y any) int {
	a, b := x.(time.Time), y.(time.Time)
	switch {
	case a.After(b):
		return 1
	case a.Before(b):
		return -1
	default:
		return 0
	}
}

// Bool 函数用于对 bool 类型数据进行类型断言, 并实现基础比较功能.
func Bool(x, y any) int {
	a, b := x.(bool), y.(bool)
	switch {
	case a == b:
		return 0
	case a:
		return 1
	default:
		return -1
	}
}

// Complex64 函数用于对 complex64 类型数据进行类型断言, 并实现基础比较功能.
func Complex64(a, b any) int {
	if x, y := a.(complex64), b.(complex64); x == y {
		return 0
	} else if rx, ry := real(x), real(y); rx < ry {
		return -1
	} else if rx == ry && imag(x) < imag(y) {
		return -1
	} else {
		return 1
	}
}

// Complex128 函数用于对 complex128 类型数据进行类型断言, 并实现基础比较功能.
func Complex128(a, b any) int {
	if x, y := a.(complex128), b.(complex128); x == y {
		return 0
	} else if rx, ry := real(x), real(y); rx < ry {
		return -1
	} else if rx == ry && imag(x) < imag(y) {
		return -1
	} else {
		return 1
	}
}

// Byte 函数用于对 byte 类型数据进行类型断言, 并实现基础比较功能.
func Byte(x, y any) int {
	a, b := x.(byte), y.(byte)
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

// Rune 函数用于对 rune 类型数据进行类型断言, 并实现基础比较功能.
func Rune(x, y any) int {
	a, b := x.(rune), y.(rune)
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

// Error 函数用于对 error 类型数据进行类型断言, 并实现基础比较功能.
func Error(x, y any) int {
	a, b := x.(error).Error(), y.(error).Error()
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}
