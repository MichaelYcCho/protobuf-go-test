package main

import (
	"fmt"
	"log"
	"my-protobuf/basic"
	"time"
)

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().Format("15:04:05") + " " + string(bytes))
}

/*
logWriter 타입은 이 Write 메서드를 구현하므로 io.Writer 인터페이스를 충족합니다.
이 경우, log.SetOutput(new(logWriter)) 코드는 로깅 출력 대상을 새로운 logWriter 인스턴스로 설정합니다.
그런 다음, log.Println()과 같은 로그 작성 함수가 호출될 때마다
log 패키지는 내부적으로 이 Write 메서드를 호출하여 로그 메시지를 씁니다.
이 Write 메서드에서는 로그 메시지에 현재 시간을 추가하고, 이를 표준 출력으로 전달하여 화면에 출력합니다.
따라서 이 Write 메서드는 직접 호출되지는 않지만, 로그 메시지를 쓰는 데 사용됩니다.
여기서는 로그 메시지에 현재 시간을 추가하는 역할을 합니다.
*/

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	//basic.BasicHello()
	basic.BasicUser()
}
