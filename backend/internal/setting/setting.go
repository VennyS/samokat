package setting

import (
	"samokat/internal/lib/env"
	cfg "samokat/internal/setting/config"
	"time"

	"go.uber.org/zap"
)

type App struct {
	cfg    cfg.Config
	logger *zap.SugaredLogger
}

func (a *App) Loadcfg() {
	a.cfg = cfg.Config{
		AppEnv: env.GetString("APP_ENV", "development"),
		Server: cfg.ServerConfig{
			Addr: env.GetString("ADDR", ":8080"),
		},
		HTTP: cfg.HTTPConfig{
			AccessSecret:  env.GetString("HTTP_ACCESS_SECRET", "randomaccesssecret"),
			RefreshSecret: env.GetString("HTTP_REFRESH_SECRET", "randomrefreshsecret"),
			AccessTTL:     env.GetDuration("HTTP_ACCESS_TTL", 1*time.Hour),
			RefreshTTL:    env.GetDuration("HTTP_REFRESH_TTL", 30*24*time.Hour),
		},
		DB: cfg.DBConfig{
			Driver:   env.GetString("DB_DRIVER", "postgres"),
			Host:     env.GetString("DB_HOST", "localhost"),
			Port:     env.GetString("DB_PORT", "5432"),
			User:     env.GetString("DB_USER", "postgres"),
			Password: env.GetString("DB_PASSWORD", "postgres"),
			Name:     env.GetString("DB_NAME", "crocus"),
			SSLMode:  env.GetString("DB_SSLMODE", "disable"),
		},
	}
}
