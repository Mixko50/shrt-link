package request

type CreateShortenLinkRequest struct {
	OriginalUrl string  `json:"original_url" validate:"required,url"`
	Slug        *string `json:"slug" validate:"omitempty,alphanum,min=6"`
}
