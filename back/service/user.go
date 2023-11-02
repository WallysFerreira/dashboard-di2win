package service

type User struct {
	id      int
	name    string
	segment string
}

func (rp *RepositorioPostgre) findUser(id int) User {
	return User{}
}
