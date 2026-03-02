package site

type SiteData struct {
	ErrorMsg        any    `json:"errorMsg,omitempty"`
	ErrorType       string `json:"errorType"`
	IsNSFW          bool   `json:"isNSFW,omitempty"`
	URL             string `json:"url"`
	URLMain         string `json:"urlMain"`
	UsernameClaimed string `json:"username_claimed"`
}

const Data = "https://raw.githubusercontent.com/sherlock-project/sherlock/master/sherlock_project/resources/data.json"
