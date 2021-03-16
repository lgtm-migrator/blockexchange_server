package web

import (
	"blockexchange/types"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (api *Api) CreateSchemaPart(w http.ResponseWriter, r *http.Request, ctx *SecureContext) {
	schemapart := types.SchemaPart{}
	err := json.NewDecoder(r.Body).Decode(&schemapart)
	if err != nil {
		SendError(w, 500, err.Error())
		return
	}

	schema, err := api.SchemaRepo.GetSchemaById(schemapart.SchemaID)
	if err != nil {
		SendError(w, 500, err.Error())
		return
	}

	if schema.UserID != ctx.Token.UserID {
		SendError(w, 403, "you are not the owner of the schema")
		return
	}

	err = api.SchemaPartRepo.CreateOrUpdateSchemaPart(&schemapart)
	if err != nil {
		SendError(w, 500, err.Error())
		return
	}

	SendJson(w, schemapart)
}

func extractSchemaPartVars(r *http.Request) (int64, int, int, int, error) {
	vars := mux.Vars(r)
	schema_id, err := strconv.Atoi(vars["schema_id"])
	if err != nil {
		return 0, 0, 0, 0, err
	}
	x, err := strconv.Atoi(vars["x"])
	if err != nil {
		return 0, 0, 0, 0, err
	}

	y, err := strconv.Atoi(vars["y"])
	if err != nil {
		return 0, 0, 0, 0, err
	}

	z, err := strconv.Atoi(vars["z"])
	if err != nil {
		return 0, 0, 0, 0, err
	}

	return int64(schema_id), x, y, z, nil
}

func (api *Api) GetSchemaPart(w http.ResponseWriter, r *http.Request) {
	schema_id, x, y, z, err := extractSchemaPartVars(r)
	if err != nil {
		SendError(w, 500, err.Error())
		return
	}

	schemapart, err := api.SchemaPartRepo.GetBySchemaIDAndOffset(int64(schema_id), x, y, z)
	if err != nil {
		SendError(w, 500, err.Error())
		return
	}

	SendJson(w, schemapart)
}

func (api *Api) GetNextSchemaPart(w http.ResponseWriter, r *http.Request) {
	schema_id, x, y, z, err := extractSchemaPartVars(r)
	if err != nil {
		SendError(w, 500, err.Error())
		return
	}

	schemapart, err := api.SchemaPartRepo.GetNextBySchemaIDAndOffset(int64(schema_id), x, y, z)
	if err != nil {
		SendError(w, 500, err.Error())
		return
	}

	SendJson(w, schemapart)
}