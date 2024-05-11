package main

import "fmt"

// defined custom type
type StatusCode int

const (
	// value of custom type as integer
	Success StatusCode = 200
	Error   StatusCode = 500
)

func (s StatusCode) String() string {
	switch s {
	case Success:
		return "Success"
	case Error:
		return "Error"
	default:
		return "Unknow status code"
	}
}

func main() {
	code := Success
	fmt.Println(code) // output : success (without stringers, this would print integer value 200)
	errCode := Error
	fmt.Println(errCode)                // output : Error (without stringersm this would print integer value 500)
	unknowCode := 404                   //not a defined constant
	fmt.Println(StatusCode(unknowCode)) //// Output: Unknown StatusCode
}

// package main

// import "fmt"

// type IPAddr [4]int

// // TODO: Add a "String() string" method to IPAddr.
// func (ip IPAddr)String() string{
//     return fmt.Sprintf("%d.%d.%d.%d",ip[0], ip[1], ip[2], ip[3])
// }

// func main() {
// 	hosts := map[string]IPAddr{
// 		"loopback":  {127, 0, 0, 1},
// 		"googleDNS": {8, 8, 8, 8},
// 	}
// 	for name, ip := range hosts {
// 		fmt.Printf("%v: %v\n", name, ip)
// 	}
// }
