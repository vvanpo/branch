package core

import (
	"github.com/vvanpo/branch/core/repo"
)

type Contact struct {
	repo *repo.Contacts
	id
	email string
}
