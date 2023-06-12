package basic

import (
	"log"
	"my-protobuf/protogen/basic"
)

func BasicHello() {
	h := basic.Hello{
		Name: "Michael",
	}

	// log.Println(&h)를 호출하면, Hello 타입의 String() 메서드가 호출되어 Hello 메시지의 문자열 표현을 반환하게 됩니다.
	log.Println(&h)
}
