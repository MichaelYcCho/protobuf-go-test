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

	jsonBytes, _ := protojson.Marshal(&u)
	log.Println(string(jsonBytes))
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

func JsonToProtoUser() {
	json := `{"id":99, "username":"batman", "is_active":true, "password":"MTIzNA==", "emails":["batman@movie.com", "batman@dc.com"], "gender":"GENDER_MALE"}`

	var p basic.User
	err := protojson.Unmarshal([]byte(json), &p)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(&p)
}
