package Store

type Config struct {
	Dbhost       string `yaml:"dbhost"`
	Dbname       string `yaml:"dbname"`
	Dbusername   string `yaml:"dbusername"`
	Dbpassword   string `yaml:"dbpassword"`
	Dockerdbport string `yaml:"dockerdbport"`
}

type Ad struct {
	Id    int64   `json:"id"`
	Title string  `json:"title"`
	Photo string  `json:"photo"`
	Price float64 `json:"price"`
}
