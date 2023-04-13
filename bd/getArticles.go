package bd

import (
	"github.com/BautistaBianculli/HelloFiber/models"
)



func GetArticles()([]models.Article,error){
	var article []models.Article
	result := PostgresDB.Find(&article)
	if result.Error!= nil{
		return article,result.Error
	}
	return article, nil
	
}