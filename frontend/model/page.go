package model

type SearchResult struct {
	Hits     int64
	Start    int
	Profiles []interface{}
	Query    string
	PrevFrom int
	NextFrom int
}
