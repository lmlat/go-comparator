package comparator

/**
 *
 * @Author AiTao
 * @Date 2024-03-09 0:23
 * @Refer
 * @GitHub
 **/

func Less(a, b any) bool {
	r, _ := compareValue(a, b, false)
	return r == less
}
