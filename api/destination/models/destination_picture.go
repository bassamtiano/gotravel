package destination_models

type DestinationPicture struct {
	Id_destination_picture string `json:"id_destination_picture"`
	File_name              string `json:"file_name"`
	Picture_title          string `json:"picture_title"`
	Picture_description    string `json:"picture_description"`
	CreatedAt              string `json:"created_at"`
	UpdatedAt              string `json:"updated_at"`
}
