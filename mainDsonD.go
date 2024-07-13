// bson.D: This represents a BSON document as a slice of bson.E elements. It's useful when order matters.


package main

import (
    "fmt"
    "go.mongodb.org/mongo-driver/bson"
)

func main() {
    // Creating a BSON document using bson.D
    doc := bson.D{
        {Key: "name", Value: "John Doe"},
        {Key: "age", Value: 30},
        {Key: "address", Value: bson.D{
            {Key: "street", Value: "123 Main St"},
            {Key: "city", Value: "New York"},
        }},
    }

    // Printing the document
    fmt.Println(doc)
}
