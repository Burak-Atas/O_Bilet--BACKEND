package controllers

import (
	"BiletAlSatArkaUc/database"
	"BiletAlSatArkaUc/models"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Ticket_Collection mongo.Collection = *database.OpenCollection(database.Client, "tickets", "ticket")

func Query_Tickets() gin.HandlerFunc {
	return func(c *gin.Context) {

		/*
				//where from
				wf := c.Query("wf")
				//to where
				tw := c.Query("tw")

				//start date
				s_dt := c.Query("st")
				//finish date
				f_dt := c.Query("ft")


				filter := bson.D{{"where_from", wf}, {"tp_where", tw}, {"start_date", s_dt}, {"finish_date", f_dt}}
				var ticket []models.Ticket

			var ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)

			cursor, err := Ticket_Collection.Find(ctx, bson.M{})
			defer cancel()

			if err != nil {
				c.JSON(200, gin.H{"error": "listelenilecek bilet bulunamadı."})
				return
			}

			err = cursor.All(ctx, ticket)

			if err != nil {

			}
			c.JSON(200, gin.H{"error": "listelenilecek bilet bulunamadı."})

			c.JSON(200, len(ticket))
		*/
		filter := bson.D{}
		opts := options.Find().SetSort(bson.D{{"rating", 1}})
		cursor, err := Ticket_Collection.Find(context.TODO(), filter, opts)
		var results []bson.D
		if err = cursor.All(context.TODO(), &results); err != nil {
			panic(err)
		}
		for _, result := range results {
			c.JSON(200, result)
		}

	}
}

func Query() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("/query/:id")
		log.Println(id)
		var bus models.Bus
		if err := c.BindJSON(&bus.Buy_Ticket); err != nil {
			return
		}

		validationErr := validate.Struct(bus.Buy_Ticket)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		filter := bson.D{{"_id", id}}

		var ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)

		err := Ticket_Collection.FindOne(ctx, filter).Decode(&bus)
		defer cancel()
		for i := 0; i < bus.Koltuk_sayisi; i++ {
			if bus.Dolu_Koltuklar[i] == bus.Buy_Ticket.Koltuk_numarasi {
				c.JSON(200, gin.H{
					"message": "seçtiğiniz koltuk satın alınmış!!Lütfen başka koltuk seciniz",
				})
				return
			}
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		c.JSON(200, bus)
	}
}

func Cancel() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

		var ticket models.Cancel_Ticket
		defer cancel()

		if err := c.BindJSON(ticket); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		filter := bson.D{{"bilet_id", ticket.Bilet_İd}}

		result := Ticket_Collection.FindOne(ctx, filter).Decode(&ticket)

		if result != nil {

			return
		}

		c.JSON(200, ticket)

		if !ticket.Valid {
			c.JSON(200, gin.H{"message": "Biletiniz iptal edilmiştir"})
		}

	}
}

func Sefer_Ekle() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)

		defer cancel()
		var seferekle models.Bus

		if err := c.BindJSON(&seferekle); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"errorBindjson": err})
			return
		}
		validationErr := validate.Struct(seferekle)

		if validationErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": validationErr})
			return
		}

		seferekle.ID = primitive.NewObjectID()
		seferekle.Bos_koltuk_sayisi = seferekle.Koltuk_sayisi - len(seferekle.Dolu_Koltuklar)

		resultInsertionNumber, insertErr := Ticket_Collection.InsertOne(ctx, seferekle)

		if insertErr != nil {
			msg := fmt.Sprintf("User item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			c.JSON(http.StatusInternalServerError, gin.H{"error": insertErr.Error()})

			return
		}
		defer cancel()
		c.JSON(http.StatusOK, resultInsertionNumber)
	}
}

func Sefer_Güncelle() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func Sefer_Sil() gin.HandlerFunc {
	return func(c *gin.Context) {

		var delete models.Cancel_Ticket
		var ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)

		if err := c.BindJSON(&delete); err != nil {

		}

		validationErr := validate.Struct(delete)

		if validationErr != nil {

		}
		filter := bson.M{"bilet_id": delete.Bilet_İd}
		result, err := Ticket_Collection.DeleteOne(ctx, filter)
		defer cancel()
		if err != nil {

		}
		c.JSON(200, result)
	}
}
