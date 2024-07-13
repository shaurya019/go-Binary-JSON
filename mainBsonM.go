// bson.M
// bson.M is a map representation of a BSON document. It's defined as:

type M map[string]interface{}

package main

import (
    "fmt"
    "go.mongodb.org/mongo-driver/bson"
)

func main() {

	doc := bson.M{
		"name": "John Doe",
		"age":30,
		"address":bson.M{
			"street": "23",
            "city":   "India",
		}
	}
	fmt.Println(doc)
}