package services

import (
	"github.com/owulveryck/restmdw/core/authentication"
	"github.com/owulveryck/toscalib"
	"io/ioutil"
	"log"
	//"github.com/owulveryck/restmdw/services/models"
	"encoding/json"
	"github.com/owulveryck/restmdw/settings"
	"gopkg.in/yaml.v2"
	"net/http"
	"os"
	"path/filepath"
)

// OfferingList parse the yaml files from the repository configuration file defined in the seettings package
// And then returns it json representation
func OfferingList(r *http.Request) (int, []byte, error) {
	settings := settings.Get()
	rootRepository := settings.OfferingRepositoryPath
	var response []byte
	sendOffering := func(path string, info os.FileInfo, err error) error {
		stat, err := os.Stat(path)
		if err != nil {
			return err

		}

		if stat.IsDir() && path != rootRepository {
			return filepath.SkipDir

		}
		matched, err := filepath.Match("*.yaml", info.Name())
		if err != nil {
			return err // this is fatal.

		}
		if matched {
			r, err := os.Open(path)
			if err != nil {
				return err
			}
			data, err := ioutil.ReadAll(r)
			defer r.Close()
			if err != nil {
				return err
			}
			var nodeTypes map[string]toscalib.NodeType
			err = yaml.Unmarshal(data, &nodeTypes)
			if err != nil {
				return err
			}
			response, err = json.MarshalIndent(nodeTypes, "", "  ")
			if err != nil {
				return err
			}
		}
		return nil

	}
	err := filepath.Walk(rootRepository, sendOffering)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, response, nil
}
