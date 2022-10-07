package tags_type_controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	tagsTypeModel "github.com/bassamtiano/gotravel/api/tags_type/models"
)

func GetTagsTypeHandler(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	fmt.Println(count)
	fmt.Println(start)

	if count > 10 || count < 1 {
		count = 10
	}

	if start < 0 {
		start = 0
	}

	tags_types, err := tagsTypeModel.GetTagsType(DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, tags_types)
}

func GetTagsTypeByIdHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))

	var tags_type tagsTypeModel.TagsType
	tags_type.Id_tag_type = id

	if err := tags_type.GetTagsTypeById(DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Product not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, tags_type)
}

func CreateTagsTypeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tags_type tagsTypeModel.TagsType

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tags_type); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := tags_type.CreateTagsType(DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, tags_type)
}

func UpdateTagsTypeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tags_type tagsTypeModel.TagsType

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tags_type); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := tags_type.UpdateTagsType(DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, tags_type)
}

func DeleteTagsTypeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tags_type tagsTypeModel.TagsType

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tags_type); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := tags_type.DeleteTagsType(DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, tags_type)
}
