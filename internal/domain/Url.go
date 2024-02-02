package domain

type Url struct {
	Id    uint64 `db:"id"`
	Url   string `db:"url"`
	Alias string `db:"alias"`
}
