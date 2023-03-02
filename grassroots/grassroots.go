package grassroots

import (
	"fmt"
	"sort"
	"time"
)

type Match struct {
	BatLimitID      int                 `json:"batLimitId"`
	DismissalModeID int                 `json:"dismissalModeId"`
	Grade           Grade               `json:"grade"`
	Innings         []Innings           `json:"innings"`
	IsBallByBall    bool                `json:"isBallByBall"`
	LegacyMatchID   int                 `json:"legacyMatchId"`
	MatchSchedule   []MatchScheduleItem `json:"matchSchedule"`
	MatchSummary    MatchSummary        `json:"matchSummary"`
	MatchType       string              `json:"matchType"`
	MatchTypeID     int                 `json:"matchTypeId"`
	Officials       []interface{}       `json:"officials"`
	Round           Round               `json:"round"`
	Status          string              `json:"status"`
	StatusID        int                 `json:"statusId"`
	Teams           []Team              `json:"teams"`
	Venue           Venue               `json:"venue"`
}

type Bowler struct {
	LegacyPlayerID  int     `json:"legacyPlayerId"`
	PlayerShortName string  `json:"playerShortName"`
	OversBowled     float64 `json:"oversBowled"`
	MaidensBowled   int     `json:"maidensBowled"`
	RunsConceded    int     `json:"runsConceded"`
	WicketsTaken    int     `json:"wicketsTaken"`
	WideBalls       int     `json:"wideBalls"`
	NoBalls         int     `json:"noBalls"`
	Economy         string  `json:"economy"`
}

type Organisation struct {
	ID        string `json:"id"`
	LogoURL   string `json:"logoUrl"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
}

type Grade struct {
	LegacyGradeID int          `json:"legacyGradeId"`
	Name          string       `json:"name"`
	Organisation  Organisation `json:"organisation"`
}

type MatchScheduleItem struct {
	MatchDay      int       `json:"matchDay"`
	StartDateTime time.Time `json:"startDateTime"`
}

func ScheduleTime(schedule []MatchScheduleItem) string {
	sort.Slice(schedule, func(i, j int) bool {
		return schedule[i].MatchDay < schedule[j].MatchDay
	})
	var str string
	for i, day := range schedule {
		str += day.StartDateTime.Local().Format("Mon 02 Jan 2006 (3:04PM)")
		if i != len(schedule)-1 {
			str += ", "
		}
	}
	return str
}

type Batter struct {
	BallsFaced      int    `json:"ballsFaced"`
	BattingMinutes  int    `json:"battingMinutes"`
	DismissalText   string `json:"dismissalText"`
	DismissalType   string `json:"dismissalType"`
	DismissalTypeID int    `json:"dismissalTypeId"`
	FoursScored     int    `json:"foursScored"`
	LegacyPlayerID  int    `json:"legacyPlayerId"`
	PlayerShortName string `json:"playerShortName"`
	RunsScored      int    `json:"runsScored"`
	SixesScored     int    `json:"sixesScored"`
	StrikeRate      string `json:"strikeRate,omitempty"`
}

type FallOfWicket struct {
	LegacyPlayerID  int    `json:"legacyPlayerId"`
	PlayerShortName string `json:"playerShortName"`
	Runs            int    `json:"runs"`
	Order           int    `json:"order"`
}

func FallOfWicketList(fow []FallOfWicket) string {
	sort.Slice(fow, func(i, j int) bool {
		if fow[i].Runs == fow[j].Runs {
			return fow[i].Order < fow[j].Order
		}
		return fow[i].Runs < fow[j].Runs
	})
	var str string
	for i, f := range fow {
		str += fmt.Sprintf("%d/%d (%s)", i+1, f.Runs, f.PlayerShortName)
		if i != len(fow)-1 {
			str += ", "
		}
	}
	return str
}

type Fielder struct {
	Catches             int    `json:"catches"`
	LegacyPlayerID      int    `json:"legacyPlayerId"`
	PlayerShortName     string `json:"playerShortName"`
	RunOuts             int    `json:"runOuts"`
	Stumpings           int    `json:"stumpings"`
	WicketKeeperCatches int    `json:"wicketKeeperCatches"`
}

type Innings struct {
	Batting               []Batter       `json:"batting"`
	Bowling               []Bowler       `json:"bowling"`
	ByesRuns              int            `json:"byesRuns"`
	FallOfWickets         []FallOfWicket `json:"fallOfWickets"`
	Fielding              []Fielder      `json:"fielding"`
	InningsNumber         int            `json:"inningsNumber"`
	InningsOrder          int            `json:"inningsOrder"`
	InningsStatus         int            `json:"inningsStatus"`
	IsDeclared            bool           `json:"isDeclared"`
	IsFollowOn            bool           `json:"isFollowOn"`
	LegByesRuns           int            `json:"legByesRuns"`
	LegacyBattingTeamID   string         `json:"legacyBattingTeamId"`
	LegacyInningsID       int            `json:"legacyInningsId"`
	Name                  string         `json:"name"`
	NoBalls               int            `json:"noBalls"`
	NumberOfWicketsFallen int            `json:"numberOfWicketsFallen"`
	OversBowled           float64        `json:"oversBowled"`
	Penalties             int            `json:"penalties"`
	RunsScored            int            `json:"runsScored"`
	TotalExtras           int            `json:"totalExtras"`
	WideBalls             int            `json:"wideBalls"`
}

type TeamSummary struct {
	BattedFirst  bool   `json:"battedFirst"`
	IsBatting    bool   `json:"isBatting"`
	IsHome       bool   `json:"isHome"`
	IsWinner     bool   `json:"isWinner"`
	LegacyTeamID string `json:"legacyTeamId"`
	ResultType   string `json:"resultType"`
	ResultTypeID int    `json:"resultTypeId"`
	ScoreText    string `json:"scoreText"`
	WonToss      bool   `json:"wonToss"`
}

type MatchSummary struct {
	ResultText string        `json:"resultText"`
	Teams      []TeamSummary `json:"teams"`
}

type Round struct {
	LegacyRoundID string `json:"legacyRoundId"`
	ShortName     string `json:"shortName"`
}

type Team struct {
	DisplayName        string       `json:"displayName"`
	LegacyTeamID       string       `json:"legacyTeamId"`
	OwningOrganisation Organisation `json:"owningOrganisation,omitempty"`
	Players            []Player     `json:"players"`
	NonPlayingMembers  []Player     `json:"nonPlayingMembers"`
}

type Player struct {
	LegacyPlayerID int           `json:"legacyPlayerId"`
	Name           string        `json:"name"`
	Roles          []interface{} `json:"roles"`
	ShortName      string        `json:"shortName"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Venue struct {
	Name           string   `json:"name"`
	Line1          string   `json:"line1"`
	PlayingSurface Location `json:"playingSurface"`
}
