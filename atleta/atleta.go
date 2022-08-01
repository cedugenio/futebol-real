package atleta

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Atleta struct {
	AtletaID    int    `json:"atleta_id"`
	NomePopular string `json:"nome_popular"`
	Nome        string `json:"nome"`
	Posicao     struct {
		Nome  string `json:"nome"`
		Sigla string `json:"sigla"`
	} `json:"posicao"`
}

func GetInfoAtleta(id int) {
	url := fmt.Sprintf("https://api.api-futebol.com.br/v1/atletas/%d", id)
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

	var atletaStruct Atleta
	json.Unmarshal(bodyBytes, atletaStruct)
	fmt.Printf("Informação do Atleta: %v %v %v\n", atletaStruct.NomePopular, atletaStruct.Posicao.Nome, atletaStruct.Posicao.Sigla)

}
