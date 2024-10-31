package eventstream

type KafkaData struct {
	UserID      string      `json:"user_id"`
	MovieID     string      `json:"movie_id"`
	Preferences Preferences `json:"preferences"`
	Next        bool        `json:"next"`
}

type Preferences struct {
	GenreScore       []Score `json:"genre_score"`
	ProtagonistScore []Score `json:"protagonist_score"`
	DirectorScore    []Score `json:"director_score"`
}

type Score struct {
	Name  string  `json:"name"`
	Score float64 `json:"score"`
}
