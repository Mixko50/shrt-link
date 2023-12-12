package types

type CreateShortenLinkResponse struct {
	LongUrl string `json:"long_url"`
	Slug    string `json:"slug"`
}
