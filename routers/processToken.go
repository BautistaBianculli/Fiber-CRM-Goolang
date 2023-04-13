package routers

import (
	"errors"
	"strings"
	"time"

	"github.com/BautistaBianculli/HelloFiber/bd"
	"github.com/BautistaBianculli/HelloFiber/models"
	jwt "github.com/dgrijalva/jwt-go"
)




func VerifyToken(tkn string)error {
	clave := []byte("adruid123")
	claims := new(models.JWTClaims)
	tkn = strings.ReplaceAll(tkn," ","")
	token := strings.Split(tkn, "Bearer")
	if len(token) != 2{
		return errors.New("Invalidtoken")
	}

	tk := strings.TrimSpace(token[1])
	_, err := jwt.ParseWithClaims(tk,claims,func(token *jwt.Token)(interface{}, error){
		return clave,nil
	})
	if err == nil{
		if time.Now().After(time.Unix(int64(claims.ExpiresAt),0)){
			return errors.New("ExpiredToken")
		}
		_, status := bd.UserCheck(claims)
		if !status {
			return errors.New("UnhautorizedUser")
		}
		return nil
	}
	return err

}