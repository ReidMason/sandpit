package anilist

type AnimeResponse struct {
	Data struct {
		Media AnimeResult `json:"Media"`
	} `json:"data"`
}

type AnimeResult struct {
	ID       int32    `json:"id"`
	Format   string   `json:"format"`
	Episodes int      `json:"episodes"`
	Synonyms []string `json:"synonyms"`
	Status   string   `json:"status"`
	EndDate  struct {
		Year  int `json:"year"`
		Month int `json:"month"`
		Day   int `json:"day"`
	} `json:"endDate"`
	StartDate struct {
		Year  int `json:"year"`
		Month int `json:"month"`
		Day   int `json:"day"`
	} `json:"startDate"`
	Title struct {
		English string `json:"english"`
		Romaji  string `json:"romaji"`
	} `json:"title"`
	Relations struct {
		Edges []struct {
			RelationType string `json:"relationType"`
		} `json:"edges"`
		Nodes []struct {
			ID      int32  `json:"id"`
			Format  string `json:"format"`
			EndDate struct {
				Year  any `json:"year"`
				Month any `json:"month"`
				Day   any `json:"day"`
			} `json:"endDate"`
			StartDate struct {
				Year  int `json:"year"`
				Month int `json:"month"`
				Day   int `json:"day"`
			} `json:"startDate"`
		} `json:"nodes"`
	} `json:"relations"`
}

type Anime struct {
	ID       int32    `json:"id"`
	Format   string   `json:"format"`
	Episodes int      `json:"episodes"`
	Synonyms []string `json:"synonyms"`
	Status   string   `json:"status"`
	EndDate  struct {
		Year  int `json:"year"`
		Month int `json:"month"`
		Day   int `json:"day"`
	} `json:"endDate"`
	StartDate struct {
		Year  int `json:"year"`
		Month int `json:"month"`
		Day   int `json:"day"`
	} `json:"startDate"`
	Title struct {
		English string `json:"english"`
		Romaji  string `json:"romaji"`
	} `json:"title"`
	Relations []Relation `json:"relations"`
}

type Relation struct {
	ID       int32  `json:"id"`
	Format   string `json:"format"`
	Relation string `json:"relation"`
	EndDate  struct {
		Year  any `json:"year"`
		Month any `json:"month"`
		Day   any `json:"day"`
	} `json:"endDate"`
	StartDate struct {
		Year  int `json:"year"`
		Month int `json:"month"`
		Day   int `json:"day"`
	} `json:"startDate"`
}

const (
	Adaptation  = "ADAPTATION"
	Prequel     = "PREQUEL"
	Sequel      = "SEQUEL"
	Parent      = "PARENT"
	SideStory   = "SIDE_STORY"
	Character   = "CHARACTER"
	Summary     = "SUMMARY"
	Alternative = "ALTERNATIVE"
	SpinOff     = "SPIN_OFF"
	Other       = "OTHER"
	Source      = "SOURCE"
	Compilation = "COMPILATION"
	Contains    = "CONTAINS"
)
