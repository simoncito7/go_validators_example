// package main

// import (
// 	"fmt"
// )

// const aConstantValue int = 100

// func main() {
// 	// string declaration
// 	var aString string = "This is a string"
// 	fmt.Println(aString)
// 	fmt.Printf("The data type is %T", aString) //Printf allow us to know the data type

// 	var anInteger int = 12
// 	fmt.Print("\nThe integer value is: ", anInteger)

// 	// we can use := instead of var for defining variables. This way of declare variables only works
// 	// inside of a function. For a global scope, we must use the "var" keyword
// 	anotherString := "This is another string"
// 	fmt.Println("\n" + anotherString)

// 	fmt.Println("\nThis is a constant value defined outside the function: ", aConstantValue)
// }
// ---------------------------------------------------------------------------------------

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	aString := bufio.NewReader(os.Stdin)
// 	fmt.Print("Please, insert a text: ")
// 	//fmt.Print(aString.ReadString)
// 	// if you want to ignore a variable, then you have to name it with an underscore
// // In this case, we are ignoring the error object
// 	input, _ := aString.ReadString('\n') // '\n' represents one byte
// 	fmt.Print("You entered: ", input)
// }

// -------------------------------------------------------------
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter a text: ")

// 	input, _ := reader.ReadString('\n')
// 	fmt.Println("You entered ", input)

// 	fmt.Print("Enter a number: ")
// 	numInput, _ := reader.ReadString('\n')
// 	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
// 	/*
// 		In Go, nil is the zero value for pointers, interfaces, maps, slices, channels and function types,
// 		representing an uninitialized value.

// 		nil doesn't mean some "undefined" state, it's a proper value in itself.
// 		An object in Go is nil simply if and only if it's value is nil, which it
// 		can only be if it's of one of the aforementioned types.

// 		nil is like the NullPointer in Java!!
// 	*/
// 	if err != nil {

// 		// es decir que si err no es igual a nil, no hay un *null pointer en err, por lo tanto, existirá
// 		// un error.
// 		fmt.Println("You entered a invalid value! ")
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("The entered value is: ", aFloat)
// 	}
// }

// -------------------------------------------------------------
// // !USING SOME MATH FUNCTIONS (FROM THE math package)
// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	// if the variables are of the same type, then you can define them in the same line
// 	i1, i2, i3 := 12, 12, 12
// 	intSum := i1 + i2 + i3
// 	fmt.Print("The result of the sum is: ", intSum)

// 	f1, f2, f3 := 21.4, 23.8, 234.1231231231
// 	floatSum := f1 + f2 + float64(f3)
// 	floatSum = math.Round(floatSum*1000) / 1000 // to round in 3 decimals after the point
// 	fmt.Print("\nThe result of the sum is: ", floatSum)

// 	// another way to add digits after the point is using the notation %.2f or %.3f for example, for two
// 	// or three digits, respectively.
// 	circleRadius := 23
// 	circleCircumference := 2 * math.Pi * float64(circleRadius)
// 	fmt.Printf("\nThe circle circumference is: %.3f", circleCircumference)

// }
// -------------------------------------------------------------
// !"Calculator"
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// // 	fmt.Print("Enter a number: ")
// // 	numInput, _ := reader.ReadString('\n')
// // 	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
// func main() {
// 	// band := true
// 	keyboard := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter the first number:  ")
// 	numInput, _ := keyboard.ReadString('\n')
// 	float1, err1 := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
// 	if err1 != nil {
// 		fmt.Print(err1)
// 		// band = false
// 		// return
// 		panic("\nValue must be a number!")
// 	}

// 	fmt.Print("\nInsert the second number:  ")
// 	numInput2, _ := keyboard.ReadString('\n')
// 	float2, err2 := strconv.ParseFloat(strings.TrimSpace(numInput2), 64)
// 	if err2 != nil {
// 		fmt.Print(err2)
// 		// band = false
// 		panic("\nValue must be a number!")
// 	}
// 	fmt.Println("The result of the sum is: ", float1+float2)

// 	m := make(map[string]int)
// 	m["hola"] = 2
// 	fmt.Println(m)
// }

// -------------------------------------------------------------
// //!POINTERS
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	anInt := 23
// 	p := &anInt

// 	fmt.Print("The integer is stored in: ", p, " and it contains the following value: ", *p)

// 	var floatPointer *float64
// 	floatValue := 23.3
// 	floatPointer = &floatValue

