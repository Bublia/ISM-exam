/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
package main


import "fmt"
import "strconv"
import "math"

func main() {
    var cash string
    var d float64 
    var coins int
    coins = 0
    fmt.Println("Enter the amount of cash you have: ")
    fmt.Scanf("%s", &cash)
    strconv.ParseFloat(cash, 64)
    
    var cts int = int (cash)
    cts := math.Round(cash)
    
}
