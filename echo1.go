//Echo1 prints its command line arguments
//Brett
package main

import (
    "fmt"
    "os"
)

func main() {

    fmt.Println( "The name of the program is " + os.Args[0] );

    for i := 1; i < len(os.Args); i++ {
        fmt.Printf( "%d %s\n", i, os.Args[i] );
    }
}
