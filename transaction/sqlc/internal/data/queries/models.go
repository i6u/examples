// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package queries

import ()

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type UserDetail struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}
