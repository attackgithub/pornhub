package pornhub

type PornhubVideo struct {
	ID           string  `json:"video_id,omitempty"`
	Duration     string  `json:"duration,omitempty"`
	Views        float64 `json:"views,omitempty"`
	Rating       string  `json:"rating,omitempty"`
	Ratings      float64 `json:"ratings,omitempty"`
	Title        string  `json:"title,omitempty"`
	URL          string  `json:"url,omitempty"`
	DefaultThumb string  `json:"default_thumb,omitempty"`
	Thumb        string  `json:"thumb,omitempty"`
	PublishDate  string  `json:"publish_date,omitempty"`
	Thumbs       []PornhubThumb
	Tags         []PornhubTag
	Pornstars    []PornhubPornstar
	Categories   []PornhubCategory
	Segment      string `json:"segment,omitempty"`
}

type PornhubThumb struct {
	Size   string `json:"size,omitempty"`
	Width  string `json:"width,omitempty"`
	Height string `json:"height,omitempty"`
	Src    string `json:"src,omitempty"`
}

type PornhubTag struct {
	Name string `json:"tag_name,omitempty"`
}

type PornhubPornstar struct {
	Name string `json:"pornstar_name,omitempty"`
}

type PornhubCategory struct {
	Name string `json:"category,omitempty"`
}