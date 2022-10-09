package models

type Personalidade struct {
	Nome     string `json:"nome"`
	Historia string `json:"historico"`
}

var Personalidades []Personalidade
