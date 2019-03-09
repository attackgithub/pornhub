package pornhub

type PornhubEmbedCode struct {
	Embed PornhubCode `json:"embed,omitempty"`
}

type PornhubCode struct {
	Code string `json:"code,omitempty"`
}
