package models

type ThreadModelWrapper struct {
	Value []ThreadModel `json:"values"`
}

type ThreadModel struct {
	Id         string              `json:"id"`
	Status     string              `json:"status"`
	IsDeleted  bool                `json:"isDeleted"`
	Properties ThreadPropertyModel `json:"properties"`
}

type ThreadPropertyModel struct {
	GeneratedBySonarCI ThreadPropertySonarCIModel `json:"generatedBySonarCI"`
}

type ThreadPropertySonarCIModel struct {
	Value string `json:"$value"`
}
