package database

import "github.com/iwanhae/Jupynetes/ent"

//GetClient return database client
func GetClient() *ent.Client {
	return client
}
