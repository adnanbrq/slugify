package dto

type CreateLinkDTO struct {
	Url  string `json:"url" valid:"string|min:5|max:64"`
	Slug string `json:"slug" valid:"string|min:1|max:64"`
}
