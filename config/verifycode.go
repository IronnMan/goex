package config

import "goex/pkg/config"

func init() {
	config.Add("verifycode", func() map[string]interface{} {
		return map[string]interface{}{
			"code_length": config.Env("VERIFY_CODE_LENGTH", 6),

			"expire_time": config.Env("VERIFY_CODE_EXPIRE", 15),

			// debug data
			"debug_expire_time":  10080,
			"debug_code":         123456,
			"debug_email_prefix": "@test.com",
		}
	})
}
