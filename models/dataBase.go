package models

import "strconv"

type ConnectionBD struct {
	Host 		string 
	User 		string
	Password 	string
	DbName 		string
	Port 		int
}

type DataBase interface{
	Credenciales()
}

func (bd *ConnectionBD)Init(){
	bd.Host = "localhost"
	bd.User = "postgres"
	bd.Password = "admin123"
	bd.DbName = "Gestion"
	bd.Port = 5432
}


func (con *ConnectionBD)Credenciales() string {
	bd := new(ConnectionBD)
	bd.Init()
	return "host=" + bd.Host + " user=" + bd.User + " password=" + bd.Password + " dbname=" + bd.DbName + " port=" + strconv.Itoa(bd.Port)
}