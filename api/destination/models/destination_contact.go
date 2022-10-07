package destination_models

type DestinationContact struct {
	Id_destination_contact int    `json:"id_destination_contact"`
	Phone                  int    `json:"phone"`
	Twitter                string `json:"twitter"`
	Facebook               string `json:"facebook"`
	Instagram              string `json:"instagram"`
	Whatsapp               string `json:"whatsapp"`
	Email                  string `json:"email"`
	CreatedAt              string `json:"created_at"`
	UpdatedAt              string `json:"updated_at"`
}
