package bd

import (
	"github.com/BautistaBianculli/HelloFiber/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var PostgresDB = ConnectionBD()

func ConnectionBD() *gorm.DB{
	con := new(models.ConnectionBD)
	db, err := gorm.Open(postgres.Open(con.Credenciales()), &gorm.Config{})
	if err != nil{
		log.Fatal(err.Error())
		return  db
	}
	dbInstance , err := db.DB()
	err = dbInstance.Ping()
	if err != nil{
		log.Fatal(err.Error())
		return db
	}
	log.Println("Conexion Creada")
	return db
}

func ConenectionCheck() (bool){
	dbInstance, err := PostgresDB.DB()
	if err != nil{
		log.Fatal(err.Error())
		return false
	}
	err = dbInstance.Ping()
	if err != nil{
		log.Fatal(err.Error())
		return false
	}
	return true
}


func CloseConnection(db *gorm.DB)(bool, error) {
	dbInstance, err := db.DB()
	if err != nil{
		return false, err
	}
	err = dbInstance.Close()
	if err != nil{
		return false, err
	}
	return true, nil
}