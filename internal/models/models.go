package models

import "time"

// Models are defined in repository package to keep GORM tags close to implementations.
// This file is left intentionally as a placeholder for pure-domain models if needed in future.

// Example domain models could live here with no ORM annotations.

type Dummy struct {
	CreatedAt time.Time
}
