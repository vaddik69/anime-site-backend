package shikimori

type Date struct {
	Year  int    `json:"year"`
	Month int    `json:"month"`
	Day   int    `json:"day"`
	Date  string `json:"date"`
}

type Poster struct {
	ID          string `json:"id"`
	OriginalURL string `json:"originalUrl"`
	MainURL     string `json:"mainUrl"`
}

type Genre struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Russian string `json:"russian"`
	Kind    string `json:"kind"`
}

type Studio struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"imageUrl"`
}

type ExternalLink struct {
	ID        string `json:"id"`
	Kind      string `json:"kind"`
	URL       string `json:"url"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type PersonPoster struct {
	ID string `json:"id"`
}

type Person struct {
	ID     string       `json:"id"`
	Name   string       `json:"name"`
	Poster PersonPoster `json:"poster"`
}

type PersonRole struct {
	ID      string   `json:"id"`
	RolesRu []string `json:"rolesRu"`
	RolesEn []string `json:"rolesEn"`
	Person  Person   `json:"person"`
}

type CharacterPoster struct {
	ID string `json:"id"`
}

type Character struct {
	ID     string          `json:"id"`
	Name   string          `json:"name"`
	Poster CharacterPoster `json:"poster"`
}

type CharacterRole struct {
	ID        string    `json:"id"`
	RolesRu   []string  `json:"rolesRu"`
	RolesEn   []string  `json:"rolesEn"`
	Character Character `json:"character"`
}

type RelatedAnime struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RelatedManga struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Related struct {
	ID           string        `json:"id"`
	Anime        *RelatedAnime `json:"anime"`
	Manga        *RelatedManga `json:"manga"`
	RelationKind string        `json:"relationKind"`
	RelationText string        `json:"relationText"`
}

type Video struct {
	ID        string `json:"id"`
	URL       string `json:"url"`
	Name      string `json:"name"`
	Kind      string `json:"kind"`
	PlayerURL string `json:"playerUrl"`
	ImageURL  string `json:"imageUrl"`
}

type Screenshot struct {
	ID          string `json:"id"`
	OriginalURL string `json:"originalUrl"`
	X166URL     string `json:"x166Url"`
	X332URL     string `json:"x332Url"`
}

type ScoresStat struct {
	Score int `json:"score"`
	Count int `json:"count"`
}

type StatusesStat struct {
	Status string `json:"status"`
	Count  int    `json:"count"`
}

type Anime struct {
	ID                string          `json:"id"`
	MalID             string          `json:"malId"`
	Name              string          `json:"name"`
	Russian           string          `json:"russian"`
	LicenseNameRu     string          `json:"licenseNameRu,omitempty"`
	English           []string        `json:"english,omitempty"`
	Japanese          []string        `json:"japanese,omitempty"`
	Synonyms          []string        `json:"synonyms,omitempty"`
	Kind              string          `json:"kind"`
	Rating            string          `json:"rating"`
	Score             float64         `json:"score"`
	Status            string          `json:"status,omitempty"`
	Episodes          int             `json:"episodes,omitempty"`
	EpisodesAired     int             `json:"episodesAired,omitempty"`
	Duration          int             `json:"duration,omitempty"`
	AiredOn           *Date           `json:"airedOn,omitempty"`
	ReleasedOn        *Date           `json:"releasedOn,omitempty"`
	URL               string          `json:"url,omitempty"`
	Season            string          `json:"season,omitempty"`
	Poster            Poster          `json:"poster,omitempty"`
	Fansubbers        []string        `json:"fansubbers,omitempty"`
	Fandubbers        []string        `json:"fandubbers,omitempty"`
	Licensors         []string        `json:"licensors,omitempty"`
	CreatedAt         string          `json:"createdAt,omitempty"`
	UpdatedAt         string          `json:"updatedAt,omitempty"`
	NextEpisodeAt     string          `json:"nextEpisodeAt,omitempty"`
	IsCensored        bool            `json:"isCensored,omitempty"`
	Genres            []Genre         `json:"genres,omitempty"`
	Studios           []Studio        `json:"studios,omitempty"`
	ExternalLinks     []ExternalLink  `json:"externalLinks,omitempty"`
	PersonRoles       []PersonRole    `json:"personRoles,omitempty"`
	CharacterRoles    []CharacterRole `json:"characterRoles,omitempty"`
	Related           []Related       `json:"related,omitempty"`
	Videos            []Video         `json:"videos,omitempty"`
	Screenshots       []Screenshot    `json:"screenshots,omitempty"`
	ScoresStats       []ScoresStat    `json:"scoresStats,omitempty"`
	StatusesStats     []StatusesStat  `json:"statusesStats,omitempty"`
	Description       string          `json:"description,omitempty"`
	DescriptionHTML   string          `json:"descriptionHtml,omitempty"`
	DescriptionSource string          `json:"descriptionSource,omitempty"`
}

type AnimeSearchResponseData struct {
	Animes []Anime `json:"animes"`
}
