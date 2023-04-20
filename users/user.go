package user

import (
	"fmt"
	"time"

	"github.com/Marian86029/godesde0/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "Mariano", time.Now(), true)
	fmt.Println(u)

}
