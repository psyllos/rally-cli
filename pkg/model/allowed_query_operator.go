package model

type AllowedQueryOperator struct {
	PersistableObject

	OperatorName string `json:"operator_name"`
}
