package models

type IInput struct {
	Inputs []IInputComponent `json:"inputs"`
}

type IInputComponent struct {
	Name      string    `json:"name"`
	Rectangle Rectangle `json:"rectangle"`
}

type Rectangle struct {
	Left   int `json:"left"`
	Top    int `json:"top"`
	Width  int `json:"width"`
	Height int `json:"height"`
}