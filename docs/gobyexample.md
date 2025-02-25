coinGo CLI: 
Go is strictly call by value language. Even when you pass a pinter inside the function a copy of the pointer is created.

```go
package main

import "fmt"

func main() {
	a := 10
	aPtr := &a
	increment(aPtr)
	fmt.Println(aPtr) // becaue go is strictly call by value thats why it stil hold the addreess
}

// a copy of pointer of a will be created in increment function
func increment(x *int) {
	x = nil
}
```


But maps and slices are curious case in GO. map and slice implemented as data stricture that hold memory addresses.(they are reference type )
Go  build: to compile the one or more file in same directory
Go run : to compile and run one or more fle in same directory
Go install: compile and install package
Go get: to download row code of another package
Go fmt: to format the coe
Go test:  to run the test

Import statement: to gain access of another packages inside of one we are creating

Basic Types in GO:

```go
String
bool
Int
float64
``` 

 Typed and Untyped Constants
In go constant are treated in a different way than any other language. GO has a very strong type system that doesn’t allow implicit conversion between any of the types. Even with the same numeric types no operation is allowed without explicit conversion. For eg you cannot add a int32 and int64 value. To add those either int32 has to be explicitly converted to int64 or vice versa. However untyped constant have the flexibility of temporary escape from the GO’s type system as we will see in this article.
Untyped constants:
Const a = 100; // hidden type is uint64
Typed constants:
Const a unit32 = 100; // type is explicitly define to unit32
Exampl: below one gives an error 
cannot use a (constant 30 of type int32) as type int64 in assignment
```go
func main() {
   const a int32 = 30
   var i1 int32
   var i2 int64
   i1 = a
   i2 = a
}
```



Character constant in Go:

NOTE: 
typed named character constant can not be assigned to different type with var key word.
Type myChar int32;
Const aa unit32 = ‘a’;

Var bb myChar = aa; // thow error , cant assign to different type

Untyped named character constant can be assign to different type with var keyword
```go
const aa = ‘a’;
 Var bb myChar = aa;
  Var cc int32 = aa;
```


Untyped unnamed char constant can be assing ed to different type with var keyword

```go
Var bb myChar = ‘a’;
Var cc int32 = ‘a’;
```


package main :  a package is a way to group functions, it is maid of all the files in  the directory. Go   consider directory a package. And a package can contain only one function with the name main,
Package == workspace == project

Two type of packages
Executable: if package main is declare on the top of the file and it must have func main.
Reusable: compiling non executable files return nothing. They are helper or dependency code,












Pointers: pointers are variables that store the memory address of another variable.
Its zero value is nil
```go
package main

import "fmt"

func main() {
	// Declare a variable of type int
	var num int = 10

	// Declare a pointer variable of type *int (pointer to int)
	var ptr *int
	fmt.Println("ptr default value ", ptr) // nil
	// Assign the address of 'num' to the pointer variable 'ptr'
	ptr = &num

	// Print the value of 'num' and the value pointed to by 'ptr'
	fmt.Println("Value of num:", num)             // Output: Value of num: 10
	fmt.Println("Value pointed to by ptr:", *ptr) // Output: Value pointed to by ptr: 10

	// Change the value of 'num' using the pointer
	*ptr = 20

	// Print the updated value of 'num'
	fmt.Println("Updated value of num:", num) // Output: Updated value of num: 20
}
```



Array:

Slice: Slice are like reference to array, slice does not hold any data , it just describe the section of underlying array. Like if changing the element of slice , modified the underlaying array element,
A slice s a data structure that holds length, capacity and pointer to the memory where data is stored.

Slicing is the process of taking the portion of slice and creating a new slice from it.the new slice refer to the same underlaying array, making any chaing to new silence will reflect in the original slice.

