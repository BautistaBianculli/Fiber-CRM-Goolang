package jwt

import(
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/BautistaBianculli/HelloFiber/models"
)



func GenerateJWT(user models.User)(string,error){

	clave := []byte("adruid123")

	payload := jwt.MapClaims{
		"ID"		:	user.ID,
		"user"		: 	user.User,
		"password"	:	user.Password,
		"FirstName"	:	user.FirstName,
		"LastName"	:	user.LastName,
		"exp"		:	time.Now().Add(time.Hour * 24).Unix(),
	}
	
	token:= jwt.NewWithClaims(jwt.SigningMethodHS256,payload)

	tokenStr, err := token.SignedString(clave)
	if err != nil{
		return tokenStr, err
	}

	return tokenStr, nil
}