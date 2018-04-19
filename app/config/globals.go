package config

var basePath string

type GetVariable struct{}

func (v GetVariable) BasePathValue() string {
	basePath = "http://localhost:8024/"
	return basePath
}
