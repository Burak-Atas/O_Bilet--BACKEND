package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ticket struct {
	ID          primitive.ObjectID `bson:"_id"`
	Where_from  string             `json:"where_from" `
	To_where    string             `json:"to_where" `
	Start_date  string             `json:"start_date" `
	Finish_date string             `json:"finish_date" `
	Bilet_id    string             `json:"bilet_id" `
}

type Bus struct {
	ID                primitive.ObjectID `bson:"_id" `
	Company_Name      string             `json:"company_name" validate:"required"`
	Koltuk_sayisi     int                `json:"koltuk_sayisi" validate:"required"`
	Dolu_Koltuklar    []int              `json:"dolu_koltuklar"`
	Yolculuk_saati    string             `json:"yolculuk_saati" validate:"required"`
	Yolculuk_süresi   string             `json:"yolculuk_suresi" validate:"required"`
	Gar               string             `json:"gar" validate:"required"`
	Bus_id            string             `json:"Bus_id"`
	Bos_koltuk_sayisi int
	Buy_Ticket
}

type Buy_Ticket struct {
	ID              string `json:"where_from" `
	Koltuk_numarasi int    `json:"koltuk_numarasi" `
	First_name      string `json:"first_name" `
	Last_name       string `json:"last_name" `
	Valid           bool   `json:"valid" `
	Price           int    `json:"price" `
}

type Cancel_Ticket struct {
	Bilet_İd   string
	First_name string
	Last_name  string
	Plaka_no   string
	Valid      bool
}
