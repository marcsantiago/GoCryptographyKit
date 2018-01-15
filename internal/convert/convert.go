package convert

import (
	"fmt"
	"io/ioutil"
	"os"
)

// RetrieveDataFromStringOrFile ...
func RetrieveDataFromStringOrFile(i interface{}) (string, error) {
	switch i.(type) {
	case string:
		return i.(string), nil
	case *os.File:
		f := i.(*os.File)
		defer f.Close()
		b, err := ioutil.ReadAll(f)
		if err != nil {
			return "", err
		}
		return string(b), nil
	default:
		return "", fmt.Errorf("Message must be of type string or file")
	}
}
