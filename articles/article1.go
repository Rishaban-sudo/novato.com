package articles

import (
	"fmt"
)

var Article1 = "The mountain gave birth to a mouse :)"

func GetArticleText() string {
	fmt.Println("Getting Article text form articles package...")
	return Article1
}
