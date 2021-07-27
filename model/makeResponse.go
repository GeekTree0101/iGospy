package model

// MakeResponse is generated spy object response
type MakeResponse struct {
	Interactor          string `json:"interactor"`
	InteractorInterface string `json:"interactorInterface"` // FIXME: interface & spy grouping :[
	Presenter           string `json:"presenter"`
	PresenterInterface  string `json:"presenterInterface"`
	Displayer           string `json:"displayer"`
	DisplayerInterface  string `json:"displayerInterface"`
}
