package tags_type

import (
	"database/sql"

	tagsTypeController "github.com/bassamtiano/gotravel/api/tags_type/controllers"
	"github.com/gorilla/mux"
)

func InitDB(Conn *sql.DB) {
	tagsTypeController.DBInit(Conn)
}

func AddTagsTypeRouterHandler(r *mux.Router) {
	url_prefix := "/tags_type/"

	r.HandleFunc(url_prefix+"get/all", tagsTypeController.GetTagsTypeHandler).Methods("GET")
	r.HandleFunc(url_prefix+"get/by_id", tagsTypeController.GetTagsTypeByIdHandler).Methods("GET")

	r.HandleFunc(url_prefix+"create", tagsTypeController.CreateTagsTypeHandler).Methods("POST")
	r.HandleFunc(url_prefix+"update", tagsTypeController.UpdateTagsTypeHandler).Methods("POST")
	r.HandleFunc(url_prefix+"delete", tagsTypeController.DeleteTagsTypeHandler).Methods("POST")
}