```go
package main


import "fmt"


func main() {
   var i = [6]int{1, 2, 3, 4, 5, 6}


   var slice []int = i[1:4] //incldue low indice or bound exclude high bound
   fmt.Println(slice)


   var str = [4]string{
       "John",
       "Paul",
       "George",
       "Ringo",
   }


   slice1 := str[0:2]
   slice2 := str[1:3]
   fmt.Println(slice1, slice2)
   slice1[0] = "XXX" // Changing the elements of a slice modifies the corresponding elements of its underlying array.


   fmt.Println(slice1, slice2)
   fmt.Println(str)


   q := []int{2, 3, 4, 5, 6}
   fmt.Println(q)
   r := []bool{true, false, true}
   fmt.Println(q, r)
   s := []struct {
       i int
       j bool
   }{
       {2, true},
       {3, false},
       {4, true},
   }
   fmt.Println("sssssss ", s)


   t := []int{2, 3, 4, 5, 6, 7}
   t = t[1:4]
   fmt.Println("1....... ", t)
   t = t[:2]
   fmt.Println("2........ ", t)
   t = t[1:]
   fmt.Println("3........ ", t)
   t = t[:]
   fmt.Println("4....... ", t)


   // below slice expression are equilentconst


   // var a[10]int
   // a[0:10]
   // a[:10]
   // a[0:]
   // a[:]


}
```




Slice default

var a [10]int
these slice expressions are equivalent:
a[0:10]
a[:10]
a[0:]
a[:]


Slice length and capacity: capacity is equal to number of element in underlaying array. Length equals to no of elements it contains.      

There are four ways of creating a slice
Using the []<type>{} format   ( s:=[]int{})
Creating a slice from another slice or array (
    numbers := [5]int{1, 2, 3, 4, 5}
    //Both start and end
    num1 := numbers[2:4]
)
Using make ( s:= make([]int, 3,5) )
Using new (  new []int  )


NOTE: zero value of slice is nil
   var s []int
   fmt.Println(s, len(s), cap(s))


   if s == nil {
       fmt.Println("nil")
   }




Maps: a map, maps key to value, default value is nil , nor any key can be added to nil map.

```go
package main


import "fmt"


type Vertex struct {
   lat, long float64
}


var m map[string]Vertex // nil map


func main() {
   m = make(map[string]Vertex)
   m["bela vita"] = Vertex{
       40.0344, -43.0877,
   }
   fmt.Println(m["bela vita"])
}
```




Map is like struct literal but keys are required
package main

```go
import "fmt"


type Vertex struct {
   lat, long float64
}


var m = map[string]Vertex{
   "bell labs": Vertex{    // Vertex can be omitted
       23.4566, -34.000,
   },
   "google": Vertex{
       21.00, -24.000,
   },
}


func main() {
   fmt.Println(m)
}
```




Function Values: function are first class citizens in golang, meaning they can be assigned to a variable , passed as an argument to another function and return from another functions.
```go
package main
import "fmt"
func add(a, b int) int {
   return a + b
}


// functiion that take a function as an argument


func applyOperation(a, b int, operation func(a, b int) int) int {
   return operation(a, b)
}


func main() {
   // Assigning the add function to a variable
   var addFunc func(a, b int) int
   addFunc = add


   // using the addFunc variable to call the add function
   sum := addFunc(2, 3)


   fmt.Println("sum ", sum)


   // passing add function as an argument to applyOperation function
   result := applyOperation(2, 3, add)
   fmt.Println("result: ", result)
}
```






Function Closers:  A powerful concept in Golang that allows a function to carry and capture the surrounding lexical scope. even when they are executed outside of their original scope. Lexical scope includes the variables that were in the scope when the closure was created.

```go
package main
import "fmt"


// outer function retrun innner function that captures the varibale x
func outerFunc() func() int {
   x := 0
   // innerFunction capute the variable x from outer Function
   innerFunc := func() int {
       x++
       return x
   }
   return innerFunc
}


func main() {


   // creating a clouser by callig outerFunction
   closuer := outerFunc()
   // calling the  closure multiple times


   fmt.Println(closuer())
   fmt.Println(closuer())
   fmt.Println(closuer())
}
```



