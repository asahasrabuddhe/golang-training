package main

import (
	"fmt"
	"github.com/anandnkhl/training/04_scope/04_visibility/mypackage"
)

func main() {
	fmt.Println(mypackage.MyName)
	mypackage.PrintCompanyName()
}
