func twoSum(intArr []int, target int) map[int]int {
	dataMap := make(map[int]int, 1)
	for i := 0; i < len(intArr); i++ {
		for j := i + 1; j < len(intArr); j++ {
			value := intArr[i] + intArr[j]
			if value == target {
				dataMap[i] = intArr[i]
				dataMap[j] = intArr[j]
				fmt.Println(intArr[i])
				fmt.Println(intArr[j])
				break
			}
		}
	}

	return dataMap
}
func main() {
	var target int = 9
	intArr := []int{2, 7, 11, 5}
	dataMap := twoSum(intArr, target)
	fmt.Println(dataMap)

}
