package main

import "fmt"

func test_for(num int) {
	for i := 0; i < num; i++ {
		// go fmt.Print(i)
		fmt.Print(i)
		fmt.Print("\n")
	}
}

func test_array() [2]int {
	array1 := [2]int{0, 0}
	array1[0] = 1
	array1[1] += 2
	return array1
}

func test_slice() {
	s := make([]int, 2)
	s[0] = 1
	s[1] = 2
	s = append(s, 3, 4, 5)
	fmt.Print(s, "\n")
}

func test_map(a ...int) (map[int]int, error) {
	m := make(map[int]int, 0)
	for _, i := range a {
		m[i] = i * 2
	}
	return m, nil
}

func test_struct() {
	type info struct {
		id   int
		name struct {
			firstname string
			lastname  string
		}
	}

	var p *info
	p = new(info)
	p.id = 1
	p.name.firstname = "Eldon"
	p.name.lastname = "Zhao"
	fmt.Print(p, "\n")
	fmt.Print(*p, "\n")
}
