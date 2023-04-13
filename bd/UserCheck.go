package bd

import (
	"github.com/BautistaBianculli/HelloFiber/models"
)


func UserCheck(claims *models.JWTClaims)(models.User, bool){
	user := new(models.User)
	result:= PostgresDB.Where(map[string]interface{}{"id":claims.ID}).Find(&user)
	if result.Error != nil{
		return *user,false
	}
	if result.RowsAffected < 1 {
		return *user, false 
	}

	return *user, true
}