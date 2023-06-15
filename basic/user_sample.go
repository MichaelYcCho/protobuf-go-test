package basic

import (
	"log"

	"github.com/MichaelYcCho/protobuf-go/protogen/basic"
	"google.golang.org/protobuf/encoding/protojson"
)

func BasicUser() {
	u := basic.User{
		Id:       99,
		Username: "superman",
		IsActive: true,
		Password: []byte("1234"),
		Emails:   []string{"superman@movie.com", "superman@dc.com"},
		Gender:   basic.Gender_GENDER_MALE,
	}

	log.Println(&u)
}

func ProtoToJsonuser() {
	p := basic.User{
		Id:       99,
		Username: "batman",
		IsActive: true,
		Password: []byte("1234"),
		Emails:   []string{"batman@movie.com", "batman@dc.com"},
		Gender:   basic.Gender_GENDER_MALE,
	}
	jsonBytes, _ := protojson.Marshal(&p)
	log.Println(string(jsonBytes))
}
