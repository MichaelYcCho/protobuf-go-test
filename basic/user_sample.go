package basic

import (
	"log"

	"github.com/MichaelYcCho/protobuf-go/protogen/basic"
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
