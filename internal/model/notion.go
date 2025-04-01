package model

type Parent struct {
	DatabaseId string `json:"database_id"`
}

type Text struct {
	Content string `json:"content"`
}

type Title struct {
	Text Text `json:"text"`
}

type CompanyName struct {
	Title []Title `json:"title"`
}

type Properties struct {
	CompanyName CompanyName `json:"company_name"`
}

type Page struct {
	Parent     Parent     `json:"parent"`
	Properties Properties `json:"properties"`
}
