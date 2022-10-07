package destination_models

type Facility struct {
	Id_facility          int    `json:"id_facility"`
	Facility_name        string `json:"facility_name"`
	Facility_description string `json:"facility_description"`
	CreatedAt            string `json:"created_at"`
	UpdatedAt            string `json:"updated_at"`
}
