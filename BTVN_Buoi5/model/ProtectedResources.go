package model

type ProtectedResources struct {
	Users      []string `json:"users"`
	Products   []string `json:"products"`
	Categories []string `json:"categories"`
}
