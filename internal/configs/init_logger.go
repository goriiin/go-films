package configs

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"time"
)

func InitLogger(lvl string) *slog.Logger {
	var logger *slog.Logger

	switch lvl {
	case "local":
		logger = newLogger(slog.LevelDebug)
	case "prod":
		logger = newLogger(slog.LevelInfo)
	}

	return logger
}

type DevJSONHandler struct {
	writer *os.File
	level  slog.Level
}

func (h *DevJSONHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return h
}

func (h *DevJSONHandler) WithGroup(_ string) slog.Handler {
	return h
}

func (h *DevJSONHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.level
}

func (h *DevJSONHandler) Handle(_ context.Context, r slog.Record) error {
	data := map[string]interface{}{
		"time":    r.Time.Format(time.RFC3339),
		"level":   r.Level.String(),
		"message": r.Message,
	}

	r.Attrs(func(a slog.Attr) bool {
		data[a.Key] = a.Value.Any()
		return true
	})

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	_, _ = fmt.Fprintf(h.writer, "%s", jsonData)

	return nil
}

func newLogger(level slog.Level) *slog.Logger {
	return slog.New(&DevJSONHandler{
		writer: os.Stdout,
		level:  level,
	})
}
