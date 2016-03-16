package main

import (
	/**
	*Period represents
	*dont have to call
	*The libraary it self..
	 */
	. "fmt"
	"math"
)

func init() {
	Println("INITALIZED!")
}

type animal struct {
	numOfLegs int
	name      string
	id        float64
}

type Box struct {
	width  float64
	height float64
	color  string
}

type Circle struct {
	radius float64
}
type Rectangle struct {
	width  float64
	height float64
}

type person struct {
	name string
	age  int
}

type Donkey struct {
	animal
	ratedByCompany float64
}
type student struct {
	person
	specality string
}

type kindaHuman struct {
	legs  bool
	hands bool
}
type alien struct {
	kindaHuman
	eyes  bool
	hands bool
}

func main() {
	/**
	 *Defing Constants..
	 */
	const PI float64 = 3.141595654353453524231543523454235423
	var i int = 3
	/**
	 *Defer gets executed last..
	 */
	defer printLine()
	Printf("Hello World \n")
	Println(i)

	var something bool
	something = false
	Println(something)
	/**
	 *If statement Syntax..
	 *Else if and else directly after..
	 */
	if x := 9; x > 10 {
		Println("x is greater than 10")
	} else if x == 9 {
		Println("x is equal to 9...")
	} else {
		Println("x is less than 10")
	}

	/**
	 * For loop syntax
	 */
	for i := 3; i < 12; i++ {
		Println("Hello")
	}

	sum := 1
	for sum < 1000 {
		sum += sum
		if sum == 2 {
			continue
		}
		Println(sum)

	}

	/**
	 * Group Defintion
	 */
	const (
		a       = iota
		pi      = 3.1415
		prefixx = "Go_"
	)

	Println("A in Const is ", a)

	/**
	 * Array's
	 */
	var arrayExample [5]int
	arrayExample[0] = 5

	Println("First example is ", arrayExample[0])

	//Declaring with previously set values...
	//Go will calculate number when you do ...
	array := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	Println(array)

	doubleArray := [4][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}, [4]int{5, 4, 3, 2}, [4]int{0, 1, 7, 9}}

	Println(doubleArray)

	/**
	 * Slice(Dynamic Array More Widely Used)
	 */

	helloworld := []byte{'H', 'e', 'l', 'l', 'o'}
	Println(helloworld)

	//Put array to Slice..

	var someSLice []int
	someSLice = array[0:2]
	Println(someSLice)

	Println("Length of slice ", len(someSLice))
	Println("Maximum of slice ", cap(someSLice))
	var aSlice []int
	aSlice = array[:3] // equals to aSlice = array[0:3] aSlice has elements 1,2,3
	aSlice = array[5:] // equals to aSlice = array[5:10] aSlice has elements 5,6,7,8,9
	aSlice = array[:]  // equals to aSlice = array[0:10] aSlice has all elements

	copy(someSLice, aSlice)
	someSLice = append(someSLice, 3, 4, 5)
	Println(someSLice)

	/**
	 *Maps! Like Dictonaries
	 */

	myFirstMap := make(map[string]int)
	myFirstMap["one"] = 1
	Println(myFirstMap["one"])

	mySecondMap := make(map[string]int)
	mySecondMap["Fift"] = 50
	Println(mySecondMap["Fift"])

	myLastMap := make(map[bool]int)
	myLastMap[false] = 0
	myLastMap[true] = 1
	Println(myLastMap[false])

	//Map Rating Example....

	ratings := make(map[string]float64)
	ratings["C++"] = 3.1
	ratings["C"] = 4.5
	ratings["Go"] = 3.4

	Println(ratings)

	delete(ratings, "C")

	Println(ratings)

	//**MAP TWO WAY BINDING**//
	myFirstBinding := make(map[string]bool)
	myFirstBinding["false"] = false
	Println(myFirstBinding)
	mySecondBinding := myFirstBinding
	mySecondBinding["false"] = true
	//My First binding changes to true
	//Since it is attached to mySecondBinding
	Println(myFirstBinding)

	/**
	 *Slice Area and itterating over slice...
	 */
	// var radius float64  = 30;
	// Println(circleAreaSlice(radius));

	newSlice := []int{1, 2, 3, 4, 5}
	for v := range newSlice {
		Println("**%d", v)
	}

	/**
	 *Switch and Slice Practice
	 */

	x := 6
	y := 7
	// 	slicee := []int{1,2,3}
	switch x {
	case 6:
		Println("Something")
		break
	default:
		Println("Wtf..")
		break
	}
	Println("Added Numbers..", addNumbers(x, y))
	/*	maxNumSlice := maxNum(slicee)
		Println("Finding Max Num in Slice : " , maxNumSlice)*/

	radius := 2.13
	Println(circleAreaSlice(radius))

	/**
	*Undecided number of args
	*See addMany()...
	 */
	Println(addMany(1, 3, 4, 5, 6, 7, 4, 3, 2, 1))

	/**
	*Gets copy of
	*Argument passed in
	*So wont affect
	*Actual Variable..
	 */

	var firstInt int = 3
	test(firstInt)

	Println("a = ", firstInt)

	/**
	**To Actualy Change it
	*We Use pointers.. just call &
	 */
	var secondInt int = 1
	test2(&secondInt)
	Println(secondInt)

	/**
	*FUNCTIONS AS TYPES
	 */

	type typeFunction func(c int) bool
	var dog animal
	dog.numOfLegs = 4
	dog.name = "Buchie"
	dog.id = 1.1
	Println("Dog", dog)

	//With Order...
	cat := animal{4, "Cattie", 30.2}
	Println("Cat", cat)
	//Without Order need to specify key:vale
	//And go will automatically reorder it...
	elephant := animal{numOfLegs: 4, name: "SmartOne", id: 5.6}
	Println("Elephant", elephant)
	/**
	 *Struct Example..
	 */
	Amanuel := person{"Amanuel", 15}
	Nos := person{"Nos", 21}
	Dad := person{"Dad", 42}
	Mom := person{"Mom", 41}

	personOlderOne, diffrenceOne, personYoungerOne := getStructDiffrenceTwo(Amanuel, Nos)
	Printf(" %s is older from %s  by %d years..\n", personOlderOne.name, personYoungerOne.name, diffrenceOne)

	personOlderTwo, diffrenceTwo, personYoungerTwo := getStructDiffrenceTwo(Dad, Mom)
	Printf(" %s is older from %s  by %d years..\n", personOlderTwo.name, personYoungerTwo.name, diffrenceTwo)

	oldestPersonAge := itterateStructDiffrenceTotal(Amanuel, Nos, Dad, Mom)
	oldestPersonName := personNameByAge(oldestPersonAge, Amanuel, Nos, Dad, Mom)
	Println("Oldest Person Name ", oldestPersonName)
	Printf("The oldest person is %s and is %d years old.. \n", oldestPersonName, oldestPersonAge)

	firstDonkey := Donkey{animal{numOfLegs: 3, name: "FirstDonkey", id: 3.134}, 78.12}
	AmanuelTheStudent := student{person{name: "Amanuel", age: 15}, "Programming"}

	AmanuelTheStudent.person = person{name: "Amanuel", age: 16}
	Println("Amanuel ", AmanuelTheStudent)

	Printf("The First Donkey %s and The first students %s \n", firstDonkey, AmanuelTheStudent)

	firstAlien := alien{kindaHuman{true, false}, true, true}

	//Structs multiple..
	Println("From KindaHuman Struct..", firstAlien.kindaHuman.hands)
	Println("From Alien Struct", firstAlien.hands)
	/**
	 *METHODS
	 */

	//Normal Method...
	r1 := Rectangle{3.212, 4.42}
	Println("Rectangular Area.. ", calculateRectngleStruct(r1))

	//** METHODS WITH STRUCTS!!
	c1 := Circle{9}
	Println("Circle Area.. ", c1.area())

	/**
	 *Box Method Exercice
	 */

	bBox := Box{2.12, 3.1, "red"}
	aBox := Box{1.23, 1.22, "Blue"}
	cBox := Box{3.32, 7.54, "blue"}
	Println("Orignal Color of b ", bBox.color)
	//Going to change actual
	//Color of b Box because i
	//used pointer on method

	bBox.setColor("White")
	Println("Changed color of b ", bBox.color)
	Println("The biggest area from all boxes is.. ", biggestAreaFromBox(aBox, bBox, cBox))
	Println("Of Color.. ", findColorFromAreaBox(biggestAreaFromBox(aBox, bBox, cBox), aBox, bBox, cBox))

	/**
	 *PANIC
	 */
	userPanicExample := ""
	panicExample(userPanicExample)

}

