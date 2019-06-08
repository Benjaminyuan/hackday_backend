package model

type ProvinceAction struct {
	Code    string    `json:"code"`
	Message string    `json:"message"`
	Data    []HotData `json:"data"`
}
type ChinaAction struct {
	Code       string    `json:"code"`
	Message    string    `json:"message"`
	List       []HotData `json:"list"`
	TotalCount int     `json:"totalCount"`
	MaxPage    int       `json:"maxPage"`
}
type SelectChooseListAction struct {
	Code       string    `json:"code"`
	Message    string    `json:"message"`
	TotalCount int    `json:"totalCount"`
	List       []HotData `json:"list"`
	MaxPage    int       `json:"maxPage"`
}
type HotData struct {
	City           string     `json:"city"`
	CreateTime     string     `json:"createTime"`
	DifferenceDay  string     `json:"differenceDay"`
	Emotion        string     `json:"emotion"`
	IncidentSeq    int        `json:"incidentSeq"`
	IncidentTitle  string     `json:"incidentTitle"`
	KeyWord        string     `json:"keyword"`
	KeyWord_1      string     `json:"keyword1"`
	KeyWord_2      string     `json:"keyword2"`
	KeyWord_3      string     `json:"keyword3"`
	KeyWord_4      string     `json:"keyword4"`
	Labels         string     `json:"labels"`
	LineData       []LineData `json:"lineData"`
	LongTile       string     `json:"longTitle"`
	Origin         string     `json:"origin"`
	OriginalUrl    string     `json:"originalUrl"`
	Province       string     `json:"province"`
	Rank           int        `json:"rank"`
	RankDifference int        `json:"rankDifference"`
	RankLast       int        `json:"rankLast"`
}
type LineData struct {
	Total string `json:"total"`
	Name  string `json:"name"`
}

type TypeProp struct {
	Code    string         `json:"code"`
	Data    map[string]int `json:"data"`
	Message string         `json:"message"`
}
type RankData struct {
	Code    string `json:"code"`
	Data    Rank   `json:"data"`
	Message string `json:"message"`
}
type Rank struct {
	HotAsc           int `json:"hotAsc"`
	HotAvg           int `json:"hotAvg"`
	HotCount         int `json:"hotCount"`
	HotDesc          int `json:"hotDesc"`
	HotDifferenceAvg int `json:"hotDifferenceAvg"`
}
