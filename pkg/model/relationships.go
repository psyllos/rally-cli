package model

type Entity struct {
	Ref    string `json:"_ref"`
	Loaded bool
}

type OneToOne struct {
	Entity

	RefObjectUUID string `json:"_refObjectUUID"`
	RefObjectName string `json:"_refObjectName"`
	Type          string `json:"_type"`
}

type OneToMany struct {
	Entity

	Ref   string `json:"_ref"`
	Type  string `json:"_type"`
	Count int    `json:"Count"`
}