func printLine()                        { Println("Hello There... Im last for some reason..") }
func addNumbers(num1 int, num2 int) int { return num1 + num2 }

func test(a int) int {
	a = 5
	return a
}

func (b *Box) setColor(colorToSetTo string) {
	b.color = colorToSetTo
}
func (b Box) findArea() float64 {
	return b.width * b.height
}
func findColorFromAreaBox(area float64, b ...Box) string {
	var colorOfArea string = ""
	for i := 0; i <= len(b)-1; i++ {
		if (b[i].height * b[i].width) == area {
			colorOfArea = b[i].color
		}
	}
	return colorOfArea
}
func biggestAreaFromBox(b ...Box) float64 {
	Println(b)
	var biggestArea float64 = 0

	for i := 0; i <= len(b)-1; i++ {
		var currentWidth float64 = b[i].width
		var currentHeight float64 = b[i].height
		if (currentWidth * currentHeight) > biggestArea {
			biggestArea = (b[i].width * b[i].height)
		}
	}
	return biggestArea
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func getStructDiffrenceTwo(p1 person, p2 person) (person, int, person) {
	if p1.age > p2.age {
		return p1, (p1.age - p2.age), p2
	}
	return p2, (p2.age - p1.age), p1
}

func personNameByAge(age int, p ...person) string {
	var fakePersonName string
	for i := 0; i <= len(p)-1; i++ {
		if p[i].age == age {
			Println(p[i])
			fakePersonName = p[i].name
		}
	}
	return fakePersonName
}

func itterateStructDiffrenceTotal(p ...person) int {
	amount := 0
	for i := 0; i <= len(p)-1; i++ {
		if p[i].age > amount {
			amount = p[i].age
		}
	}
	return amount
}
func test2(a *int) int {
	*a = 5
	return *a

}
func panicExample(user string) {
	//Defer will login since above panic
	defer deferExample()
	if user == "" {
		panic("NO USER!")
	} else {
		Println(user)
	}

}

func calculateRectngleStruct(r Rectangle) float64 {
	return r.width * r.height
}
func deferExample() {
	Println("DEFERED!")
}
func circleAreaSlice(radius float64) float64 {
	const PI = 3.1456
	var areaValue float64
	areaValue = (PI * (radius * radius))
	return areaValue
}

/*func maxNum(num2 []int) int{
    var Max []int;
    Max[0] = 0;
	for i:=0; i<=len(num2)-1;i++{
	    Println(num2[i])
	    if(num2[i]>Max[0]){
	        Max[0] = num2[i]
	    }
	}
	return Max[0];
}*/
/**
 *Multi-Value Return..
 */
func addAndMultiply(a int, b int) (int, int) {
	return a + b, a * b
}
func addMany(arg ...int) int {
	//DONT FORGET arg is a array..
	//In this case a type int array
	added := 0
	for i := 0; i <= len(arg)-1; i++ {
		added = added + arg[i]
	}
	return added
}
