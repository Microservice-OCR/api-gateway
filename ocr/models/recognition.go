package models

type Recognition struct {
	Name string `json: "imagePath"`
	FullText string `json: "fulltext"`
}

func Tostring(recognition Recognition)[]string{
	var str []string
	str = append(str, recognition.Name)
	str = append(str, recognition.FullText)
	return str
}

type IComplexRecognition map[string]IComplexElement

type IComplexElement struct {
	Value  string `bson:"value"`
	Left   int    `bson:"left"`
	Top    int    `bson:"top"`
	Width  int    `bson:"width"`
	Height int    `bson:"height"`
}