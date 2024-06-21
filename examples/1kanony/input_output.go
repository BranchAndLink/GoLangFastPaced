package main

import (
	"fmt"
	"os"
)

func main() {
	//////////////////////////////////////////////////////////////////
	// OUTPUT
	//////////////////////////////////////////////////////////////////

	fmt.Print("Salam", "Necesen?") // without \n
	fmt.Println("\nHello World")
	fmt.Printf("%d\n", 123 + 3433) // formatted + no new line by default

	fmt.Printf( "%f\n", 'A' ) // causes silent error msg

	emojiInNumberForm := '🙂'
	emoji := string(emojiInNumberForm)

	fmt.Println(emojiInNumberForm)
	fmt.Println(emoji)

	// manually setting stream
	fmt.Fprintf( os.Stdout, "Emoji for Stdout: %s\n", emoji )
	fmt.Fprintf( os.Stderr, "Emoji for Stderr: %s\n", emoji )

	// also we can simply use
	print( "So simple\n" )
	print( "Agreed :)\n" )


	//////////////////////////////////////////////////////////////////
	// INPUT
	//////////////////////////////////////////////////////////////////

	var(
		a int
		b string
		c string
	)

	fmt.Sscanf( "100 kere salam dostum", "%d %s salam %s", &a, &b, &c )
	fmt.Println( a, b, c )

	var d float32

	fmt.Sscan( "200 23.67", &a, &d )
	fmt.Println( a, d * 2.4 )

	str := "999 \n asdf"
	fmt.Sscanln( str, &a, &b ) // b won't change for it stops after first line finished
	fmt.Println( a, b )
	
	fmt.Scan( &b )
	fmt.Println( b )
}
