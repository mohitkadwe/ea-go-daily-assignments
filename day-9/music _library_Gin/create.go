package songs

type songResponse struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Singer        string `json:"singer"`
	Writer        string `json:"writer"`
	Released_date string `json:"released_date"`
}

type songRequest struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Singer        string `json:"singer"`
	Writer        string `json:"writer"`
	Released_date string `json:"released_date"`
}

var songs = []songResponse{
	{
		ID:            1,
		Title:         "Song 1",
		Singer:        "Singer 1",
		Writer:        "Writer 1",
		Released_date: "2021-01-01",
	},
	{
		ID:            2,
		Title:         "Song 2",
		Singer:        "Singer 2",
		Writer:        "Writer 2",
		Released_date: "2021-01-02",
	},
}
