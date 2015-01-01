package main

import ( 
	"io/ioutil"
	"encoding/json"
	"net/http"
	"log"
)

//Reads the configuration and builds a map out of it.
func configure() map[string]string{
	config := make(map[string]string)
	f, err := ioutil.ReadFile("./dns4e-auth.json")
	if err !=  nil {
		log.Fatal(err)
		panic(err)
	}
	jsonErr :=json.Unmarshal(f, &config)
	if jsonErr !=  nil {
		log.Fatal(jsonErr)
		panic(jsonErr)
	}
	config["url"] = "https://api.dns4e.com/v7/"+config["address"]+"/a"
	return config
}


func main(){
	var conf = make(map[string]string)
	conf = configure()
	
	req, err := http.NewRequest("POST",conf["url"],nil)
	req.SetBasicAuth(conf["pubkey"], conf["seckey"])	
	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err !=  nil {
		log.Fatal(err)
	}
	rbody, err := ioutil.ReadAll(resp.Body)		
	if err !=  nil {
		log.Fatal(err)
		log.Print(rbody)
	}
	resp.Body.Close()
}