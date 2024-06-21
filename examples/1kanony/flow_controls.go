package main

import "fmt"

func main() {
	////////////////////////////////////////////
	// Scopes
	////////////////////////////////////////////

	var someVar = 'A'

	{ // brand-new scope
		someVar = 'E' // overwrite top scope someVar

		var someVar = "123" // brand-new variable with the same name
		println( someVar )
	}

	println( someVar )


	////////////////////////////////////////////
	// If Statements
	////////////////////////////////////////////

	if a := true; a == true {
		println( "Correct!" )
	}

	if false {
		// nope
	} else if true {
		println( "Yes!" );
	} else {
		// nope
	}


	////////////////////////////////////////////
	// Switch-case Statement
	////////////////////////////////////////////

	someRandomValue := 200

	switch {
		case someRandomValue == 100: println( "100" )
		case someRandomValue == 200: println( "200" )
		case someRandomValue == 300: println( "300" )
		default: print( "Default" )
	}

	// or

	switch someRandomValue {
	case 100: println( "100" )
	case 200: println( "200" )
	case 300: println( "300" )
	default: print( "Default" )
	}

	startDay := 3

	switch startDay {
	case 1: println( "Monday" ); fallthrough
	case 2: println( "Tuesday" ); fallthrough
	case 3: println( "Wednesday" ); fallthrough
	case 4: println( "Thursday" ); fallthrough
	case 5: println( "Friday" ); fallthrough
	case 6: println( "Saturday" ); fallthrough
	case 7: println( "Sunday" )
	default: println( "None" )
	}

	var i interface{} = func(...int) int { return 100 * 100 };

	switch i.(type) {
	case int: println( "INT" )
	case string: println( "STRING" )
	case func(...int) int: println( "FUNC" )
	}


	////////////////////////////////////////////
	// Loops
	////////////////////////////////////////////

	for i := 0; i < 5; i++ {
		println( i )
	}

	var j int

	for { // infinite loop
		fmt.Scanf( "%d", &j )

		println( j )

		if j == 0 {
			break
		}
	}

	for i, char := range "salamm" {
		fmt.Printf( "%d: %s\n", i, string(char) )
	}

	rootFor:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				if i > 7 && j == k {
					println( "BREAKING" )
					break rootFor
				} else if i > 5 && j < k {
					println( "CONTINUING" )
					continue rootFor
				}
			}
		}
	}

	////////////////////////////////////////////
	// GOTO
	////////////////////////////////////////////

	goto theEnd
	println( "Heyy" )
	theEnd:
}

