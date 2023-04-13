package bd


import (
	"github.com/BautistaBianculli/HelloFiber/models"
	"github.com/google/uuid"
	"math"
)




func CreateArticle(article *models.Article)(bool, error){
	article.ID = uuid.NewString()
	article.Prventa = math.Round((article.PrecioLista * (1 - article.Bonif1/100) * (1 - article.Bonif2/100) * 1.21)*100)/100
	result := PostgresDB.Create(&article)
	if result.Error != nil{
		return false, result.Error
	}
	return true, nil
}