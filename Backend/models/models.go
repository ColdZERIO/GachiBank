package models

type UserAuthDB struct {
	Login    string `bson:"login"`
	Password string `bson:"password"`
}
