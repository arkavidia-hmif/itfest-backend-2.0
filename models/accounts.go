package models

import "itfest-backend-2.0/types"

type Accounts struct {
	id       int64
	username string
	password types.EncryptedString
	role     types.Role
}
