package types

type SchemaStar struct {
	UserID   int64 `json:"user_id" db:"user_id"`
	SchemaID int64 `json:"schema_id" db:"schema_id"`
}

type SchemaStarResponse struct {
	Count   int  `json:"count"`
	Starred bool `json:"starred"`
}
