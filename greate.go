package comparator

/**
 *
 * @Author AiTao
 * @Date 2024-03-09 0:23
 * @Refer
 * @GitHub
 **/

func Greater(a, b any) bool {
	r, _ := compareValue(a, b, false)
	return r == greater
}
