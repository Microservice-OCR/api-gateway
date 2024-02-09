package models

type IResponseIMG struct {
	Message string `json: "message"`
	Id string `json: "ID"`
}

func Tostring(response IResponseIMG) []string {
	var str []string
	str = append(str, response.Message)
	str = append(str, response.Id)
	return str
}