Exp2: 

```go
package main
import "fmt"
// outer function calls the innerFuc that add sum to input x
func outerFun() func(int) int {
   sum := 0
   innerFunc := func(x int) int {
       sum += x
       return sum
   }
   return innerFunc


}


func main() {
   // create the two instance of outerFunc each will contain surrouding lexical scope indepently
   pos, neg := outerFun(), outerFun()


   for i := 0; i < 10; i++ {
       // Call the pos closure with the loop index.
       // This adds the loop index to the positive sum.
       // Also, call the neg closure with a negative multiple of the loop index.
       // This subtracts twice the loop index from the negative sum.
       fmt.Println(
           pos(i),
           neg(-2*i),
       )
   }
}```


Methods: a method is a function with a special receiver argument. Argument appear in between func keyword and method name.
```go
package main
import (
   "fmt"
   "math"
)
type Vertex struct {
   x, y float64
}
func (v Vertex) Abs() float64 {
   return math.Sqrt(v.x*v.x + v.y*v.y)
}
func main() {
   v := Vertex{2, 3}
   fmt.Println(v.Abs())
}
```

Note: method can be defined on non struct type like below

```go
package main

import "fmt"

type MyFloat float64

// method can be declared on non struct type
func (f MyFloat) Abs() float64 {
   if f < 0 {
       return float64(-f)
   }
   return float64(f)
}


func main() {
   f := MyFloat(4)
   fmt.Println(f.Abs())
}
```



NOTE: method can not be declared on receiver whose type is defined in other package included int,
NOTE: method can be declared on receiver that are defined in same package as method.


NOTE: methods can be defined with pointer receivers, pointer receivers are more common than value receiver, A method wita  pointer receiver can modified the value to which the pinter point , in below example pointer receiver in scale modified the Vertex value
```go 
package main
import (
   "fmt"
   "math"
)
type Vertex struct {
   x, y float64
}


func (v Vertex) abs() float64 {
   return math.Sqrt(v.x*v.x + v.y*v.y)
}
func (v *Vertex) Scale(f float64) {
   v.x = v.x * f
   v.y = v.y * f
}
func main() {
   v := Vertex{2, 3}
   v.Scale(10)
   fmt.Println(v.abs())
}
```

Pointer and Method Indirection:
Function with pointer argument must take an pointer
Var v Vertex
Scale(&v, 10)   // OK
Scale(v, 10) // compiler error
But the method with a pointer receiver can take either a value or a pointer


Var v Vertex
v.Scale(10) // ok
P := &v;
p.Scale(10) // ok
The method with a pointer receiver called automatically, go interpret v.Scale(10) to (&v).Scale(10)

```go
package main
import "fmt"
type Vertex struct {
   x, y float64
}
func (v *Vertex) Scalee(f float64) {
   v.x = v.x * f
   v.y = v.y * f
}
func ScaleFunc(v *Vertex, f float64) {
   v.x = v.x * f
   v.x = v.x * f
}
func main() {
   v := Vertex{2, 3}
   v.Scalee(10)
   ScaleFunc(&v, 10)
   p := &Vertex{2, 3}
   p.Scalee(10)
   ScaleFunc(p, 10)


   fmt.Println(v, p)
}
```


Reverse is also true like if method with receiver value can take pointer as well as value 
Var v Vertex
v.Scalee(10)  // OK
(&v).Scalle(10)  // OK
But for function
ScaleFunc(v, 10)
ScaleFunc(&v, 10) // compiler error
Choosing a value or pointer receiver: 
1st Method can modified the value  pointer point to 
2nd avoid coping the value each time method is called , it can be efficient when receiver is larger struct even though method dont modified the receiver value,
```go
package main
import (
   "fmt"
   "math"
)
type Vertex struct {
   X, Y float64
}
func (v *Vertex) Scale(f float64) {
   v.X = v.X * f
   v.Y = v.Y * f
}
func (v *Vertex) Abs() float64 {
   return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func main() {
   v := &Vertex{3, 4}
   fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
   v.Scale(5)
   fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
```

