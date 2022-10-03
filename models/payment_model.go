package models

type Money struct {
	Amount_toload int
	Payment       `json:"payment" bson:"payment"`
	Save_card     bool
}

type Payment struct {
	Card_Number     string `json:"card_number" validate:"required,min=6"`
	Expiration_date string `json:"expiration_date" validate:"required,min=6"`
	Name_onthecard  string `json:"name_onthecard"`
	CSV             string `json:"csv"`
}
