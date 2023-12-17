package response

type CreateShortenLinkResponse struct {
	OriginalUrl string `json:"original_url"`
	Slug        string `json:"slug"`
}