Interfaces: interface defined a set of method signatures. A value of interface type can hold any value that implements those methods.
```go
package main
import (
   "fmt"
   "math"
)
type Absr interface {
   Abs() float64
}
type Vertex struct {
   x, y float64
}


type myFloat float64
func main() {
   var d Absr
   f := myFloat(-2)
   v := Vertex{2, 3}
   d = f  // a myFloat implentes interface Absr
   d = &v // a *Vertex implement interface Absr
   // v is vertex not *vertex and does not implement Absr
   // a = v; // a is vertex and dont implement Absr
   fmt.Println(d.Abs())
}


func (a myFloat) Abs() float64 {
   if a < 0 {
       return float64(-a)
   }
   return float64(a)
}


func (a *Vertex) Abs() float64 {
   return math.Sqrt(a.x*a.x + a.y*a.y)
}
```

Interface value: interface value can be thought of tuple of value and underlying concrete type.
{value, type}
Interface value hold a specific underlying concrete type.
Calling a method on a interface value call the same function with underlayting type value 
```go
package main
import (
   "fmt"
   "math"
)
type I interface {
   M()
}


type T struct {
   S string
}
func (t *T) M() {
   fmt.Println(t.S)
}
type F float64
func (f F) M() {
   fmt.Println(f)
}
func main() {
   var i I
   i = &T{"Hello"}
   describe(i)
   i.M()
   i = F(math.Sqrt2)
   describe(i)
   i.M()
}
func describe(i I) {
   fmt.Printf("(%v, %T) \n", i, i)
}
``` 


Interface values with nil underlaying values:
If concrete underlaying value is nil , method would be called with nill receiver,
In some language it trigger nil pointer exception but in Go it is common to write methods that gradually handle the nil receiver value like in below example method “M” handle the nil receiver.
```go
package main
import "fmt"
type I interface {
   M()
}
type T struct {
   S string
}
func (t *T) M() {
   if t == nil {
       fmt.Println("<nil>")
       return
   }
   fmt.Println(t.S)
}
func main() {
   var i I
   var t *T
   i = t
   describe(i)
   i.M()
   i = &T{"hello"}
   describe(i)
   i.M()
}
func describe(i I) {
   fmt.Printf("(%v, %T)\n", i, i)
}
```

NOTE: a nil interface value hold neither value nor concrete type.  Calling a method on nil interface is a run time error because there is no type inside the interface tuple.
```go
package main
import "fmt"
type I interface {
   M()
}
func M() {
}
func main() {
   var i I
   describee(i)
   i.M()
}
func describee(i I) {
   fmt.Println("(%v, %T)\n ", i, i)
}
```

Empty interface : an interface define zero method known as empty interface. An empty interface may hold values of any type. In below example we are assigning string and integer to empty interface
Empty interface is used by the code that handle value of unknow type. Ex: frm.Print takes any number of arguments.
```go
package main
import "fmt"
func main() {
   / / Using empty interface to hold values of different types
   var i interface{}
   i = "hello"
   describeee(i)
   i = 42
   describeee(i)
}
func describeee(i interface{}) {
   fmt.Printf("%v, %T \n", i, i)
}
```

NOTE: The empty interface allows us to work with values of different types without having to define multiple function signatures or data structures for each type. It is commonly used in Go for situations where you need to handle various types dynamically, such as in reflection or when designing generic algorithms. However, excessive use of empty interfaces can make the code harder to understand and maintain, so it should be used judiciously.
Type Assertions: 
A assertion type provide a interface value’s underlaying concrete value
     t: = i.(T)
