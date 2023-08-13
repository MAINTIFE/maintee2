package main

import (
	"fmt"
	"strings"
)

func main() {
	itemsAsSlice := []string{"boot", "top", "trouser", "gown", "jumpsuit"}
	// itemsAsString(itemsAsSlice)
	result := itemsAsString2(itemsAsSlice)
	fmt.Println(result)
	// wishes(wishitems)

	// engine1 := buildEngine("v8")

	// engine2 := buildEngine("v12")

	// fmt.Println("Engine 1", engine1)

	// fmt.Println("Engine 2", engine2)

	// itemsString := convertItemsToString(wishitems)
	// fmt.Println(itemsString)
	// result := strings.Join(wishitems, ",")
	// fmt.Println(result)
	// "boot, top, trouser, gown, jumpsuit"
}

// func wishes(item []string) {
// 	for _, item := range item {
// 		fmt.Printf("I will get a %v.\n", item)
// 	}
// 	fmt.Println("Peace Out!")
// }

// func buildEngine(engineType string) string {
// 	builtEngine := "Engine has been built " + engineType
// 	return builtEngine
// }

// func agbero(insult string) string {
// 	return fmt.Sprintf("Slap!....because you said %v", insult)
// }

/**
create a function called wishes that accepts a slice of items
loop through the items and print
 out I will get a + {item}
then outside the loop, print out "Peace Out!"
*/
// {trouser, jumpsuit, top, boot, gown}


// func itemsAsString(itemsAsSlice []string) {
// 	result := strings.Join(itemsAsSlice, ", ")
// 	fmt.Println(result) 
// }
func itemsAsString2(itemsAsSlice []string) string{
	return strings.Join(itemsAsSlice, ", ")
}