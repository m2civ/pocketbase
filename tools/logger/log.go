package logger

import (
	"log/slog"
	"time"

	"github.com/m2civ/pocketbase/tools/types"
)

// Log is similar to [slog.Record] bit contains the log attributes as
// preformatted JSON map.
type Log struct {
	Time    time.Time
	Data    types.JSONMap[any]
	Message string
	Level   slog.Level
}
