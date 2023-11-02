package service

type Repositorio interface {
	findUser(id int) User
	findUsers() []User
}
