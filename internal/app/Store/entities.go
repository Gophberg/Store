package Store

type Config struct {
	Dbhost       string `yaml:"dbhost"`
	Dbname       string `yaml:"dbname"`
	Dbusername   string `yaml:"dbusername"`
	Dbpassword   string `yaml:"dbpassword"`
	Dockerdbport string `yaml:"dockerdbport"`
}

type Ad struct {
	Photo        string  `json:"photo"`
	Id           int64   `json:"id"`
	Title        string  `json:"title"`
	Content      string  `json:"content"`
	Price        float64 `json:"price"`
	CreationDate string  `json:"datecreated"`
}

type Result struct {
	Id     int64
	Status bool
	Reason string
}
