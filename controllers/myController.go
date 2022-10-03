package controllers

import (
	helper "BiletAlSatArkaUc/helpers"
	"BiletAlSatArkaUc/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Load_money() gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err_t := c.Cookie("token")
		var money models.Money

		if err_t != nil {
			c.JSON(500, gin.H{"error": err_t})

		}

		claims, msg := helper.ValidateToken(token)

		if msg != "" {

		}

		if claims.User_type != "ADMIN" {

			c.JSON(500, gin.H{
				"Error": "Bu işlem için yetkiniz yoktur lütfen giriş yapmayı deneyin",
			})

			c.JSON(500, gin.H{
				"Error": c.GetString("User_type"),
			})
			return
		}

		if err := c.BindJSON(&money); err != nil {
			c.JSON(500, err)
			return
		}

		validationErr := validate.Struct(&money)

		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_validate": validationErr.Error()})
			return
		}
		c.JSON(500, money.Amount_toload)

		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User

		result := userCollection.FindOne(ctx, bson.M{
			"email": claims.Email,
		}).Decode(&user)
		c.JSON(200, gin.H{"Hesabınızdaki para miktarı": user})

		if result != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_Result": result})

		}

		if user.Money == 0 {
			user.Money = money.Amount_toload
		} else {
			user.Money += money.Amount_toload
		}

		defer cancel()

	}
}

func My_ticket() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")

		if err != nil {
			c.JSON(500, gin.H{"error": "token hatalı"})
			return
		}

		claims, msg := helper.ValidateToken(token)

		if msg != "" {
			c.JSON(200, gin.H{"error": msg})
			return
		}

		if claims.User_type != "ADMIN" || claims.User_type != "USER" {
			c.JSON(200, gin.H{"error": "ADMIN veya USER kullanıcı tipine sahip değilsiniz"})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User

		result := userCollection.FindOne(ctx, bson.M{
			"email": claims.Email,
		}).Decode(&user)

		if result != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_Result": result})

		}

		defer cancel()

		for _, i := range user.Ticket_Array {
			c.JSON(200, "Biletleriniz aşşagıda listelenmektedir")
			c.JSON(200, i)
		}
	}

}

func My_Profil() gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := c.Cookie("token")
		var user models.User

		if err != nil {
			c.JSON(400, gin.H{
				"message": "tokeninizin süresi dolmus,token dogrulanamadı",
			})
			return
		}

		if err != http.ErrNoCookie {
			c.JSON(404, gin.H{
				"message": "cokkie bulunamadı",
			})
			return
		}
		claims, msg := helper.ValidateToken(token)

		if msg != "" {
			c.JSON(500, msg)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Millisecond)
		filter := bson.D{{"email", claims.Email}}

		resulterr := userCollection.FindOne(ctx, filter).Decode(&user)
		defer cancel()

		if resulterr != nil {
			return
		}
		c.JSON(200, user)
	}
}

func Updateitem() gin.HandlerFunc {
	return func(c *gin.Context) {
		//var query_params = []string{"firstname", "lastname", "email", "password"}

		query := c.Request.URL.Query()
		var ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)

		firstname := query["firstname"]
		lastname := query["lastname"]
		email := query["email"]
		password := query["password"]

		filter := bson.D{
			"set":bson.D{
			"firstname"	
			}
		}

		userCollection.UpdateMany(ctx)
		defer cancel()
	}
}
