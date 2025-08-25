package sqlcqueries

import (
	"context"

	"github.com/alterejoe/generate/sqlc-go-helper/gov2/db"
	"github.com/jackc/pgx/v5/pgtype"
)

type SelectAllEmbroideryCustomers struct {
}

func (saec *SelectAllEmbroideryCustomers) Query(query *db.Queries, r context.Context) (any, error) {
	results, err := query.SelectAllEmbroideryCustomers(r)
	return results, err
}

type SelectAllScreenPrintCustomers struct {
}

func (saspc *SelectAllScreenPrintCustomers) Query(query *db.Queries, r context.Context) (any, error) {
	results, err := query.SelectAllScreenPrintCustomers(r)
	return results, err
}

type SelectAllTransferCustomers struct {
}

func (satc *SelectAllTransferCustomers) Query(query *db.Queries, r context.Context) (any, error) {
	results, err := query.SelectAllTransferCustomers(r)
	return results, err
}

type SelectCustomerByID struct {
	Id pgtype.UUID
}

func (scbid *SelectCustomerByID) Query(query *db.Queries, r context.Context) (any, error) {
	results, err := query.SelectCustomerByID(r, scbid.Id)
	return results, err
}

type SelectLatestCompletedCustomers struct {
}

func (slcc *SelectLatestCompletedCustomers) Query(query *db.Queries, r context.Context) (any, error) {
	results, err := query.SelectLatestCompletedCustomers(r)
	return results, err
}

type CompleteCustomer struct {
	Id pgtype.UUID
}

func (cc *CompleteCustomer) Query(query *db.Queries, r context.Context) (any, error) {
	err := query.CompleteCustomer(r, cc.Id)
	return nil, err
}

type InsertCustomer struct {
	Params db.InsertCustomerParams
}

func (ic *InsertCustomer) Query(query *db.Queries, r context.Context) (any, error) {
	results, err := query.InsertCustomer(r, ic.Params)
	return results, err
}

type RemoveCustomer struct {
	Id pgtype.UUID
}

func (rc *RemoveCustomer) Query(query *db.Queries, r context.Context) (any, error) {
	err := query.RemoveCustomer(r, rc.Id)
	return nil, err
}

type RestoreCustomer struct {
	Id pgtype.UUID
}

func (rc *RestoreCustomer) Query(query *db.Queries, r context.Context) (any, error) {
	err := query.RestoreCustomer(r, rc.Id)
	return nil, err
}

type UpdateCustomer struct {
	Params db.UpdateCustomerParams
}

func (uc *UpdateCustomer) Query(query *db.Queries, r context.Context) (any, error) {
	results, err := query.UpdateCustomer(r, uc.Params)
	return results, err
}

type Testing struct {
	Dollar_1 pgtype.Int4
}

func (t *Testing) Query(query *db.Queries, r context.Context) (any, error) {
	results, err := query.Testing(r, t.Dollar_1)
	return results, err
}
