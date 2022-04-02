package arrays

func Sum(numbers [5]int) int {
	sum := 0
	//range 会迭代数组，每次迭代都会返回数组元素的索引和值。我们选择使用 _ 空白标志符 来忽略索引
	for _, v := range numbers {
		sum += v
	}
	return sum
}
