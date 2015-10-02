package models

type OfferingList struct {
	Members []OfferingMember `json:"offerings"`
}

type OfferingMember struct {
	Key         string `json:"key"`
	Description string `json:"description"`
}
