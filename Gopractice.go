// // package main

// // import "fmt"

// // func main() {
// // 	fmt.Println("Hello, Go!")
// // }

package main
// // import (
// // 		"fmt"
// // 		"math"
// // 		"strings"
// // )
// // func main() {
// // 	fmt.Println(math.Floor(2.75))
// // 	fmt.Println(strings.Title("head first go"))
// // }

// package main

// import (
// 	"fmt"
// 	"reflect"
// 	"sort"
// 	"time"
// )

// func main() {
// 	var now time.Time = time.Now()
// 	var year int = now.Year()
// 	fmt.Println(year)
// 	fmt.Println(now)
// 	fmt.Println(reflect.TypeOf(now))
// 	fmt.Printf("%T\n", now)

// 	sort.Ints([]int{})

// 	// laptop("Hp", []string{"type a", "type c"}, 540_000, "15 inches")

// 	// agbero(laptop)
// 	// chivita(500)

// 	humanbeing(chivita)
// }

// // func laptop(name string, charger []string, price int, screensize string) {
// // 	fmt.Printf("the name of your laptop is %v. The laptop is charging using %v charger. The price is %v naira with screensize %v.", name, len(charger), price, screensize)
// // }

// // func samSung(charger string) {
// // 	fmt.Printf("The name of the charger is %v.", charger)
// // }

// // func agbero(laptop func(name string, charger []string, price int, screensize string)) {
// // 	laptop("Hp", []string{"type a", "type c"}, 540_000, "15 inches")
// // }

// /**
// function 1 chivita (price int)
// function2 humanbeing(drink func(price int))
// */

// func chivita(price int) {
// 	fmt.Printf("the price of the chivita is %v naira.", price)
// }

// func humanbeing(drink func(price int)) {
// 	drink(100)
// }
