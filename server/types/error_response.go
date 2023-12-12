package types

import "errors"

var (
	ErrCheckExistingUrl     = errors.New("error while checking existing url")
	ErrSlugAlreadyExists    = errors.New("slug already exists")
	ErrCannotCreateShrtLink = errors.New("unable to create shorten link")
	ErrSlugNotFound         = errors.New("slug not found")
	ErrSlugNotAlphanumeric  = errors.New("slug must be alphanumeric")
	ErrCannotUpdateVisit    = errors.New("unable to update visit")
)
