// BSON (Binary JSON) is a binary-encoded serialization of JSON-like documents. It's the primary data format used by MongoDB to store documents and make remote procedure calls.


// bson.E
// bson.E represents a BSON element, which is a key-value pair. It's defined as:


// type E struct {
// 	Key string
// 	Value interface{}
// }


package main

import (
	"fmt"
    "go.mongodb.org/mongo-driver/bson"
)

func main () {
	ele := bson.E{Key:"name",Value: "John Doe"}

	fmt.Println(ele)
}