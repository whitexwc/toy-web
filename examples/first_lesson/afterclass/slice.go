package main

func main() {
	s := []int{1, 2, 4, 7}
	// 结果应该是 5, 1, 2, 4, 7
	s = Add(s, 0, 5)

	// 结果应该是5, 9, 1, 2, 4, 7
	s = Add(s, 1, 9)

	// 结果应该是5, 9, 1, 2, 4, 7, 13
	s = Add(s, 6, 13)

	// 结果应该是5, 9, 2, 4, 7, 13
	s = Delete(s, 2)

	// 结果应该是9, 2, 4, 7, 13
	s = Delete(s, 0)

	// 结果应该是9, 2, 4, 7
	s = Delete(s, 4)

}

func Add(s []int, index int, value int) []int {
	length := len(s)
	if index >= length {
		s = append(s, value)
	} else if index < length && index >= 0 {
		s = append(s, s[length-1])
		i := length - 1
		for ; i > index; i -= 1 {
			s[i] = s[i-1]
		}
		s[i] = value
	}
	return s
}

func Delete(s []int, index int) []int {
	length := len(s)
	if index < 0 || index >= length {
		return s
	}
	for i := index; i < length-1; i += 1 {
		s[i] = s[i+1]
	}
	s = s[:length-1]
	return s
}
