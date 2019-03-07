package pornhub

type PornhubVideo struct {
	ID           string  `json:"video_id"`
	Duration     string  `json:"duration"`
	Views        float64 `json:"views"`
	Rating       string  `json:"rating"`
	Ratings      float64 `json:"ratings"`
	Title        string  `json:"title"`
	URL          string  `json:"url"`
	DefaultThumb string  `json:"default_thumb"`
	Thumb        string  `json:"thumb"`
	PublishDate  string  `json:"publish_date"`
	Thumbs       []PornhubThumb
	Tags         []PornhubTag
	Pornstars    []PornhubPornstar
	Categories   []PornhubCategory
	Segment      string `json:"segment"`
}

type PornhubThumb struct {
	Size   string `json:"size"`
	Width  string `json:"width"`
	Height string `json:"height"`
	Src    string `json:"src"`
}

type PornhubTag struct {
	Name string `json:"tag_name"`
}

type PornhubPornstar struct {
	Name string `json:"pornstar_name"`
}

type PornhubCategory struct {
	Name string `json:"category"`
}