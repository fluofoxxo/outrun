package obj

type LineAuth struct {
	UserID            string `json:"userId"`
	Password          string `json:"password"`
	MigrationPassword string `json:"migrationPassword"`
}
