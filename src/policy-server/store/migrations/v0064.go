package migrations

var migration_v0064 = map[string][]string{
	"mysql": {
		`ALTER TABLE ip_ranges ADD COLUMN description text DEFAULT '';`,
	},
	"postgres": {
		`ALTER TABLE ip_ranges ADD COLUMN description text DEFAULT '';`,
	},
}

