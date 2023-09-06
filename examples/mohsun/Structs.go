package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Declaration of a Struct Type

type circle struct {
	length float64
	radius float64
	color  string
}

type rectangle struct {
	length  int
	breadth int
	color   string

	geometry struct {
		area      int
		perimeter int
	}
}

type Salary struct {
	Basic, HRA, TA float64
}

type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

// Field tags
type newEmployee struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	City      string `json:"city"`
}

// Add Method to Struct Type
func (e Employee) EmpInfo() string {
	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)
	for _, info := range e.MonthlySalary {
		fmt.Println("===================")
		fmt.Println(info.Basic)
		fmt.Println(info.HRA)
		fmt.Println(info.TA)
	}
	return "----------------------"

}

//ask about runnning sequence.

type Employees struct {
	Name string
	Age  int
}

func (obj *Employees) Info() {
	if obj.Name == "" {
		obj.Name = "John Doe"
	}
	if obj.Age == 0 {
		obj.Age = 25
	}
}

func main() {
	fmt.Println(circle{10.5, 25.10, "red"})
	//The var keyword initializes a variable rect. Using dot notation, values are assigned to the struct fields.
	var rect rectangle // dot notation
	rect.length = 10
	rect.breadth = 20
	rect.color = "Green"

	rect.geometry.area = rect.length * rect.breadth
	rect.geometry.perimeter = 2 * (rect.length + rect.breadth)

	fmt.Println(rect)
	fmt.Println("Area:\t", rect.geometry.area)
	fmt.Println("Perimeter:", rect.geometry.perimeter)

	var cir1 = circle{10, 20, "Green"}
	fmt.Println(cir1)

	var cir2 = circle{length: 10, color: "Green"} // breadth value skipped
	fmt.Println(cir2)

	cir3 := circle{10, 20, "Green"}
	fmt.Println(cir3)

	cir4 := circle{length: 10, radius: 20, color: "Green"}
	fmt.Println(cir4)
	cir4.length = 20
	fmt.Println("value of cir4:", cir4)

	cir5 := circle{radius: 20, color: "Green"} // length value skipped
	fmt.Println(cir5)
	//An instance of a struct can also be created with the new keyword. It is then possible to assign data values to the data fields using dot notation.

	circ1 := new(circle) // circ1 is a pointer to an instance of circle
	circ1.length = 10
	circ1.radius = 20
	circ1.color = "Green"
	fmt.Println(circ1)

	var circ2 = new(circle) // circ2 is an instance of circle
	circ2.length = 10
	circ2.color = "Red"
	fmt.Println(circ2)

	var circle1 = &circle{10, 20, "Green"} // Can't skip any value
	fmt.Println(circle1)

	var circle2 = &circle{}
	circle2.radius = 10
	circle2.color = "Red"
	fmt.Println(circle2)

	var circle3 = &circle{}
	(*circle3).radius = 10
	(*circle3).color = "Blue"
	fmt.Println(circle3)
	fmt.Println(circle3.radius > cir4.length)
	var cc = circle{10, 20, "green"}
	var c1 = cc
	c1.radius = 7
	fmt.Println(cc, c1)

	//Nested Struct Type
	//Struct can be nested by creating a Struct type using other Struct types as the type for the fields of Struct.
	//Nesting one struct within another can be a useful way to model more complex structures.

	e := Employee{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []Salary{
			Salary{
				Basic: 15000.00,
				HRA:   5000.00,
				TA:    2000.00,
			},
			Salary{
				Basic: 16000.00,
				HRA:   5000.00,
				TA:    2100.00,
			},
			Salary{
				Basic: 17000.00,
				HRA:   5000.00,
				TA:    2200.00,
			},
		},
	}
	fmt.Println("name:", e.FirstName, e.LastName)
	fmt.Println("age:", e.Age)
	fmt.Println("email:", e.Email)
	fmt.Println("salary:", e.MonthlySalary[0])
	fmt.Println("salary:", e.MonthlySalary[1])
	fmt.Println("salary:", e.MonthlySalary[2])

	//During the definition of a struct type, optional string values may be added to each field declaration.

	json_string := `
    {
        "firstname": "Rocky",
        "lastname": "Sting",
        "city": "London"
    }`

	emp1 := new(newEmployee)
	json.Unmarshal([]byte(json_string), emp1)
	fmt.Println(*emp1)

	emp2 := new(newEmployee)
	emp2.FirstName = "Ramesh"
	emp2.LastName = "Soni"
	emp2.City = "Mumbai"
	jsonStr, _ := json.Marshal(emp2)
	fmt.Printf("%s\n", jsonStr)

	//running the method for the struct.
	fmt.Println("Method result:")
	fmt.Println(e.EmpInfo())

	employee1 := Employees{Name: "Mr. Fred"}
	employee1.Info()
	fmt.Println(employee1)

	employee2 := Employees{Age: 26}
	employee2.Info()
	fmt.Println(employee2)

	//find type of struct

	fmt.Printf("type of cir1: %T \n", cir1)   // main.circle
	fmt.Println(reflect.ValueOf(cir1).Kind()) // struct

	fmt.Printf("type of cir2: %T \n", cir2)   // main.circle
	fmt.Println(reflect.ValueOf(cir2).Kind()) // struct

	fmt.Printf("type of circle2: %T \n", circle2) // *main.circle
	fmt.Println(reflect.ValueOf(circle2).Kind())  // ptr

	// fmt.Println(reflect.TypeOf(circle3))         // *main.circle
	fmt.Printf("type of circle3: %T \n", circle3)

	//Comparing Structs
	rect3 := new(rectangle)
	var rect4 = &rectangle{}

	// rect3 := circle{10, 20, "Green"}
	// rect4 := circle{10, 20, "Green"}

	if rect3 == rect4 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
	fmt.Println(&rect3, &rect4)

	//embedding

	Ystack := branchandlink{
		number_of_collaborators: 3,
		our_project:             project{learnprogramming: "golang"},
	}

	var ystack branchandlink
	ystack.number_of_collaborators = 3
	ystack.our_project.learnprogramming = "kubernetes"

	fmt.Println(Ystack)
	fmt.Println(ystack)

}

type project struct {
	learnprogramming string
}

type branchandlink struct {
	number_of_collaborators int
	our_project             project
}
