package Store

import (
	"github.com/joho/godotenv"
	"log"
)

func (c *Config) NewConfig() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	read, err := godotenv.Read()
	if err != nil {
		return
	}
	if err != nil {
		log.Println("[CONGIF]", err)
	}
	c.Dbhost = read["POSTGRES_HOST"]
	c.Dbname = read["POSTGRES_DB_NAME"]
	c.Dbusername = read["POSTGRES_USER"]
	c.Dbpassword = read["POSTGRES_PASSWORD"]
	c.Dockerdbport = read["DOCKER_DB_PORT"]

	//log.Println("[CONFIG] env:", os.Getenv("POSTGRES_HOST"))
	//log.Println("[CONFIG] env:", read)

	//f, err := os.Open("configs/store.yaml")
	//defer func(f *os.File) {
	//	err := f.Close()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}(f)
	//fields, err := io.ReadAll(f)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//err = yaml.Unmarshal(fields, &c)
	//if err != nil {
	//	log.Fatal(err)
	//}
}

//func NewConfig() (*Config, error) {
//	ctx := context.Background()
//	var c Config
//	if err := envconfig.Process(ctx, &c); err != nil {
//		return nil, err
//	}
//
//	return &c, nil
//}
