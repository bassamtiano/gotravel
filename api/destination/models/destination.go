package destination_models

import "database/sql"

type Destination struct {
	Id_denstination               int     `json:"id_destinaion"`
	Destination_name              string  `json:"destination_name`
	Location                      string  `json:"location"`
	Coordinate_Lon                float64 `json:"coordinate_lo"`
	Coordinate_Lat                float64 `json:"coordinate_lat"`
	Description                   string  `json:"description"`
	Operation_start               string  `json:"operation_sart"`
	Operation_end                 string  `json:"operation_end"`
	Id_region                     int     `json:"id_region"`
	Id_destination_classification int     `json:"id_destinatin_classification"`
	CreatedAt                     string  `json:"created_at"`
	UpdatedAt                     string  `json:"updated_at"`
}

var currentDatetime string

func (dest *Destination) GetDestinationById(db *sql.DB) error {
	return db.QueryRow(
		"SELECT id_destination, destination_name, , id_entity_group WHERE id_entity=?",
		en.Id_entity).Scan(&en.Id_entity, &en.Title, &en.Step, &en.Id_entity_group)
}

func GetDestinationById(db *sql.DB, start, count int) ([]TagsType, error) {

	rows, err := db.Query(
		"SELECT * FROM destination LIMIT ? OFFSET ?",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	tags_types := []TagsType{}

	for rows.Next() {
		var tagt TagsType
		if err := rows.Scan(&tagt.Id_tag_type, &tagt.Tag_type); err != nil {
			return nil, err
		}
		tags_types = append(tags_types, tagt)
	}
	return tags_types, nil
}

func (dest *Destination) CreateDestination(db *sql.DB) error {
	currentDatetime = time.Now().Format("2006-01-02T15:04:05")
	_, err := db.Exec(
		"INSERT INTO destination(destination_name, location, coordinate_lon, coordinate_lat, description, operation_start, operation_end, id_region, id_destination_classification, created_at) VALUES(?, ?)",
		engr.Entity_group_name, currentDatetime)

	return err
}

func (dest *Destination) UpdateDestination(db *sql.DB) error {
	currentDatetime = time.Now().Format("2006-01-02T15:04:05")
}

func (dest *Destination) DeleteDestination(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM entity WHERE id_entity=?", en.Id_entity)
	return err
}