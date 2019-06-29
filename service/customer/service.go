package customer

import "database/sql"

type CustomerService struct {
	Database *sql.DB
}