package comparator

/**
 *
 * @Author AiTao
 * @Date 2024-03-09 0:23
 * @Refer
 * @GitHub
 **/

func Equals(a, b interface{}) bool {
	r, _ := compareValue(a, b, false)
	return r == equal
}
