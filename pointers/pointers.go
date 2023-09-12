package pointers

import "fmt"

//lets declarea an integer variable X
var x int = 200

//now lets declare a pointer to an integer
var ptr *int

func Pointers() {
	//make ptr as pointer to x
	ptr = &x
	//log ptr
	fmt.Println(ptr) //this prints address but not value : 0x2342b8

	//access value from that pointer
	val := *ptr
	// log access value
	fmt.Println(val) //this prints value : 200

	//modify value of that x with pointer
	*ptr = 500
	fmt.Println("value of val : ", val) // still prints 200
	fmt.Println("value of x : ", x)     // 500

}
