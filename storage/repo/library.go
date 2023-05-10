package repo

import "context"
type LibraryStorageI interface{
	SignUpUser(ctx context.Context)(string, error)
}