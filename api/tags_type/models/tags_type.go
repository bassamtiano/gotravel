package tags_type_models

import (
	"database/sql"
	"time"
)

type TagsType struct {
	Id_tag_type int    `json:"id_tag_type"`
	Tag_type    string `json:"tag_type"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func (tagt *TagsType) GetTagsTypeById(db *sql.DB) error {
	return db.QueryRow(
		"SELECT id_tag_type, tag_type FROM tags_type WHERE id_tag_type=?",
		tagt.Id_tag_type).Scan(&tagt.Id_tag_type, &tagt.Tag_type)
}

func GetTagsType(db *sql.DB, start, count int) ([]TagsType, error) {

	rows, err := db.Query(
		"SELECT id_tag_type, tag_type FROM tags_type LIMIT ? OFFSET ?",
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

func (tagt *TagsType) CreateTagsType(db *sql.DB) error {
	currentDatetime := time.Now().Format("2006-01-02T15:04:05")
	_, err := db.Exec(
		"INSERT INTO tags_type(tag_type, created_at) VALUES(?, ?)",
		tagt.Tag_type, currentDatetime)

	return err
}

func (tagt *TagsType) UpdateTagsType(db *sql.DB) error {
	currentDatetime := time.Now().Format("2006-01-02T15:04:05")
	_, err := db.Exec(
		"UPDATE tags_type SET tag_type=?, updated_at=? WHERE id_tag_type=?)",
		tagt.Tag_type, currentDatetime, tagt.Id_tag_type)
	return err
}

func (tagt *TagsType) DeleteTagsType(db *sql.DB) error {
	_, err := db.Exec(
		"DELETE FROM tag_type WHERE id_tag_type=?", tagt.Id_tag_type)
	return err
}
