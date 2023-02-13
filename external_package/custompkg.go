package main

import (
	Sub "ext/sub_external_package"
	"fmt"
)

func main() {
	Sub.SubExternalPrint()
	fmt.Println("Hello")
}
