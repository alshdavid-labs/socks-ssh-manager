package appstate

type State struct {
	configPath      string
	Connected       bool
	ClientAddress   string
	ProxyStrategy   string
	ProxyList       map[string]bool
	ProxyBypassList map[string]bool
}

func NewState(configPath string) *State {
	return &State{
		configPath:      configPath,
		Connected:       false,
		ClientAddress:   "",
		ProxyStrategy:   "none",
		ProxyList:       map[string]bool{},
		ProxyBypassList: map[string]bool{},
	}
}

func (s *State) SetClientAddress(address string) {
	s.ClientAddress = address
}

func (s *State) SetProxyStrategy(strategy string) {
	s.ProxyStrategy = strategy
}

func (s *State) PutProxyList(domain string) {
	s.ProxyList[domain] = true
}

func (s *State) DeleteProxyList(domain string) {
	delete(s.ProxyList, domain)
}

func (s *State) PutProxyBypassList(domain string) {
	s.ProxyBypassList[domain] = true
}

func (s *State) DeleteProxyBypassList(domain string) {
	delete(s.ProxyBypassList, domain)
}
