package bd

import (
	"github.com/BautistaBianculli/HelloFiber/models"
	"github.com/google/uuid"

)


func CreateUser(user models.User) (bool, error){
	user.ID = uuid.NewString()
	result:= PostgresDB.Create(&user)
	if result.Error != nil{
		return false, result.Error
	}
	return true, nil
}