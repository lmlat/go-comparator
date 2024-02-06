package comparator

/**
 *
 * @Author AiTao
 * @Date 2024-02-06 19:43
 * @Url
 **/

// Iface 是一个自定义可比较接口, 实现该接口能够赋予实现类型具有可比较功能, 按照指定字段进行有序操作。
// 举个例子, 假设 Example 类型实现了 Iface 接口, 则可以通过以下方式对 Example 类型中的指定字段进行比较操作。
//
//		func (ex Example) CompareTo(c Comparable) int {
//				a, b := ex.FieldA, c.(Example).FieldA
//				switch {
//				case a > b:
//					return 1
//				case a < b:
//					return -1
//	         	default:
//					return 0
//				}
//			}
type Iface interface {
	CompareTo(Iface) int
}

type Type func(any, any) int
