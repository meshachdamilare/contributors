package github

type Contributor struct {
	Login         string `json:"login"`
	Url           string `json:"url"`
	Contributions int    `json:"contributions"`
}