// 	fmt.Print("\nMemory location: ", floatPointer, "\nContained value: ", *floatPointer)
// }

// -------------------------------------------------------------
// !ARRAYS AND SLICES
// DIfference between slices and arrays is that for slices you don't have to
// define a specific number of elements
// package main

// import "fmt"

// func main() {
// 	var colors [3]string
// 	colors[0] = "Red"
// 	colors[1] = "Green"
// 	colors[2] = "Blue"
// 	count := len(colors)
// 	for i := 0; i < count; i++ {
// 		fmt.Println(colors[i])
// 	}

// 	//array of numbers
// 	var intArr = [3]int{1, 2, 3}

// 	for i := 0; i < len(intArr); i++ {
// 		fmt.Println(intArr[i])
// 	}

// 	var slice = []string{"red", "blue", "green"}
// 	fmt.Println(slice)

// 	slice = append(slice, "purple")
// 	fmt.Println(slice)

// 	// we can create a slice with //!make(arrayType, arrayLen, arrayCapacity)
// 	// if we define make without capacity, we can add as many elements as we want
// 	numArray := make([]int, 5, 6)

// 	for i := 0; i < 5; i++ {
// 		numArray[i] = i
// 	}
// 	fmt.Println(numArray)

// 	numArray = append(numArray, 6)
// 	for i := 0; i < 6; i++ {
// 		numArray[i] = i
// 	}
// 	fmt.Println(numArray)

// 	numArray = append(numArray, 7)
// 	for i := 0; i < 6; i++ {
// 		numArray[i] = i
// 	}
// 	fmt.Println(numArray)

// 	numArray = append(numArray, 6)
// 	for i := 0; i < 6; i++ {
// 		numArray[i] = i
// 	}
// 	fmt.Println(numArray)
// }

// -------------------------------------------------------------
//!Unordered values in maps
// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	countries := make(map[string]string)
// 	fmt.Println(countries)
// 	countries["AR"] = "Argentina"
// 	countries["BR"] = "Brazil"
// 	countries["UR"] = "Uruguay"
// 	fmt.Println(countries)
// 	fmt.Println(countries["AR"])

// 	// for adding a value
// 	countries["PE"] = "Peru"
// 	fmt.Println(countries)
// 	countries["VE"] = "Venezuela"
// 	fmt.Println(countries)
// 	delete(countries, "VE")
// 	fmt.Println(countries)

// 	// to go through the map, we can do
// 	for k, v := range countries {
// 		fmt.Printf("%v: %v\n", k, v)
// 	}

// 	// if I want to put in order the elements of countries, I can do:
// 	keys := make([]string, len(countries))
// 	i := 0
// 	for k := range countries {
// 		keys[i] = k
// 		i++
// 	}
// 	fmt.Println("KEYS")
// 	fmt.Println(keys)
// 	sort.Strings(keys)
// 	fmt.Println("SORTED KEYS")
// 	fmt.Println(keys)
// 	for i := range keys {
// 		fmt.Println(countries[keys[i]])
// 	}
// }

// -------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	pomponzuelo := Dog{"Caniche", "Pompón", 10}
// 	fmt.Print(pomponzuelo)
// 	fmt.Printf("\n%+v\n", pomponzuelo)
// 	fmt.Printf("Breed: %v\nName: %v\nAge: %v", pomponzuelo.Breed, pomponzuelo.Name, pomponzuelo.Age)
// }

// // This is a Dog structure
// type Dog struct {
// 	Breed string
// 	Name  string
// 	Age   int
// }

// // -------------------------------------------------------------
// // !SWITCH STATEMENTS
// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func main() {
// 	rand.Seed(time.Now().Unix())
// 	dow := rand.Intn(7) + 1
// 	fmt.Println("Day", dow)

// 	var option string
// 	switch dow {
// 	case 1:
// 		option = "opcion 1"
// 	case 2:
// 		option = "option 2"
// 	default:
// 		option = "None"
// 	}
// 	println(option)
// }
// -------------------------------------------------------------
// // !LOOPS
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	colors := []string{"Red", "Blue", "Green"}
// 	fmt.Println(colors)
// 	// * Using the "normal" for loop
// 	for i := 0; i < len(colors); i++ {
// 		fmt.Println(colors[i])
// 	}

// 	// * Using the range
// 	for i := range colors {
// 		fmt.Println(colors[i])
// 	}

