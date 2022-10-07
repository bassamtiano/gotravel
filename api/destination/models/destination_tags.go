package destination_models

type DestinationTags struct {
	Id_denstination_tag int    `json:"id_destination_tag"`
	Id_destination      int    `json:"id_destination"`
	Id_tags_type        int    `json:"id_tags_type"`
	CreatedAt           string `json:"created_at"`
	UpdatedAt           string `json:"updated_at"`
}
