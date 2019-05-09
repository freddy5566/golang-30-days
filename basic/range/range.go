package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum: ", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}

	kevs := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range kevs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kevs {
		fmt.Println("key: ", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
