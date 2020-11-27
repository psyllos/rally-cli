package model

type AllowedAttributeValue struct {
	PersistableObject

	IntegerValue         int64  `json:"integer_value"`
	LocalizedStringValue string `json:"localized_string_value"`
	StringValue          string `json:"string_value"`
	ValueIndex           int64  `json:"value_index"`
}
