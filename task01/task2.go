func isPalindrome(x int) bool {
    num := x
	numStr := strconv.Itoa(num)

	charArray := []rune(numStr)
	strArray := make([]string, len(charArray))
	for index, value := range charArray {
		//digit, _ := strconv.Atoi(string(value))
		// if value == '-' {
		// 	fmt.Printf("不是回数: ", num)
		// 	return
		// }
		strArray[index] = string(value)
		//fmt.Println(string(value))
	}
	// 倒序
	for i := 0; i < len(strArray)/2; i++ {

		strArray[i], strArray[len(strArray)-1-i] = strArray[len(strArray)-1-i], strArray[i]
	}
	fmt.Println(strArray)
	result := strings.Join(strArray, "")
	// 如果是负数，自动转成0
	digit, _ := strconv.Atoi(result)

	if digit != num {
		return false
	}
    return true

}