The statement asserts that interface value i hold the concrete type T and its concrete value assign to t. If it does not hold the data type T, it will trigger a panic.
```go
package main
import "fmt"
func main() {
   var i interface{} = "hello"
   s := i.(string)
   fmt.Println(s)
   s, ok := i.(string)
   fmt.Println(s, ok)


   f, ok := i.(float64)
   fmt.Println(f, ok)
   j := i.(float64) // panic
   fmt.Println(j)
}
```

Usages:  type assertions in Go offer a way to tell the compiler that you are certain a value has specific underlaying type, even though compiler might not be able to infer statically .
Use case 1: 
While working with assertions you may need to call methods and fields on certain concrete type implementation the interface, type assertion allow to convert a interface value to its underlaying type.
var value interface{}
If user, ok := value.(User); ok{
     // user have type User now you can access its methods like user.getName()
}
Usecaes 2 :  can use type assertions in switch startemnt to do certain task based on the value type
Usecase 3: Error Handling:  some type a function can return a error interface, type assertion can be used to check if error is specific type for more precise handling
   err:= someFunction()
 If networkError, ok:= err.(net.Error()); ok{
           fmt.Println(“network error : ”,networkError )
}
Type Switches:  
A type switches is a construct that permits several type assertions in series.
It is like switch statement but cases in type switch specify types instead of values.
switch s:= i.(type){
Case T:
   //
Case S:
//
Case default:
//
}

```go
package main


import "fmt"


func do(i interface{}) {
   switch v := i.(type) {
   case string:
       fmt.Printf("%T, %v\n", v, v)
   case int:
       fmt.Printf("%T, %v\n", v, v)
   default:
       fmt.Printf("i do not know about type %T\n", v)


   }
}


func main() {
   do("hello")
   do(21)
   do(true)
}
```

Stringers: in golang Stringers provide a way to define a method named String() on a custom type that automatically converts the type’s value into human readable string representation.
Benefits:
Improved readability : Stringers make your code more readable by providing a clean string representation of custom types.
Flexibility: you can customized the string representation of your types based on your needs.

```go
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
   fmt.Println(errCode) // output : Error (without stringersm this would print integer value 500)
   unknowCode := 404 //not a defined constant
   fmt.Println(StatusCode(unknowCode))//// Output: Unknown StatusCode
}
```

NOTE: stringers are offen used by go generate tooll to automatically create String() method for a set of constant, this can be useful for managing large set of constant with their string representations.
```go
package main
import "fmt"
type IPAddr [4]int


// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr)String() string{
   return fmt.Sprintf("%d.%d.%d.%d",ip[0], ip[1], ip[2], ip[3])
}
func main() {
   hosts := map[string]IPAddr{
       "loopback":  {127, 0, 0, 1},
       "googleDNS": {8, 8, 8, 8},
   }
   for name, ip := range hosts {
       fmt.Printf("%v: %v\n", name, ip)
   }
}
```

In above function
We define a method named String on the IPAddr type. This method takes the receiver (ip of type IPAddr) and returns a string.


Errors: the error type is build in interface like fmt.stringers , fmt package looks for error interface while printing value. This require Error() string method implementation thats reuters human readable error message.
```go
package main
import "fmt"
type insuffError struct {
   val int
}
func (e insuffError) Error() string {
   return fmt.Sprintf("Insufficient funds you need at least %d funds\n", e.val)
}
func withdraw(balance, amout int) error {
   if balance < amout {
       return insuffError{val: amout}
   }
   return nil
}
func main() {
   balanace := 100
   withdrawl := 150
   err := withdraw(balanace, withdrawl)
   if err != nil {
       fmt.Println("Error ", err)
   } else {
       fmt.Println("withdrawl successful")
   }
}
```

NOTE: function ofter return err value, calling code must handle it
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)


