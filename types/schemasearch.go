package types

type SchemaSearch struct {
	UserID     *float64 `json:"user_id"`
	SchemaID   *float64 `json:"schema_id"`
	SchemaName *string  `json:"schema_name"`
	UserName   *string  `json:"user_name"`
	Keywords   *string  `json:"keywords"`
}

type SchemaSearchResult struct {
	Schema
	User *User       `json:"user"`
	Mods []string    `json:"mods"`
	Tags []SchemaTag `json:"tags"`
}
