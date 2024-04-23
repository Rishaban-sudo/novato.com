package main

import (
	"fmt"

	//"novato.com/articles"

	"novato.com/json"
)

func main() {
	// fmt.Println("Hello, world!")
	// fmt.Print("Test name: " + TestName + "\n")
	// fmt.Println("Hello variable : ", TestName)

	// fmt.Println(articles.Article1)

	// fmt.Println(articles.GetArticleText())

	var jsonObj *json.JsonObject
	fmt.Println(jsonObj)
	fmt.Println(jsonObj.GetJsonRepresentation())

	jsonObj = json.NewJsonObject()

	fmt.Println("After creating plain print  ->", jsonObj)
	fmt.Println("After creating print  ->>>", jsonObj.GetJsonRepresentation())

	jsonObj.AddOrUpdate("isSaturated", "true")
	jsonObj.AddOrUpdate("isOffline", "false")

	fmt.Println(jsonObj.GetJsonRepresentation())

	fmt.Println(jsonObj.Remove("isSaturated"))

	fmt.Println(jsonObj.GetJsonRepresentation())
}
