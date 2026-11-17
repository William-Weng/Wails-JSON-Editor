package main

type DroppedJSONFile struct {
	Name  string `json:"name"`
	Text  string `json:"text"`
	Error string `json:"error"`
}
