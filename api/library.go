package api

import (
	"Library/entity"
	"net/http"
)

func(api *api)SignUpUser(w http.ResponseWriter, r *http.Request){
	var body entity.User

	if err:=BodyParser(r, &body); err!=nil {
		HandleBadRequestErrWithMessage(w, err, "error in parsing json")
		return
	}

	
	
}