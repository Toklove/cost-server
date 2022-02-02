package Config

type Yaml struct {
	Database
	Server
}
type Database struct {
	Hostname string `yaml:"hostname"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}
type Server struct {
	Port int `yaml:"port"`
}