// 	// * Using the "key, value" model
// 	for _, color := range colors {
// 		fmt.Println(color)
// 	}
// }
// -------------------------------------------------------------
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Printf("The result of the sum is: %v", sum(2, 2))
// 	fmt.Printf("\nThe result of the sum is: %v", sumAllValues(2, 2, 2, 2, 2))
// 	total, avg := sumAllValuesAndAvg(2, 2, 2, 2, 2)
// 	fmt.Printf("\nThe result of the sum is: %v  %v", total, avg)

// }

// func sum(a int, b int) int {
// 	return a + b
// }

// func sumAllValues(values ...int) int {
// 	total := 0
// 	for _, v := range values {
// 		total += v
// 	}
// 	return total
// }

// func sumAllValuesAndAvg(values ...int) (int, float64) {
// 	total := 0
// 	for _, v := range values {
// 		total += v
// 	}
// 	// var avg float64
// 	avg := float64(total / len(values))
// 	return total, avg
// }
// -------------------------------------------------------------
// // !Functions as methods of custom types
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	pomponzuelo := Dog{"Caniche", "Pompón", 10, "Woof!"}
// 	fmt.Print(pomponzuelo)
// 	fmt.Printf("\n%+v\n", pomponzuelo)
// 	// fmt.Printf("Breed: %v\nName: %v\nAge: %v", pomponzuelo.Breed, pomponzuelo.Name, pomponzuelo.Age)
// 	pomponzuelo.SoundType()
// }

// // This is a Dog structure
// type Dog struct {
// 	Breed string
// 	Name  string
// 	Age   int
// 	Sound string
// }
// // SoundType is the function for know what sound is like.
// func (d Dog) SoundType() {
// 	fmt.Println(d.Sound)
// }
// -------------------------------------------------------------
// !CREATING AND READING FILES
// package main

// import (
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"os"
// )

// func main() {
// 	content := "This is a text!!!"
// 	file, err := os.Create("./newFile.txt")
// 	checkError(err)
// 	length, err := io.WriteString(file, content)
// 	checkError(err)
// 	fmt.Printf("Wrote a file with %v characters! ", length)
// 	defer file.Close()
// 	defer readFile("./newFile.txt")

