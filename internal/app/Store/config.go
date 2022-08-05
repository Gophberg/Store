package Store

import (
	yaml "gopkg.in/yaml.v2"
	"io"
	"log"
	"os"
)

func (c *Config) NewConfig() {
	f, err := os.Open("configs/store.yaml")
	log.Println("[CONFIG]", f)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)
	fields, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(fields, &c)
	if err != nil {
		log.Fatal(err)
	}
}
