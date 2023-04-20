package modelos

import "time" //todas las estructuras empiezan con type

type User struct {
	Id        int
	Name      string // el paquete modelo en Go se usa solamente para guardar definiciones y estructuras
	CreatedAt time.Time
	Status    bool
}

func (usuario *User) AddUser(id int, name string, createdAd time.Time, status bool) {
	usuario.Id = id
	usuario.Name = name
	usuario.CreatedAt = createdAd
	usuario.Status = status

}
