package dto

type UrlWithAliasDto struct {
	Url   string `json:"url" bson:"url"`
	Alias string `json:"alias" bson:"alias"`
}
