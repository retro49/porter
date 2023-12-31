package scanner

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/retro49/porter/plogger"
)

const JSON_PATH = "/usr/share/porter/ports.json"

var LOADER_ERROR_READING_FILE_SIZE error = errors.New("error reading file size")
var LOADER_ERROR_OPENING_FILE error = errors.New("error opening file")
var LOADER_ERROR_READING_CONTENT error = errors.New("error reading file content")

// lets call it a pojo.
// used for storing information about a port.
type portInfo struct {
	Name        string
	Description string
	Port        string
}

// Enables to create a property about a
// specific port number
// with the name of the service and
// description about the service.
func NewPortInfo(port, name, description string) portInfo {
	return portInfo{
		Name:        name,
		Description: description,
		Port:        port,
	}
}

// returns the name of the port.
func (p portInfo) GetName() string { return p.Name }

// returns the description of the port.
func (p portInfo) GetDescription() string { return p.Description }

func (p portInfo) GetPort() string { return p.Port }

// reads the json file and returns the stream
func readJSON() ([]byte, error) {
	var fileSize int64 = 0
	// get the file status
	if status, err := os.Stat(JSON_PATH); err != nil {
		return nil, LOADER_ERROR_READING_FILE_SIZE
	} else {
		fileSize = status.Size()
	}
	buffer := make([]byte, fileSize)
	// read into the file
	jsonFile, err := os.Open(JSON_PATH)
	if err != nil {
		return nil, LOADER_ERROR_OPENING_FILE
	}

	if _, err := jsonFile.Read(buffer); err != nil {
		return nil, LOADER_ERROR_READING_CONTENT
	}

	jsonFile.Close()

	return buffer, nil
}

// loads the json and sends the result in the channel
func LoadPortInfo(ch chan any) {
	// validate the json file
	jsonData, err := readJSON()
	if err != nil {
		plogger.NewPlogger().Error("read json error", "unable to read port json file")
		ch <- nil
	}

	validJSON := json.Valid(jsonData)
	if !validJSON {
		plogger.NewPlogger().Error("invalid json", "error while validating json")
		ch <- nil
	}
	// json unmarshaller
	portInfo := make(map[string]map[string]string)
	if err := json.Unmarshal(jsonData, &portInfo); err != nil {
		plogger.NewPlogger().Error("error decodig", "error while decoding json")
		ch <- nil
	}
	ch <- &portInfo
}

// a scan info for  providing the scanner
type ScanInfo struct {
	Network   string
	Host      string
	StartPort int
	EndPort   int
	Step      int
	Skip      []int
	Threads   int
	Timeout   int
	Format    string
	Output    string
}

func (s ScanInfo) GetNetwork() string { return s.Network }
func (s ScanInfo) GetHost() string    { return s.Host }
func (s ScanInfo) GetStart() int      { return s.StartPort }
func (s ScanInfo) GetEnd() int        { return s.EndPort }
func (s ScanInfo) GetStep() int       { return s.Step }
func (s ScanInfo) GetSkip() []int     { return s.Skip }
func (s ScanInfo) GetThreads() int    { return s.Threads }
func (s ScanInfo) GetFormat() string  { return s.Format }
func (s ScanInfo) GetOutput() string  { return s.Output }
func (s ScanInfo) GetTimeout() int    { return s.Timeout }

func ToStringArr(ports []int, service string) []string {
	res := make([]string, 0)
	for _, port := range ports {
		prt := fmt.Sprintf("%d/%s", port, service)
		res = append(res, prt)
	}
	return res
}

func InJson(jsonData *map[string]map[string]string, key string) portInfo {
	var name string
	var description string
	jsonValue, found := (*jsonData)[key]
	if !found {
		return NewPortInfo(key, "unkown", "unkown")
	}
	name, found = jsonValue["name"]
	if !found {
		name = "unkownn"
	}
	description, found = jsonValue["description"]
	if !found {
		description = "unkown"
	}
	return NewPortInfo(key, name, description)
}
