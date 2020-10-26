package domain

//User respresents a typical person
type User struct {
	ID             int
	Name           string
	PatronymicName string
	FamilyName     string
	Roles          []Role
}

//Role represents role
type Role struct {
	Name string
}
