package sub_external_package

import (
	"fmt"

	"github.com/guptarohit/asciigraph"
)

// sub_external_package
func SubExternalPrint() {
	fmt.Println("This is sub external package")

	//expkg.PrintSample()
	data := []float64{3, 4, 5, 6, 9, 7, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}
