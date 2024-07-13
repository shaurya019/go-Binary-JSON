// bson.A: This represents a BSON array as a slice of interface{}.

package main

import (
    "fmt"
    "go.mongodb.org/mongo-driver/bson"
)

func main() {
    // Creating a BSON array using bson.A
    arr := bson.A{"John Doe", 30, bson.M{"street": "123 Main St", "city": "New York"}}

    // Printing the array
    fmt.Println(arr)
}
