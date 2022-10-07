package destination_models

type DestinationFacility struct {
	Id_destination int    `json:"id_destination"`
	Id_facility    int    `json:"id_facility"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
