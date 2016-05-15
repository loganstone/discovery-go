package methods

import "fmt"

func ExampleMutilSet() {
	val := MultiSet{}
	fmt.Println(val)
	val.Insert("ventus")
	val.Insert("ventus")
	val.Insert("ventus")
	fmt.Println(val)
	val.Insert("maestro")
	fmt.Println(val.Count("maestro"))
	val.Erase("maestro")
	fmt.Println(val.Count("maestro"))
	fmt.Println(val)
	fmt.Println(val.Count("ventus"))
	// Output:
	// { }
	// { ventus ventus ventus }
	// 1
	// 0
	// { ventus ventus ventus }
	// 3
}
