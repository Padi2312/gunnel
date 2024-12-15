package internal

import (
	"context"
	"encoding/json"
	"os"
	"path"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Config struct {
	Username          string   `json:"username"`
	Tunnels           []Tunnel `json:"tunnels"`
	SSHPrivateKeyPath string   `json:"sshPrivateKeyPath"`
}

type Store struct {
	root   string
	Config *Config
	ctx    context.Context
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (s *Store) Startup(ctx context.Context) {
	s.ctx = ctx
}

func NewStore() *Store {
	// Get user's home directory
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	// Create a new directory named "gunnel" in the user's home directory
	root := path.Join(home, "/gunnel")
	if _, err := os.Stat(root); os.IsNotExist(err) {
		os.Mkdir(root, 0755)
	}

	// Return a new Store instance
	return &Store{
		root: root,
	}
}

func (s *Store) GetConfig() *Config {
	return s.Config
}

func (s *Store) Init() {
	// Check if the store is already initialized with `gunnel.config.json`
	_, err := os.Stat(path.Join(s.root, "gunnel.config.json"))
	if os.IsNotExist(err) {
		// If the file does not exist, create a new store with an empty configuration
		s.Config = &Config{
			Username: "",
			Tunnels:  []Tunnel{},
		}
		s.Save()
		return
	}

	// Read the contents of the file
	file, err := os.Open(path.Join(s.root, "gunnel.config.json"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var config *Config = &Config{}
	// Decode the contents of the file into the Store instance
	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		panic(err)
	}

	s.Config = config
}

func (s *Store) SetUsername(username string) {
	s.Config.Username = username
	s.Save()
}

func (s *Store) AddTunnel(tunnel Tunnel) {
	// Create a new Tunnel instance
	// Append the new tunnel to the store
	s.Config.Tunnels = append(s.Config.Tunnels, tunnel)

	// Save the store
	s.Save()
}

func (s *Store) RemoveTunnel(id string) {
	// Find the tunnel with the given id
	for i, tunnel := range s.Config.Tunnels {
		if tunnel.ID == id {
			// Remove the tunnel from the store
			s.Config.Tunnels = append(s.Config.Tunnels[:i], s.Config.Tunnels[i+1:]...)
			break
		}
	}

	// Save the store
	s.Save()
}

func (s *Store) Save() {
	// Open the file for writing
	file, err := os.Create(path.Join(s.root, "gunnel.config.json"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Encode the store into the file
	encoder := json.NewEncoder(file)
	err = encoder.Encode(s.Config)
	if err != nil {
		panic(err)
	}
}

func (s *Store) SelectAndSetSSHPrivateKeyPath() string {
	file, err := runtime.OpenFileDialog(s.ctx, runtime.OpenDialogOptions{
		Title: "Select a file",
	})
	if err != nil {
		panic(err)
	}
	if file != "" {
		s.Config.SSHPrivateKeyPath = file
	} else {
		println("No file selected.")
	}

	s.Save()

	return s.Config.SSHPrivateKeyPath
}
