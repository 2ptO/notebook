package main

import "fmt"

/*
 * temp file to play with various features of Go lang.
 */

//Student record
type Student struct {
	id   int    //id of the student.
	name string //name of the student.
}

func main() {
	//playing with slices and maps.
	var students = []Student{{1, "alice"}, {2, "bob"}, {3, "cathy"}}
	var sMap = make(map[int]Student)

	//adding new item to a slice. append doubles the capacity during each op
	students = append(students, Student{4, "deesha"})

	//range over slice returns the index and the object.
	for _, s := range students {
		sMap[s.id] = s
	}

	//range over maps returns the key and the value
	for k, v := range sMap {
		fmt.Println(k, v.name)
	}

	fmt.Printf("Total students = %d, capacity = %d\n", len(students), cap(students))
}
