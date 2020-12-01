package model

type Workspace struct {
	PersistableObject

	Name            string           `json:"name"`
	Description     string           `json:"description"`
	SchemaVersion   string           `json:"schema_version"`
	TypeDefinitions []TypeDefinition `json:"type_definitions"`
}