// }

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func readFile(filename string) {
// 	data, err := ioutil.ReadFile(filename)
// 	checkError(err)
// 	fmt.Println("The entered text is: ", string(data))
// }
// -------------------------------------------------------------
// !READING CONTENT FROM THE WEB
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"strings"
// )

// const url = "http://services.explorecalifornia.org/json/tours.php"

// func main() {
// 	resp, err := http.Get(url)
// 	checkError(err)
// 	fmt.Printf("The code of the response is: %T\n", resp) // I get a pointer to an object name response
// 	defer resp.Body.Close()

// 	bytes, err := ioutil.ReadAll(resp.Body) // we read the content in bytes, therefore, we need
// 	// to transform it in string
// 	checkError(err)

// 	contentResp := string(bytes)
// 	fmt.Print(contentResp, "\n")

// 	tours := toursFromJson(contentResp)
// 	for _, v := range tours {
// 		fmt.Println(v.Name)
// 	}

// }

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func toursFromJson(content string) []Tour {
// 	tours := make([]Tour, 0, 20)

// 	decoder := json.NewDecoder(strings.NewReader(content))
// 	_, err := decoder.Token()
// 	checkError(err)

// 	var tour Tour
// 	for decoder.More() {
// 		err := decoder.Decode(&tour)
// 		checkError(err)
// 		tours = append(tours, tour)
// 	}
// 	return tours
// }

// type Tour struct {
// 	Name  string
// 	Price string
// }

// --------------------------------------
// !HACKER-RANK 2
// package main

// import (
// 	"fmt"
// )

// /*
//  * Complete the 'simpleArraySum' function below.
//  *
//  * The function is expected to return an INTEGER.
//  * The function accepts INTEGER_ARRAY ar as parameter.
//  */

// func simpleArraySum(ar []int32) int32 {
// 	sum := int32(0)
// 	for _, value := range ar {
// 		sum += value
// 	}

// 	return sum
// }

// func main() {

// 	// var ar []int32
// 	ar := []int32{1, 2, 3, 4, 5}
// 	result := simpleArraySum(ar)

// 	fmt.Printf("The result is %d\n", result)
// }

// --------------------------------------
// !HACKER-RANK 3
// package main

// import (
// 	"fmt"
// )

/*
 * Complete the 'compareTriplets' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

// func compareTriplets(a []int32, b []int32) []int32 {
// 	aliceBob := []int32{0, 0}
// 	for i := 0; i < len(a) && i < len(b); i++ {
// 		switch {
// 		case a[i] > b[i]:
// 			aliceBob[0] += 1
// 		case a[i] < b[i]:
// 			aliceBob[1] += 1
// 		}
// 	}

// 	return aliceBob
// }

// func main() {
// 	var a []int32
// 	var b []int32

// 	a = append(a, 22, 89, 49)
// 	b = append(b, 2, 6, 9)

// 	result := compareTriplets(a, b)

// 	fmt.Printf("The result is %d", result)
// }
// --------------------------------------
// !HACKER-RANK 4
// package main

// import "fmt"

// func aVeryBigSum(ar []int64) int64 {
// 	acu := int64(0)
// 	for _, v := range ar {
// 		acu += v
// 	}
// 	return acu
// }

// func main() {

// 	ar := []int64{1111111, 222222, 3333333, 44444444444, 5555555555}

// 	result := aVeryBigSum(ar)
// 	fmt.Printf("The result is %d ", result)
// }

// --------------------------------------
// !HACKER-RANK 5
// package main

// import (
// 	"fmt"
// 	"math"
// )

// func diagonalDifference(ar [][]int64) int64 {
// 	absSum := int64(0)
// 	ltr_acu := int64(0)
// 	rtl_acu := int64(0)
// 	for k1, v1 := range ar {
// 		for k2, v2 := range v1 {
// 			if k1 == k2 {
// 				ltr_acu += int64(v2)
// 				fmt.Print("\nv1: ", v2)
// 			}
// 			if (k1 + k2) == (len(v1) - 1) {
// 				rtl_acu += int64(v2)
// 				fmt.Print("\nv2: ", v2)
// 			}
// 		}
// 	}
// 	absSum = int64(math.Abs(float64(ltr_acu) - float64(rtl_acu)))
// 	return absSum
// }

// func main() {

// 	ar := [][]int64{{66, 55, 44}, {11, 22, 33}, {29, 30, 48}}

// 	result := diagonalDifference(ar)
// 	fmt.Printf("The result is %d ", result)
// }
// -----------------------------------------------------------------
// !HACKER-RANK 6
// package main

// import (
// 	"fmt"
// )

// func plusMinus(arr []int) {
// 	var posNum, negNum, Ze float32
// 	for _, v := range arr {
// 		if v > 0 {
// 			posNum += 1
// 		}
// 		if v < 0 {
// 			negNum += 1
// 		}
// 		if v == 0 {
// 			Ze += 1
// 		}
// 	}
// 	n := float32(len(arr))
// 	ratios := []float32{posNum / n, negNum / n, Ze / n}
// 	for _, v := range ratios {
// 		fmt.Printf("%.6f\n", v)
// 	}
// }

// func main() {
// 	arr := []int{-4, 3, -9, 0, 4, 1}

// 	plusMinus(arr)

// }
// -----------------------------------------------------------------
// !HACKER-RANK 7
// package main

// import (
// 	"fmt"
// 	"time"
// )

// import "fmt"

// func staircase(n int32) {
// 	var cad string

// 	var i, j, k int32
// 	for i = 1; i <= n; i++ {
// 		for j = n - 1; j >= i; j-- {
// 			cad += " "
// 		}

// 		for k = 1; k <= i; k++ {
// 			cad += "#"
// 		}

// 		// new line
// 		cad += "\n"
// 	}

// 	fmt.Print(cad)
// }

// func main() {

// 	n := int32(5)
// 	staircase(n)

// }

// -----------------------------------------------------------------
//  !HACKER-RANK 8 - 1536115
// import "fmt"

// func sum(arr []int64) int64 {
// 	var tot int64 = 0
// 	for _, v := range arr {
// 		tot += v
// 	}
// 	return tot
// }

// func miniMaxSum(arr []int64) {
// 	var temp int64
// 	// order elements first
// 	if len(arr) != 5 {
// 		arr = arr[:5]
// 	}
// 	for k := 0; k < len(arr)-1; k++ {
// 		if arr[k] > arr[k+1] {
// 			temp = arr[k]
// 			arr[k] = arr[k+1]
// 			arr[k+1] = temp
// 			k = -1
// 		}
// 	}
// 	min := sum(arr[:len(arr)-1])
// 	max := sum(arr[1:])
// 	fmt.Print(arr)
// 	fmt.Printf("%v %v", min, max)
// }

// func main() {
// 	var arr = [6]int64{256741038, 623958417, 467905213, 714532089, 1138071625}
// 	miniMaxSum(arr[:])
// }
// ------------------------------------------------------------------------
// !HACKER-RANK 9 - 1434369

// func birthdayCakeCandles(arr []int64) int64 {
// 	var max int64
// 	var count int64 = 0

// 	for k := range arr {
// 		if k == 0 {
// 			max = arr[k]
// 		} else {
// 			if arr[k] > max {
// 				max = arr[k]
// 			}
// 		}
// 	}

// 	for k := range arr {
// 		if max == arr[k] {
// 			count++
// 		}
// 	}

// 	return count
// }

// func main() {
// 	var arr = []int64{1, 2, 4, 4, 4}
// 	fmt.Print(birthdayCakeCandles(arr[:]))
// }

// ------------------------------------------------------------------------
// !HACKER-RANK 10 - 1341715

// func timeConversion(s string) string {
// 	originalFormat := "03:04:05PM"
// 	requiredFormat := "15:04:05"
// 	t, err := time.Parse(originalFormat, s)
// 	if err != nil {
// 		return err.Error()
// 	}

// 	return t.Format(requiredFormat)
// }

// func main() {
// 	originalFormat := "08:04:05PM"

// 	fmt.Println(timeConversion(originalFormat))
// }

//
//
//
//
//
//
//
//
//
//
//
//
//
//

// !github.com/asaskevich/govalidator
// // CamelCaseToUnderscore converts from camel case form to underscore separated form. Ex.: MyFunc => my_func
// func CamelCaseToUnderscore(str string) string

// // IsASCII checks if the string contains ASCII chars only. Empty string is valid.
// func IsASCII(str string) bool

// // This will only fail when Email is an invalid email address but not when it's empty:
// type User struct {
// 	Name  string `valid:"-"`
// 	Email string `valid:"email,optional"`
//   }
// // -------------------------------------------------

// type Person struct {
// Name string      `valid:"type(string)"`  // ------> !govalidator.IsType("Bob", "string")
// Age  int         `valid:"type(int)"`
// }

// ---------------------------------------------------------------------------------------------------------------------------------------------------

// ! github.com/go-ozzo/ozzo-validation/v4

// package main

// import (
// 	"fmt"

// 	"github.com/go-ozzo/ozzo-validation/v4"
// 	"github.com/go-ozzo/ozzo-validation/v4/is"
// )

// // Validating a Simple Value
// func main() {
// 	data := "example"
// 	err := validation.Validate(data,
// 		validation.Required,       // not empty
// 		validation.Length(5, 100), // length between 5 and 100
// 		is.URL,                    // is a valid URL
// 	)
// 	fmt.Println(err)
// 	// Output:
// 	// must be a valid URL
// }

// // Validating a Struct
// type Address struct {
// 	Street string
// 	City   string
// 	State  string
// 	Zip    string
// }

// func (a Address) Validate() error {
// 	return validation.ValidateStruct(&a,
// 		// Street cannot be empty, and the length must between 5 and 50
// 		validation.Field(&a.Street, validation.Required, validation.Length(5, 50)),
// 		// City cannot be empty, and the length must between 5 and 50
// 		validation.Field(&a.City, validation.Required, validation.Length(5, 50)),
// 		// State cannot be empty, and must be a string consisting of two letters in upper case
// 		validation.Field(&a.State, validation.Required, validation.Match(regexp.MustCompile("^[A-Z]{2}$"))),
// 		// State cannot be empty, and must be a string consisting of five digits
// 		validation.Field(&a.Zip, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{5}$"))),
// 	)
// }

// a := Address{
//     Street: "123",
//     City:   "Unknown",
//     State:  "Virginia",
//     Zip:    "12345",
// }

// err := a.Validate()
// fmt.Println(err)
// // Output:
// // Street: the length must be between 5 and 50; State: must be in a valid format.

// ---------------------------------------------------------------------------------------------------------------------------------------------------

// // ! Paquete github.com/go-playground/validator/v10

// package main

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"

// 	"github.com/go-playground/validator/v10"
// )

// type Person struct {
// 	Name     string `validate:"omitempty,min=2,max=30"`
// 	Nickname string `json:"nick_name,omitempty" validate:"min=2,max=50"`
// 	Surname  string `validate:"required_with=Name,min=2,max=30"`
// 	Adress   string `validate:"required,min=2,max=30"`
// 	DNI      string `validate:"allowed_length"`
// 	Age      string `validate:"omitempty,numeric,is_of_legal_age"`
// 	Country  string `validate:"iso3166_1_alpha2"`
// 	IP       string `validate:"ipv6"`
// }

// type Gamer struct {
// 	Name    string   `validate:"required,min=2,max=30"`
// 	Jugador string   `validate:"required,min=2,max=30"`
// 	Bot     string   `validate:"required,min=2,max=30"`
// 	Games   []string `validate:"max=3,dive,min=3,max=20"`
// 	Console string   `validate:"omitempty,oneof=Family Sega PC"`
// 	Level   string   `validate:"max=50"`
// }

// type Fingerprint struct {
// 	BrowserIP           string `validate:"omitempty,max=45,ip"`
// 	BrowserLanguage     string `validate:"required,min=1,max=8,bcp47_language_tag"`
// 	BrowserColorDepth   string `validate:"oneof=1 4 8 15 16 24 32 48"`
// 	BrowserScreenHeight string `validate:"required,min=1,max=8"`
// 	BrowserScreenWidth  string `validate:"required,min=1,max=8"`
// 	BrowserTimeZone     string `validate:"required,min=1,max=5"`
// 	BrowserUserAgent    string `validate:"required,max=2048"`
// 	BrowserAcceptHeader string `validate:"required,max=2048"`
// 	BrowserJavaEnabled  bool
// }

// func main() {
// 	valid := validator.New()
// 	_ = valid.RegisterValidation("is_of_legal_age", isOfLegalAge) // register custom validation func
// 	valid.RegisterAlias("allowed_length", `min=8,max=10,numeric`) // register custom alias

// 	// properties := []string{"Language", "ColorDepth", "ScreenHeight", "ScreenWidth",
// 	// 	"TimeZone", "UserAgent", "AcceptHeader"}

// 	fp := Fingerprint{
// 		BrowserIP:           "ads",
// 		BrowserLanguage:     "es-ES123",
// 		BrowserScreenHeight: "768",
// 		BrowserScreenWidth:  "1024",
// 		BrowserTimeZone:     "180",
// 		BrowserUserAgent:    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.71 Safari/537.36",
// 		BrowserAcceptHeader: "accept,accept-language,accept-encoding",
// 		BrowserJavaEnabled:  false,
// 	}

// 	// var fp Fingerprint

// 	err := valid.Struct(fp)
// 	validationErrors := err.(validator.ValidationErrors)
// 	// fmt.Println(validationErrors)
// 	var failProperties []string

// 	for _, err := range validationErrors {
// 		failProperties = append(failProperties, strings.TrimPrefix(err.Field(), "Browser"))
// 		fmt.Println(strings.TrimPrefix(err.Field(), "Browser"))
// 	}

// 	fmt.Println(failProperties)

// 	fmt.Println(fp.Validate())
// }

// // isOfLegalAge returns true if the age is a numeric value and
// // if it's greater or equal than 18
// func isOfLegalAge(fl validator.FieldLevel) bool {
// 	age, err := strconv.Atoi(fl.Field().String())
// 	if err != nil {
// 		return false
// 	}

// 	if age >= 18 {
// 		return true
// 	}

// 	return false
// }

// func (f *Fingerprint) Validate() (bool, validator.ValidationErrors) {
// 	valid := validator.New()

// 	err := valid.Struct(f)
// 	if err != nil {
// 		validationErrors := err.(validator.ValidationErrors)
// 		return false, validationErrors
// 	}

// 	return true, nil
// }

// ----------------------------------------------------
// package practicefolder

// import "fmt"

// func remove(s []int, i int) []int {
// 	s[i] = s[len(s)-1]
// 	return s[:len(s)-1]
// }

// func main() {
// 	all := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	fmt.Println(all) //[0 1 2 3 4 5 6 7 8 9]
// 	all = remove(all, 5)
// 	fmt.Println(all) //[0 1 2 3 4 6 7 8 9]
// }

package main

type Payment struct {
	SiteID string
}

// type Masker interface {
// 	Name(string) string
// 	Email(string) string
// 	Phone(string) string
// 	Address(string) string
// }

type Person struct {
	Name    string
	Email   string
	Phone   string
	Address string
}

// type Store struct {
// 	masker Masker
// }

func main() {
	// array of strings.
	// str := []string{"Geeks", "For", "Geeks"}

	// var masker *Masker

	var person Person

	person.Name = "Simon"

}
