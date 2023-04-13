package bd

import (
	"github.com/BautistaBianculli/HelloFiber/models"
)


func Profile(id string) (models.User,bool , error){
	user := new(models.User)
	result:= PostgresDB.Where(map[string]interface{}{"id":id}).Find(&user)
	if result.Error != nil{
		return *user,false, result.Error  
	}
	if result.RowsAffected < 1 {
		return *user,false, result.Error
	}
	return *user,true, nil
}