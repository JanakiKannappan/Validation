package input
import (
	"fmt"
	"strings"
)



func Validation(name string) {
		
		
	if strings.Contains(name, " ") {
		fmt.Println("Error: Name should not contain whitespace.")
	} else {
		fmt.Println("Name is valid.")
	}

}