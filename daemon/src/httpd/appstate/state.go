package appstate

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"socks-manager/src/platform/freeport"
	"strconv"
)

var ProxyStrategyType = struct {
	BypassAll       string
	ProxyAll        string
	BypassAllExcept string
	ProxyAllExcept  string
}{
	BypassAll:       "BypassAll",
	ProxyAll:        "ProxyAll",
	BypassAllExcept: "BypassAllExcept",
	ProxyAllExcept:  "ProxyAllExcept",
}

type State struct {
	configPath      string
	connected       bool
	connectionPort  int
	command         *exec.Cmd
	ClientAddress   string          `json:"clientAddress"`
	ProxyStrategy   string          `json:"proxyStrategy"`
	ProxyList       map[string]bool `json:"proxyList"`
	ProxyBypassList map[string]bool `json:"proxyBypassList"`
}

func NewState(configPath string) *State {
	state := &State{
		configPath:      configPath,
		connected:       false,
		ClientAddress:   "",
		ProxyStrategy:   "none",
		ProxyList:       map[string]bool{},
		ProxyBypassList: map[string]bool{},
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		state.persist()
	} else {
		state.loadPersistant()
	}

	return state
}

func (s *State) Connect() error {
	if s.connected == true {
		return errors.New("Already connected")
	}

	inputReader, inputWriter := io.Pipe()
	freeport, _ := freeport.GetFreePort()

	s.connectionPort = freeport
	s.command = exec.Command("ssh", "-D", strconv.Itoa(freeport), s.ClientAddress)

	s.command.Stdin = inputReader

	go io.Copy(inputWriter, os.Stdin)

	s.command.Start()

	s.connected = true

	return nil
}

func (s *State) Disconnect() {
	s.command.Process.Kill()
	s.connected = false
}

func (s *State) IsConnected() bool {
	return s.connected
}

func (s *State) GetConnectionPort() int {
	return s.connectionPort
}

func (s *State) SetClientAddress(address string) {
	s.ClientAddress = address
	s.persist()
}

func (s *State) SetProxyStrategy(strategy string) {
	s.ProxyStrategy = strategy
	s.persist()
}

func (s *State) PutProxyList(domain string) {
	s.ProxyList[domain] = true
	s.persist()
}

func (s *State) DeleteProxyList(domain string) {
	delete(s.ProxyList, domain)
	s.persist()
}

func (s *State) PutProxyBypassList(domain string) {
	s.ProxyBypassList[domain] = true
	s.persist()
}

func (s *State) DeleteProxyBypassList(domain string) {
	delete(s.ProxyBypassList, domain)
	s.persist()
}

func (s *State) persist() {
	out, _ := json.MarshalIndent(s, "", "  ")
	ioutil.WriteFile(s.configPath, out, 0644)
}

func (s *State) loadPersistant() {
	data, _ := ioutil.ReadFile(s.configPath)
	json.Unmarshal(data, s)
}
