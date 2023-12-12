package utils

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func TimeToPgTime(t time.Time) pgtype.Timestamptz {
	return pgtype.Timestamptz{Time: t, InfinityModifier: pgtype.InfinityModifier(0), Valid: true}
}
