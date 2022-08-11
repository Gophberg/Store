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

}
