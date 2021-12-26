package godsa

import "fmt"


/**
  number of duplicates in a given list of numbers (int)
 */
func main() {

	x := []int{2, 3, 2, 4, 5, 5, 4, 8, 8, 10, 11, 11, 11}
	fmt.Println(x)
	fmt.Println(numberOfOccurrence(x...))
}

func numberOfOccurrence(numbers ...int) map[int]int {
	xMap := make(map[int]int)

	//store in a map
	for _, item := range numbers {
		if _, ok := xMap[item]; ok {
			xMap[item] = xMap[item] + 1
		} else {
			xMap[item] = 1
		}
	}
	return xMap
}
