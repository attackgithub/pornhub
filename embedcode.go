package pornhub

type PornhubEmbedCode struct {
	Embed PornhubCode `json:"embed"`
}

type PornhubCode struct {
	Code string `json:"code"`
}