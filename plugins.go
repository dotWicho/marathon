package marathon

// Plugins Array of
type Plugins struct {
	Plugins []Plugin `json:"plugins"`
}

// Plugin holds information of Marathon server plugins
type Plugin struct {
	ID             string     `json:"id"`
	Implementation string     `json:"implementation"`
	Info           PluginInfo `json:"info"`
	Plugin         string     `json:"plugin"`
	Tags           []string   `json:"tags"`
}

// PluginInfo version Info
type PluginInfo struct {
	Version string `json:"version"`
	Array   []int  `json:"array"`
	Test    bool   `json:"test"`
}
