package main

type LineEntity struct {
	UrlID string `json:"stemId"`
	TheCode string `json:"theCode,omitempty"`
	Code string `json:"Code,omitempty"`
	BrowserUrl string `json:"browserUrl"`
	CodeRange string `json:"codeRange"`
}
func (lentity LineEntity) ID() string {
	return GetIDfromUrl(lentity.UrlID)
}
func (lentity LineEntity) GetCode() string {
	if lentity.TheCode == "" && lentity.Code == ""{
		return lentity.CodeRange
	}else if lentity.CodeRange == "" && lentity.Code == ""{
		return lentity.TheCode
	}else {
		return lentity.Code
	}
}