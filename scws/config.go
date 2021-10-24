package scws

type ConfigurationDict struct {
	Type string
	File string
}
type Configuration struct {
	Dict   []ConfigurationDict
	Rule   string
	Listen string
}
