package client

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/viper"

	"main/pkg/config"
)

const (
	pathToConfig = "./configs/config.json"
)

var (
	cfg                 = getConfig(pathToConfig)
	authorizationHeader = getAuthorizationHeader(cfg)
	tr                  = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client              = &http.Client{Transport: tr}
)

func GetTargetTables() {
	url := fmt.Sprintf("%s/rest/v1/table/task", cfg.GetString("INSTANCE_URL"))
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Panic(err.Error())
	}

	req.Header.Add("authorization", authorizationHeader)
	fmt.Println(req.Header)
	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err.Error())
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	// t = CreateTable()
}

func getConfig(p string) *viper.Viper {
	cfg, err := config.LoadConfig(p)
	if err != nil {
		log.Panic(err.Error())
	}
	return cfg
}

func getAuthorizationHeader(c *viper.Viper) string {
	authorizationHeader := fmt.Sprintf("Bearer %s", c.GetString("TOKEN"))
	return authorizationHeader
}
