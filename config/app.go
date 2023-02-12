package config

import "goex/pkg/config"

func init() {
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{}{
			// Application name
			"name": config.Env("APP_NAME", "Goex"),

			// Current environment, used to distinguish the environment, usually local, stage, production, test
			"env": config.Env("APP_ENV", "production"),

			// Whether to enter debug mode
			"debug": config.Env("APP_DEBUG", false),

			// Application service port
			"port": config.Env("APP_PORT", "3000"),

			// Encrypted sessions, JWT encryption
			"key": config.Env("APP_KEY", ""),

			// To generate links
			"url": config.Env("APP_URL", "http://localhost:3000"),

			// Set time zone
			"timezone": config.Env("TIMEZONE", "Asia/Shanghai"),
		}
	})
}
