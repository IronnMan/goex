package config

import "goex/pkg/config"

func init() {
	config.Add("captcha", func() map[string]interface{} {
		return map[string]interface{}{
			"height": 80,

			"width": 240,

			"length": 6,

			"maxskew": 0.7,

			"dotcount": 80,

			"expire_time": 15,

			"debug_expire_time": 10080,

			"testing_key": "captcha_skip_test",
		}
	})
}
