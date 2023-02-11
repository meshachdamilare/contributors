package github

type Contributor struct {
	Login         string `json:"login"`
	Url           string `json:"url"`
	Contributions string `json:"contributions"`
}
