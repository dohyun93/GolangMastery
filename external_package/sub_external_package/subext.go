package sub_external_package

import (
	"fmt"

	"github.com/guptarohit/asciigraph"
)

// 1. 패키지 임포트 시, 먼저 전역변수를 초기화 한다.
// 따라서 f 호출이 두 번 일어나서 d는 5가 된다. b 는 4, c 는 5, a 는 9가 된다. (d는 5)
var (
	a = c + b
	b = f()
	c = f()
	d = 3
)

func f() int {
	d++
	fmt.Println("f() d:", d)
	return d
}

// 2. 그 다음, init 이 존재하므로 호출된다.
// d 는 5였기 때문에, 6이 출력된다.
func init() {
	// init() 은 패키지가 import 될 때 존재하는경우 호출되어 패키지를 초기화 해준다.
	d++
	fmt.Println("Init function", d)
}

// 3. 상위 패키지의 로직들이 돌고 호출된 아래 함수로 d 의 값이 출력된다.
func PrintD() {
	fmt.Println("d: ", d)

}

// sub_external_package
func SubExternalPrint() {
	fmt.Println("This is sub external package")

	//expkg.PrintSample()
	data := []float64{3, 4, 5, 6, 9, 7, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}
