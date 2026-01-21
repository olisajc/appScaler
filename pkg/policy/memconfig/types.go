package memconfig

type MemConfig struct {
	Request string `json:"request,omitempty"`
	Limit   string `json:"limit,omitempty"`
}
