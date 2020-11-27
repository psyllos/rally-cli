package model

type Workspace struct {
	PersistableObject

	Name            string            `json:"name"`
	Description     string            `json:"description"`
	TypeDefinitions []TypeDefinitions `json:"type_definitions"`
	SchemaVersion   string            `json:"schema_version"`
}
