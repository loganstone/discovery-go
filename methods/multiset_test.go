package methods

import "fmt"

func ExampleMultiSet() {
	val := MultiSet{}
	fmt.Println(val)
	val.Insert("ventus")
	val.Insert("ventus")
	val.Insert("ventus")
	fmt.Println(val)
	val.Insert("maestro")
	fmt.Println(val.Count("maestro"))
	val.Erase("maestro")
	fmt.Println(val)
	fmt.Println(val.Count("ventus"))
	// Output:
	// { }
	// { ventus ventus ventus }
	// 1
	// { ventus ventus ventus }
	// 3
}
