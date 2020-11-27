package model

type TypeDefinition struct {
	PersistableObject

	Abstract    bool                  `json:"abstract"`
	Attributes  []AttributeDefinition `json:"attributes"`
	Copyable    bool                  `json:"copyable"`
	Creatable   bool                  `json:"creatable"`
	Deletable   bool                  `json:"deletable"`
	DisplayName string                `json:"display_name"`
	ElementName string                `json:"element_name"`
	IDPrefix    string                `json:"id_prefix"`
	Name        string                `json:"name"`
	Note        string                `json:"note"`
	Ordinal     int64                 `json:"ordinal"`
	ParentUUID  string                `json:"parent_uuid"`
	Queryable   bool                  `json:"queryable"`
	ReadOnly    bool                  `json:"read_only"`
	Restorable  bool                  `json:"restorable"`
	TypePath    string                `json:"type_path"`
}
