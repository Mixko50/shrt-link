package types

type CreateShortenLinkRequest struct {
	LongURL string  `json:"long_url" validate:"required,url"`
	Slug    *string `json:"slug" validate:"omitempty,alphanum"`
}
