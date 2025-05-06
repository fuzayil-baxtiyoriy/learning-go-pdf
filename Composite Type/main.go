package main

import "fmt"

func main() {
	// var myArray = [3]int{1, 2, 3}
	// fmt.Println(myArray)

	// var mySlices = []int{10, 20, 30}
	// fmt.Println(mySlices)

	// var x  = []int{20}
	// fmt.Println(x == nil)
	// x = append(x, 10)
	// fmt.Println(x)

	// myMake := make([]int, 5)
	// myMake = append(myMake, 10, 30)
	// fmt.Println(myMake)
	// clear(myMake)
	// fmt.Println(myMake)

	// x1 := []int{1, 2, 3, 4}
	// y := make([]int, 5)
	// num := copy(y, x1)
	// fmt.Println(y, num)

	// xArray := [4]int{5, 6, 7, 8}
	// xSlice := xArray[:]
	// fmt.Println(xSlice)


	var nilMap map[string]int
	fmt.Println(nilMap)

	totalWins := map[string]int{}
	fmt.Println(totalWins)

	teams := map[string][]string {
		"Orcas": []string{"Fred", "Ralph", "Bijou"},
		"Lions": []string{"Sarah", "Peter", "Billie"},
    	"Kittens": []string{"Waldo", "Raul", "Ze"},
	}
	fmt.Println(teams)

	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	fmt.Println(totalWins)

	m := map[string]int{
		"hello": 5,
		"world": 0,
	}

	v, ok := m["hello"]
	fmt.Println(v, ok)

}
