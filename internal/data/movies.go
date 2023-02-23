package data

import (
	"DotaReplays/internal/validator"
	"time"
)

// Annotate the Replay struct with struct tags to control how the keys appear in the
// JSON-encoded output.
type Replay struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Heroes    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}

func ValidateMovie(v *validator.Validator, replay *Replay) {
	v.Check(replay.Title != "", "title", "must be provided")
	v.Check(len(replay.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(replay.Year != 0, "year", "must be provided")
	v.Check(replay.Year >= 2011, "year", "must be greater than 2011")
	v.Check(replay.Year <= int32(time.Now().Year()), "year", "must not be in the future")
	v.Check(replay.Runtime != 0, "runtime", "must be provided")
	v.Check(replay.Runtime > 0, "runtime", "must be a positive integer")
	v.Check(replay.Heroes != nil, "heroes", "must be provided")
	v.Check(len(replay.Heroes) == 10, "heroes", "must contain 10 heroes")
	v.Check(validator.Unique(replay.Heroes), "heroes", "must not contain duplicate values")
}
