package api

import (
	"encoding/json"
	"github.com/airbloc/airbloc-go/api"
	"github.com/airbloc/airbloc-go/schemas"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type API struct {
	schemas *schemas.Schemas
}

func New(backend *api.AirblocBackend) (api.API, error) {
	schemaManager := schemas.New(backend.MetaDatabase, backend.Ethclient)
	return &API{schemaManager}, nil
}

func (api *API) AttachToAPI(service *api.APIService) {
	RegisterSchemaServer(service.GrpcServer, api)
}

// TODO
func (api *API) Create(ctx context.Context, req *CreateSchemaRequest) (*CreateSchemaResult, error) {
	data := make(map[string]interface{})
	err := json.Unmarshal([]byte(req.Schema), &data)
	if err != nil {
		return nil, err
	}

	id, err := api.schemas.Register(req.Name, data)
	if err == schemas.ErrNameExists {
		return &CreateSchemaResult{Exists: true}, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to register schema")
	}

	return &CreateSchemaResult{
		Id:     id.String(),
		Exists: false,
	}, nil
}