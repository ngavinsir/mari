package graph

import "database/sql"

// Resolver { db }
type Resolver struct{
	DB	*sql.DB
}
