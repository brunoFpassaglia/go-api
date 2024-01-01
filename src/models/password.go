package models

type Password struct {
	Old string `json:"old"`
	New string `json:"new"`
}
