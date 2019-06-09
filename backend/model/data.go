package model

type News struct {
	Multi    string `json:"multi"`
	Text     string `json:"text"`
	User     User   `json:"user"`
	HasUrl   bool   `json:"has_url"`
	Comments int    `json:"comments"`
	Pics     int    `json:"pics"`
	Source   string `json:"source"`
	Likes    int    `json:"likes"`
	Time     int    `json:"time"`
	Reposts  int    `json:"reposts"`
}
type User struct {
	Verified     bool   `json:"verified"`
	Description  bool   `json:"description"`
	Gender       string `json:"gender"`
	Messages     int    `json:"messages"`
	Followers    int    `json:"followers"`
	Location     string `json:"location"`
	VerifiedType int    `json:"verified_type"`
}
type Comment struct {
	Kid    []string `json:"kids"`
	Uid    string   `json:"uid"`
	Parent string   `json:"parent"`
	Text   string   `json:"text"`
	Mid    string   `json:"mid"`
	Date   string   `json:"data"`
}
