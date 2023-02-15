package main

import (
	Sub "ext/sub_external_package"
	"fmt"
)

func main() {
	Sub.SubExternalPrint()
	fmt.Println("Hello")
	// 정리하면, 상위 디렉토리에서 go.mod 를 만들고 (go mod init '모듈명'), 해당 모듈명으로 하위 패키지를 임포트하여 사용할 수 있다.

	Sub.PrintD()
}
