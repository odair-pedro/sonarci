package models

type ThreadModelWrapper struct {
	Value []ThreadModel `json:"value"`
}

type ThreadModel struct {
	Id         int                  `json:"id"`
	Status     string               `json:"status"`
	IsDeleted  bool                 `json:"isDeleted"`
	Comments   []ThreadCommentModel `json:"comments"`
	Properties ThreadPropertyModel  `json:"properties"`
}

type ThreadCommentModel struct {
	Id        int  `json:"id"`
	IsDeleted bool `json:"isDeleted"`
}

type ThreadPropertyModel struct {
	GeneratedBySonarCI ThreadPropertySonarCIModel `json:"generatedBySonarCI"`
	Tag                ThreadPropertySonarCIModel `json:"tag"`
}

type ThreadPropertySonarCIModel struct {
	Value string `json:"$value"`
}
