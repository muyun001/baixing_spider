package models

type HtmlExtractorResult struct {
	Title          string `json:"title"`
	ReleaseTime    string `json:"release_time"`
	CompanyName    string `json:"company_name"`
	ServiceContent string `json:"service_content"`
	ServiceRange   string `json:"service_range"`
	ContactPeople  string `json:"contact_people"`
	ContactPhone   string `json:"contact_phone"`
	ContactWeixin  string `json:"contact_weixin"`
	Poster         string `json:"poster"`
}

type UrlAndHtml struct {
	Url  string `json:"url"`
	Html string `json:"html"`
}

type SaveCsvResult struct {
	ExtractorResult HtmlExtractorResult `json:"extractor_result"`
	Url             string              `json:"url"`
	ChooseDate      string              `json:"choose_date"`
}