Readers:  io package provides the Reader interface, which represent stream/sequence  of bytes that can be read from. 
Readers are commonly used to read data from various sources such as files, network connections, or in-memory buffers.
Read(p []byte) (n int, err error) 
This allow to read underlaying data upto len(p) bytes into the byte slice p and returns the number of bytes read(n) and error if any 
```go
package main
import (
   "fmt"
   "io"
   "strings"
)


func main() {
   // crearting a reader from a string
   reader := strings.NewReader("hello world!")
   // create a buffer to hold the read data
   buffer := make([]byte, 5)
   for {
       // read data from reader into the buffer
       n, err := reader.Read(buffer)
       if err != nil {
           fmt.Println("Error reading: ", err)
           return
       }
       fmt.Printf("Read %d bytes: %s\n", n, buffer[:n])
       // when reader stream ends returns io.EOF error      
       if err == io.EOF {
           break
       }
   }
}
```



rot13reader:  a common pattern that raps a io.Reader into another io.Reader


Images: images are represented using the “image” package , which provide interfaces to work with images and implementations for common image format such as JPEG, PNG adn GIF.
type image interface {
      ColourMode() colour.Model
     Bounds()  Rectangle
    At()  colour.Colour
}

```go
package main
import (
   "fmt"
   "image"
)
func main() {
   m := image.NewRGBA(image.Rect(0, 0, 100, 100))
   fmt.Println(m.Bounds())
   fmt.Println(m.At(0, 0).RGBA())
}
```



Type parameters:  go functions can be written to work on multiple type using Type parameters , type parameters appear between the brackets before function arguments.
```go
package main
import "fmt"
func Index[T comparable](s []T, x T) int {
   // v and x are type T which has comparable
   for i, v := range s {
       if v == x {
           return i
       }
   }
   return -1
}


func main() {
   si := []int{1, 2, 3, 4}
   fmt.Println(Index(si, 3))
   ss := []string{"foo", "bar", "buz"}
   fmt.Println(Index(ss, "hello"))
}
```


Generic Types: in addition to generic functions go supports generic types, a type can be parameterized with a type parameter, which could be useful for implementing generic data structure,
```go
package main


// List represents a singly-linked list that holds any type
type List[T any] struct {
   next *List[T]
   val  T
}
func main() {


}
```
GOroutunes: a lightweight thread managed by GO runtime
Go routine allow you to perform the concurrent operations concurrently without the need to explicitly create and manage threads
To create a goroutine just prefix a function call with the go keyword. When a function is called as a goroutine , it is executed concurrently with the rest of the program, allowing multiple functions to run concurrently. Like handing multiple clients coneection, IO operations, parallelizing cpu bound task.

```go
package main


import (
   "fmt"
   "time"
)


func sayHello(s string) {
   for i := 0; i < 5; i++ {
       // sleep for a short duration to allow the goroutie to execute
       time.Sleep(100 * time.Millisecond)
       fmt.Println(s)
   }
}


func main() {
   // call sayHello as go routine
   go sayHello("hello")


   sayHello("world")


}
```

## go mod init

go mod init <module path>

Ex: 
go mod init example/m

its creates go.mod file in current directory  to track all package dependencies


## go mod tidy

find module(imported packages in a go file) for packages(go file)

```go
package main

import "fmt"

import "rsc.io/quote"

func main() {
    fmt.Println(quote.Go())
}
```

## Start a module that others can use

mkdir greetings
cd greetings

go mod init create-module/greetings  // it will create go.mod file under greetings dir

## Create the workspace

creating go.work file to specifiy a workspace(multi-module) with the module

initialize the work space

multi-module> go work init ./hello

Run the program in the workspace directory(multi-module)
mutli-module> go run ./hello

NOTE: go commands include all the modules in workspace as main modules. THis allow 
us to refer to a package in the module, even outside the module.


## go mod vendor

go mod tidy
go mod vendor

this places the external dependencies for your chaincode into a local vendor directory




