package clube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Clube struct {
	TimeID      int    `json:"time_id"`
	NomePopular string `json:"nome_popular"`
	Sigla       string `json:"sigla"`
	Nome        string `json:"nome"`
	Apelido     string `json:"apelido"`
}

func GetInfoClube(id int) {
	url := fmt.Sprintf("https://api.api-futebol.com.br/v1/times/%d", id)
	var bearer = "Bearer " + ""
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", bearer)
	if err != nil {
		log.Println("Ocorreu um erro ", err)
	}
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println("Ocorreu um erro ", err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var clubeStruct Clube
	json.Unmarshal(bodyBytes, &clubeStruct)
	fmt.Printf("O ID do time %v tem o nome conhecido pelo torcedor %v\n", clubeStruct.TimeID, clubeStruct.NomePopular)

}
