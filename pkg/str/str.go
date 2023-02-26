package str

import (
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

// Plural to pluralize user -> users
func Plural(word string) string {
	return pluralize.NewClient().Plural(word)
}

// Singular to singular users -> user
func Singular(word string) string {
	return pluralize.NewClient().Singular(word)
}

// Snake to snake_case, such as TopicComment -> topic_comment
func Snake(s string) string {
	return strcase.ToSnake(s)
}

// Camel to CamelCase, such as topic_comment -> TopicComment
func Camel(s string) string {
	return strcase.ToCamel(s)
}

// LowerCamel to lowerCamelCase, such as TopicComment -> topicComment
func LowerCamel(s string) string {
	return strcase.ToLowerCamel(s)
}
