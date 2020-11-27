package model

type AttributeDefinition struct {
	PersistableObject

	AllowedQueryOperators []AllowedQueryOperator `json:"allowed_query_operators"`
	AllowedValues         []AllowedValue         `json:"allowed_values"`
	AllowedValueTypeUUID  string                 `json:"allowed_value_type_uuid"`
	AttributeType         string                 `json:"attribute_type"`
	Constrained           bool                   `json:"constrained"`
	Custom                bool                   `json:"custom"`
	DetailedType          string                 `json:"detailed_type"`
	ElementName           string                 `json:"element_name"`
	Filterable            bool                   `json:"filterable"`
	Hidden                bool                   `json:"hidden"`
	Hideable              bool                   `json:"hideable"`
	MaxFractionalDigits   int64                  `json:"max_fractional_digits"`
	MaxLength             int64                  `json:"max_length"`
	Name                  string                 `json:"name"`
	Note                  string                 `json:"note"`
	Owned                 bool                   `json:"owned"`
	ReadOnly              bool                   `json:"read_only"`
	RealAttributeType     string                 `json:"real_attribute_type"`
	Required              bool                   `json:"required"`
	Sortable              bool                   `json:"sortable"`
	SystemRequired        bool                   `json:"system_required"`
	Type                  string                 `json:"type"`
	VisibleOnlyToAdmins   bool                   `json:"visible_only_to_admins"`
}
