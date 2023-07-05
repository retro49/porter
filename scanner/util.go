package scanner

import (
    "encoding/json"
    "errors"
    "sync"
    "os"
)

const JSON_PATH string = "/usr/share/porter/ports.json"
const JSON_TEST_PATH string = "$HOME/tests/porter/ports.json"

var LOADER_ERROR_READING_FILE_SIZE error = errors.New("error reading file size")
var LOADER_ERROR_OPENING_FILE error = errors.New("error opening file")
var LOADER_ERROR_READING_CONTENT error = errors.New("error reading file content")


// lets call it a pojo.
// used for storing information about a port.
type portInfo struct {
    Name string
    Description string
    Port int
}

// Enables to create a property about a 
// specific port number
// with the name of the service and 
// description about the service.
func NewPortInfo(name, description string, port int)portInfo {
    return portInfo{
        Name: name,
        Description: description,
        Port: port,
    }
}

// returns the name of the port.
func (p portInfo)GetName() string {
    if p.Name == ""{
        return "unknown"
    }
    return p.Name
}

// returns the description of the port.
func (p portInfo)GetDescription()string{
    if p.Description == ""{
        return "unkown"
    }
    return p.Description
}


// returns the port number
func (p portInfo)GetPort()int{
    return p.Port
}

// reads the json file and returns the stream
func readJSON()([]byte, error){
    var fileSize int64 = 0
    // get the file status
    if status, err := os.Stat(JSON_TEST_PATH); err != nil{
        return nil, LOADER_ERROR_READING_FILE_SIZE 
    } else {
        fileSize = status.Size()
    }
    buffer := make([]byte, fileSize)
    // read into the file
    jsonFile, err := os.Open(JSON_TEST_PATH)
    if err != nil {
        return nil, LOADER_ERROR_OPENING_FILE 
    }

    if _, err := jsonFile.Read(buffer); err != nil{
        return nil, LOADER_ERROR_READING_CONTENT 
    }

    jsonFile.Close()

    return buffer, nil
}

// loads the json and sends the result in the channel
func LoadPortInfo(ch chan any, wg *sync.WaitGroup){
    defer wg.Done()
    // validate the json file
    jsonData, err := readJSON()
    if err != nil {
        ch <- nil
    }

    validJSON := json.Valid(jsonData)
    if !validJSON {
        ch <- nil
    }
    // json unmarshaller
    portInfo := make(map[string]map[string]string)
    if err := json.Unmarshal(jsonData, &portInfo); err != nil{
        ch <- nil
    }

    ch <- portInfo
}