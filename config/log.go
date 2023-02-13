package config

import "goex/pkg/config"

func init() {
	config.Add("log", func() map[string]interface{} {
		return map[string]interface{}{
			// The log level must be one of the following options:
			// "debug" -- The amount of information is large, and it is generally opened during debugging. Detailed
			// operation logs of system modules, such as HTTP requests, database requests, sending emails, sending SMS
			// "info" -- Business-level operation logs, such as user login, user logout, and order cancellation.
			// "warn" -- Information that is of interest and needs attention. For example, print debugging information
			// when debugging (command line output will be highlighted).
			// "error" -- Log error messages. Panic or Error. Such as database connection error, HTTP port is occupied, etc.
			// General production environment.
			"level": config.Env("LOG_LEVEL", "debug"),

			// Type of log, optional:
			// "single" -- Separate file
			// "daily" -- One per day by date
			"type": config.Env("LOG_TYPE", "single"),

			/* ------------------ Rolling Log Configuration ------------------ */
			// Log file path
			"filename": config.Env("LOG_NAME", "storage/logs/logs.log"),
			// The maximum size of each log file saved Uint: M
			"max_size": config.Env("LOG_MAX_SIZE", 64),
			// The maximum number of log files to save, 0 is unlimited, and will still be deleted when MaxAge is reached
			"max_age": config.Env("LOG_MAX_AGE", 30),
			// Whether to compress
			"compress": config.Env("LOG_COMPRESS", false),
		}
	})
}
