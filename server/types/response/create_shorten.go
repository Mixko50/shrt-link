package response

type CreateShortenLinkResponse struct {
	OriginalUrl string `json:"orginal_url"`
	Slug        string `json:"slug"`
}
