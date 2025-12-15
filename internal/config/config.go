package config

import (
	"flag"
	"log"
	"os"
	"github.com/ilyakaznacheev/cleanenv"
)

//ques : ye sab bkcd hum kr kyu rhe hai?
//ans : usually hum jitne bhi configurations + secret data like api keys wagera jo hota hai usko hum .env mai delfine krte hai and then wha se use krte hai, but in go(abhi tak muje jitna pata chla hai) we have 2 ways ki hum .env as well as config files ka use bhi kar skta hai, this is the reason, hum iss project mai config files ka use kr rhe hai.

//steps of this : sabse phle humne, base directy mai config folder bankr ek local.ymal file banyi aur jitni cheeze hume chiye thi unko define kr diya, i think ye ek terha se hmamare liye iss project ko structure provide krta hai, ki kis type se ky ky cheeze lagegai for config
//then humne vo clean env go wali repo ko install kiya then uska use krke struct annotation create ki jo `...` aise likha hua hai
//finally hum ek mustload func likh rhe hai jo ki hamare server start hote time make sure krega ki saari config acche se load hui ya nh


type HttpServer struct {
	Addr string
}

// env-default:"production"
type Config struct { //ye struct hum, jo humne base direc mai config ke andar local.yaml mai likha tha uss basis par ban rhe hai
	Env         string `yaml:"env" env:"ENV" env-required:"true"` //struct tags
	StoragePath string `yaml:"storage_path" env-required:"true"` //these annotations ko hum vo "go clean env" wali repo se use kr paa rhe hai
	HttpServer  `yaml:"http_server"`
}



//about mustload func : sbse phle hum configPath get krke ki kosis kr rhe hai using os.Getenv method, agr isse nh mil rha then hum jab app hote time usually jo flags ke saath keys and paths daalte hai uske through parse krne ki kosish kr rhe hai, aur agr hume dono cases se bhi path nh mila tab hum fatal error degai since then aage nh badh skte

func MustLoad() *Config{
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "config file path")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("Config path is not set")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist : %s", configPath)
	}

	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("Can't read config file %s", err.Error())
	}
	
	return &cfg
}
