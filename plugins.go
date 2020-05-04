package marathon

type Plugins struct {
	Plugins []Plugin `json:"plugins"`
}

type Plugin struct {
	ID             string   `json:"id"`
	Implementation string   `json:"implementation"`
	Info           Info     `json:"info"`
	Plugin         string   `json:"plugin"`
	Tags           []string `json:"tags"`
}

type Info struct {
	Version string `json:"version"`
	Array   []int  `json:"array"`
	Test    bool   `json:"test"`
}
