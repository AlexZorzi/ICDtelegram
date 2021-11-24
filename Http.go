package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func SearchICDLinEntity(query string) ([]LineEntity, error) {
	latestRelease := "2021-05"
	lineentityList := struct {
		DestinationEntities []LineEntity
	}{}
	bodyBytes, err := ApiRequest("https://icd.api.watn3y.net/icd/release/11/"+latestRelease+"/mms/search?q="+query)
	if err := json.Unmarshal(bodyBytes, &lineentityList); err != nil {
		PrintErr(err)
	}
	return lineentityList.DestinationEntities, err
}

func GetICDLinEntityByID(ID string) (LineEntity, error) {
	lineEntity := LineEntity{}
	bodyBytes, err := ApiRequest("https://icd.api.watn3y.net/icd/release/11/2021-05/mms/"+ID)
	if err := json.Unmarshal(bodyBytes, &lineEntity); err != nil {
		PrintErr(err)
	}
	return lineEntity, err
}


func GetICDLinEntityByCode(code string) (LineEntity, error) {
	latestRelease := "2021-05"
	lineEntity := LineEntity{}
	bodyBytes, err := ApiRequest("https://icd.api.watn3y.net/icd/release/11/"+latestRelease+"/mms/codeinfo/"+code)
	if err := json.Unmarshal(bodyBytes, &lineEntity); err != nil {
		PrintErr(err)
	}
	return lineEntity, err
}

func GetICDFoundationByID(ID string) (Entity, error) {
	entity := Entity{}
	bodyBytes, err := ApiRequest("https://icd.api.watn3y.net/icd/entity/"+ID)
	if err := json.Unmarshal(bodyBytes, &entity); err != nil {
		PrintErr(err)
	}
	return entity, err
}

func GetICD10ByCode(code string) (Entity10, error) {
	entity10 := Entity10{}
	bodyBytes, err := ApiRequest("http://icd10api.com/?code="+code+"&desc=long&r=json")
	if err := json.Unmarshal(bodyBytes, &entity10); err != nil {
		PrintErr(err)
	}
	return entity10, err
}

func ApiRequest(apiUrl string) ([]byte, error) {
	apiUrl = strings.ReplaceAll(apiUrl, " ", "%20") // could be better but fixes multi words queries 'Autism spectrum disorder' -> 'Autism%20spectrum%20disorder'
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		PrintErr(err)
	}
	client := &http.Client{}
	///////////////////////////////////
	req.Header.Set("Accept-Language", "en")
	req.Header.Set("API-Version", "v2")
	req.Header.Set("accept","application/json")
	// Headers not needed for ICD-10 Api but whatever
	///////////////////////////////////
	resp, err := client.Do(req)
	if err != nil {
		PrintErr(err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	return bodyBytes, err
}


