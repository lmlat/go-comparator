package comparator

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

/**
 *
 * @Author AiTao
 * @Date 2023/10/10 14:46
 * @Url
 **/

var (
	arr1 = [5]int{1, 2, 3, 4, 5}
	arr2 = [5]int{1, 2, 3, 4, 5}
	arr3 = [5]int{1, 2, 3, 4, 6}
	arr4 = [3]int{7, 8, 9}

	slice1 = []int{1, 2, 3, 4, 5}
	slice2 = []int{1, 2, 3, 4, 5}
	slice3 = []int{1, 2, 3, 4, 6}
	slice4 = []int{7, 8, 9}

	map1 = map[string]int{"a": 1, "b": 2, "c": 3}
	map2 = map[string]int{"a": 1, "b": 2, "c": 3}
	map3 = map[string]int{"a": 1, "b": 2, "c": 4}
	map4 = map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	map5 = map[string]any{"d": map1, "e": map2}
	map6 = map[string]any{"d": map1, "e": map2}
	map7 = map[string]any{"d": map1, "f": map3}
)

type Str string

func TestEquals(t *testing.T) {
	m1 := map[string]string{"sex": "man", "name": "aitao"}
	m2 := map[string]string{"name": "aitao", "sex": "man"}
	m3 := map[string]string{"name": "aitao", "birthday": "2022-03-04"}

	fmt.Println(Equals(m1, m2), reflect.DeepEqual(m1, m2))
	fmt.Println(Equals(m1, m3), reflect.DeepEqual(m1, m3))
	fmt.Println(Equals(m2, m3), reflect.DeepEqual(m2, m3))

	a := "aitao"
	b := "lml"
	c := "tn"
	s1 := []uintptr{uintptr(unsafe.Pointer(&a)), uintptr(unsafe.Pointer(&b)), uintptr(unsafe.Pointer(&c))}
	s2 := []uintptr{uintptr(unsafe.Pointer(&a)), uintptr(unsafe.Pointer(&b)), uintptr(unsafe.Pointer(&c))}
	s3 := []uintptr{uintptr(unsafe.Pointer(&a)), uintptr(unsafe.Pointer(&c)), uintptr(unsafe.Pointer(&b))}
	fmt.Println(Equals(s1, s2), reflect.DeepEqual(s1, s2))
	fmt.Println(Equals(s1, s3), reflect.DeepEqual(s1, s3))
	fmt.Println(Equals(s2, s3), reflect.DeepEqual(s2, s3))
}

func TestCompare(t *testing.T) {
	fmt.Println("=======================Array=================================")
	show(Compare(arr1, arr2), "arr1", "arr2") // equal
	show(Compare(arr1, arr3), "arr1", "arr3") // less
	show(Compare(arr1, arr4), "arr1", "arr4") // invalid
	show(Compare(arr3, arr2), "arr3", "arr2") // greater

	fmt.Println("=======================Slice=================================")
	show(Compare(slice1, slice2), "slice1", "slice2") // equal
	show(Compare(slice1, slice3), "slice1", "slice3") // less
	show(Compare(slice1, slice4), "slice1", "slice4") // greater
	show(Compare(slice3, slice2), "slice3", "slice2") // greater

	fmt.Println("=======================Map=================================")
	show(Compare(map1, map2), "map1", "map2") // equal
	show(Compare(map1, map3), "map1", "map3") // less
	show(Compare(map1, map4), "map1", "map4") // less
	show(Compare(map3, map2), "map3", "map2") // greater
	show(Compare(map5, map6), "map5", "map6") // equal
	show(Compare(map5, map7), "map5", "map7") // invalid

	fmt.Println("=======================Pointer=================================")
	show(Compare(&slice1, &slice2), "p1", "p2") // equal
	show(Compare(&slice1, &slice3), "p1", "p3") // less
	var a, b ***int
	show(Compare(a, b), "pa", "pb") // equal
	//
	fmt.Println("=======================error=================================")
	show(Compare(errors.New("我爱你"), errors.New("我爱你")), "err1", "err2")             // equal
	show(Compare(fmt.Errorf("我爱%s", "你"), fmt.Errorf("我爱%s", "你")), "err3", "err4") // equal
	//
	fmt.Println("=======================包装基础类型=================================")
	show(Compare(Str("aitao"), Str("aitao")), "aitao", "aitao")
	show(Compare(Str("aitao"), Str("aitap")), "aitao", "aitap")
	show(Compare(Str("aitax"), Str("aitao")), "aitax", "aitao")
}

func show(r int, names ...string) {
	switch r {
	case equal:
		fmt.Printf("%s is【equal】than %s\n", names[0], names[1])
	case less:
		fmt.Printf("%s is【less】than %s\n", names[0], names[1])
	case greater:
		fmt.Printf("%s is【greater】than %s\n", names[0], names[1])
	default:
		fmt.Println("invalid")
	}
}

func TestCompareAnySlice(t *testing.T) {
	s1 := []any{"aitao", []string{"go", "python", "java"}, 100, true, 16.8}
	s2 := []any{"aitao", []string{"go", "python", "java"}, 100, true, 16.8}
	fmt.Println(Equals(s1, s2))
}

func BenchmarkCompare(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// BenchmarkCompare-12      10178936               117.6 ns/op            96 B/op		2 allocs/op
		reflect.DeepEqual(map1, map2)
		// BenchmarkCompare-12      9751435               123.8 ns/op            48 B/op		2 allocs/op
		//Equals(map1, map2)
		// BenchmarkCompare-12       601328              1939 ns/op             848 B/op		21 allocs/op
		//cmp.Equal(arr1, arr2)
		//Equals(arr1, arr2)
	}
	b.ReportAllocs()
}
