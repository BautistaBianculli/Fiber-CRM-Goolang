package bd

import(
	"github.com/BautistaBianculli/HelloFiber/models"
	// "github.com/google/uuid"
)


func Login(user *models.User) (bool, error){

	result := PostgresDB.Where(map[string]interface{}{"user":user.User,"password":user.Password}).Find(&user)
	if result.Error != nil{
		return false, result.Error
	}
	if result.RowsAffected < 1 {
		return false, nil
	}
	return true, nil
}