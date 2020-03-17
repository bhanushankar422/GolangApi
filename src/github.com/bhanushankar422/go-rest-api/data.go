package main

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Title 1",
		Description: "First element",
	},
	{
		ID:          "2",
		Title:       "Title 2",
		Description: "Second element",
	},
}
