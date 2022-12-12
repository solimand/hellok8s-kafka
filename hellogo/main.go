package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

type conf struct {
	Hi string `yaml:"greetings"`
}

func (c *conf) getConf() *conf {
	yamlFile, err := ioutil.ReadFile("/config/app.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

/*Handler Wrapper - reads ENV fromn OS env. variables*/
// func handler(w http.ResponseWriter, r *http.Request) {
// 	// greeting := os.Getenv(("GREETING"))
// 	greeting := os.Getenv(("greetings"))
// 	fmt.Fprintf(w, "%s, 世界\n", greeting)
// }

/*Handler Wrapper - for auto detect changes in ENV file*/
func wrapHandler(env *conf) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		env.getConf()
		greeting := env.Hi
		fmt.Fprintf(w, "%s, 世界\n", greeting)
	}
}

func main() {
	var c conf
	c.getConf()

	http.HandleFunc("/", wrapHandler(&c))
	fmt.Println("Running demo app. Press Ctrl+C to exit...")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
