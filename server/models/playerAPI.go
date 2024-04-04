package models

// PlayerAPI représente la structure des données de l'API du joueur
type PlayerAPI struct {
	ID            string    `json:"_id"`
	Username      string    `json:"username"`
	Role          string    `json:"role"`
	Ts            *string   `json:"ts,omitempty"`
	Botmaster     *string   `json:"botmaster,omitempty"`
	Badges        []Badge   `json:"badges"`
	XP            float64   `json:"xp"`
	GamesPlayed   int       `json:"gamesplayed"`
	GamesWon      int       `json:"gameswon"`
	GameTime      float64   `json:"gametime"`
	Country       *string   `json:"country,omitempty"`
	Badstanding   *bool     `json:"badstanding,omitempty"`
	Supporter     bool      `json:"supporter"`
	SupporterTier int       `json:"supporter_tier"`
	Verified      bool      `json:"verified"`
	League        *League   `json:"league,omitempty"`
	Avatar        *Avatar   `json:"avatar,omitempty"`
	Banner        *Banner   `json:"banner,omitempty"`
	Bio           *string   `json:"bio,omitempty"`
	Connections   *struct{} `json:"connections,omitempty"`
	FriendCount   int       `json:"friend_count"`
	Distinguishment *struct {
		Type string `json:"type"`
	} `json:"distinguishment,omitempty"`
}

// Badge représente un badge de l'utilisateur
type Badge struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Ts    *string `json:"ts,omitempty"`
}

// League représente le standing de l'utilisateur dans la ligue Tetra
type League struct {
	GamesPlayed  int     `json:"gamesplayed"`
	GamesWon     int     `json:"gameswon"`
	Rating       float64 `json:"rating"`
	Rank         string  `json:"rank"`
	BestRank     string  `json:"bestrank"`
	Standing     int     `json:"standing"`
	StandingLocal int    `json:"standing_local"`
	NextRank     *string `json:"next_rank,omitempty"`
	PrevRank     *string `json:"prev_rank,omitempty"`
	NextAt       int     `json:"next_at"`
	PrevAt       int     `json:"prev_at"`
	Percentile   float64 `json:"percentile"`
	PercentileRank string `json:"percentile_rank"`
	Glicko        *float64 `json:"glicko,omitempty"`
	RD            *float64 `json:"rd,omitempty"`
	Apm           *float64 `json:"apm,omitempty"`
	Pps           *float64 `json:"pps,omitempty"`
	Vs            *float64 `json:"vs,omitempty"`
	Decaying      bool     `json:"decaying"`
}

// Avatar représente l'avatar de l'utilisateur
type Avatar struct {
	Revision int `json:"revision"`
}

// Banner représente la bannière de l'utilisateur
type Banner struct {
	Revision int `json:"revision"`
}

type PlayerInfoResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
	Data    struct {
		Player PlayerAPI `json:"user"`
	} `json:"data,omitempty"`
}