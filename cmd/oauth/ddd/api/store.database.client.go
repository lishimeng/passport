package openapi

import (
	"context"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/lishimeng/app-starter"
)

type DbClientStore struct {
}

func NewClientStore() (cs oauth2.ClientStore) {
	dcs := DbClientStore{}
	cs = &dcs
	return
}

func (cs *DbClientStore) GetByID(_ context.Context, id string) (ci oauth2.ClientInfo, err error) {
	ctx := app.GetOrm()
	ci, err = GetClientByAppKey(*ctx, id)
	return
}
