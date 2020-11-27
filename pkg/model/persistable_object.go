package model

import "time"

type PersistableObject struct {
	CreationDate time.Time `json:"creation_date"`
	ObjectID     int64     `json:"object_id"`
	ObjectUUID   string    `json:"object_uuid"`
	VersionId    int       `json:"version_id"`
}
