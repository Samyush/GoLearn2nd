package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var aString = "this is GO!!!"
var anInteger int

const value string = "hello"

func main() {

	// loading data off of web
	loadingWeb()

	// write a txt file
	writeFile()

	// reading of the created file
	defer readFile("./fromString.txt")

	// implementation of struct
	usingStruct()

	// total values
	addAllValues(1, 2, 3, 4, 5)
	// imp of maps
	impOfMap()
	// Implements of slice
	usingSlice()
	// using of array
	usingArray()
	// pointers here
	pointersH()
	// working with time
	workingMath()
	// below is working with date and time
	dateAndTime()
	// the below are for output
	fmt.Println("hello world - " + aString)
	fmt.Printf("The varaible type is %T\n", aString)

	fmt.Println(anInteger)
	fmt.Printf("The varaible type is %T\n", anInteger)

	// the underneath are for inputs of the code
	getInput := bufio.NewReader(os.Stdin) //the Stdin stands for standard input
	fmt.Print("Enter text: ")
	input, _ := getInput.ReadString('\n') //if we are to ignore the variable then we are to use _
	fmt.Println("You entered: ", input)

	// conversion of string input to other types
	fmt.Print("Enter a number: ")
	numInput, _ := getInput.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number", aFloat)
	}

	var anInt int = 5
	var aFlt float64 = 42

	sum := float64(anInt) + aFlt

	fmt.Printf("Sum: %v, Type: %T\n", sum, sum)
}

func workingMath() {
	i1, i2, i3, i4, i5 := 12, 13, 14, 15, 16

	intSum := i1 + i2 + i3 + i4 + i5
	fmt.Println("Math operation complete as sum is ", intSum)

	f1, f2, f3 := 23.5, 23.6, 14.34
	floatSum := f1 + f2 + f3
	fmt.Println("the sum is float", floatSum)

	floatSum = math.Round(floatSum*100) / 100
	fmt.Println("The sum now is ", floatSum)
}

func dateAndTime() {
	n := time.Now()
	fmt.Println("i am execuitng this program at ", n)

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("Go launced at ", t)
	fmt.Println(t.Format(time.ANSIC))
}

// implementation of pointers
func pointersH() {
	anInt := 42
	var p = &anInt
	fmt.Println("Value of p: ", *p)

	value1 := 23.23
	pointer1 := &value1
	fmt.Println("Value 1 is: ", *pointer1)

	*pointer1 = *pointer1 / 31
	fmt.Println("the new value of pointer 1 is: ", pointer1)
}

// implementation of array
func usingArray() {
	var colour [3]string
	colour[0] = "Red"
	colour[1] = "Yellow"
	colour[2] = "Purple"

	fmt.Println(colour)

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
}

func usingSlice() {
	var sliceIs = []string{"one", "two", "three"}
	fmt.Println(sliceIs)

	sliceIs = append(sliceIs, "Purple")
	fmt.Println(sliceIs)

	sliceIs = append(sliceIs[1:])
	fmt.Println(sliceIs)

	sliceIs = append(sliceIs[:len(sliceIs)-1])
	fmt.Println(sliceIs)

	// declare with a type and a initial size with built in make fuction
	// takes in three function 1> type of argument 2>initial length 3>optional capacity

	numbers := make([]int, 5, 5)
	numbers[0] = 123
	numbers[1] = 143
	numbers[2] = 163
	numbers[3] = 128
	numbers[4] = 121

	fmt.Println(numbers)

	numbers = append(numbers, 777)
	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)

}

// implementation of maps
func impOfMap() {
	maps := map[string]float64{
		"mapper":    34,
		"jerry":     34,
		"dapper":    34,
		"scrambler": 34,
	}
	fmt.Println(maps)

	for k, v := range maps {
		fmt.Printf("%v: %v\n", k, v)
	}
}

// the struct not working

// struct in go is a custom type,, struct in go is a data structure
type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

func usingStruct() {
	poodle := Dog{"Poodle", 10, "woof"}
	fmt.Println(poodle)
	poodle.Speak()

	poodle.Sound = "Arf!"
	poodle.Speak()
}

// func (sth *Dog) Here() bool {
// 	myDog := sth.Breed
// 	if sth.Weight {
// 		myDog = "zaira"
// 	}
// 	return true
// }

func addAllValues(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

// using funciton with struct --> Receiver (d Dog)
func (d Dog) Speak() {
	d.Sound = "Meoo"
	fmt.Println(d.Sound)
}

// writing and reading files
func writeFile() {
	content := "Hello form Go, this is my trying to write in txt file using go!"
	file, err := os.Create("./fromString.txt")
	checkError(err)

	// this must be returning length of the conten written
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("Wrote a file with %v characters\n", length)
	defer file.Close()
}

// checks for any type of error
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// reading of file
func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkError(err)
	fmt.Println("Text read from file is as: ", string(data))
}

// working with we and url to read data
const url = "http://services.explorecalifornia.org/json/tour.php"

func loadingWeb() {
	resp, err := http.Get(url)
	checkError(err)

	fmt.Printf("response type: %T\n", resp)

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	content := string(bytes)
	// fmt.Print(content)

	tours := toursFromJSON(content)
	for _, tour := range tours {
		fmt.Println(tour.Name)
	}
}

func toursFromJSON(content string) []Tour {

	// practice.Second()
	tours := make([]Tour, 0, 20)

	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	checkError(err)

	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		checkError(err)
	}
	return tours
}

type Tour struct {
	Name, Price string
}
