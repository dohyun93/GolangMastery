package main

import (
	_ "fmt" // 패키지를 사용하지 않을 때, _ 로 앨리어싱을 주면 된다.
)

func main() {
	// 1.16 버전부터 Go 모듈 사용이 기본이 되었다.
	// 이전 까지는 Go 모듈을 만들지 않는 프로젝트는 모두 GOPATH/src 폴더 아래에 있어야 했지만,
	// Go 모듈이 기본이 되면서 모든 Go 코드는 Go 모듈 아래에 있어야 한다.

	// go build 를 하려면 반드시 Go 모듈 루트 폴더에 go.mod 파일이 있어야 한다.

	// go.mod : 모듈 이름, Go 버전, 필요한 외부 패키지명
	// go.sum : 외부 패키지 버전
	// 위 두 가지가 있을 때 go build 를 하면 실행파일이 만들어진다.

	// Go 모듈은 'go mod init [패키지명]' 명령을 통해 만들 수 있다.
}
