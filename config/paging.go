package config

import "goex/pkg/config"

func init() {
	config.Add("paging", func() map[string]interface{} {
		return map[string]interface{}{

			"perpage": 10,

			"url_query_page": "page",

			"url_query_sort": "sort",

			"url_query_order": "order",

			"url_query_per_page": "per_page",
		}
	})
}
