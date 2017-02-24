package transform

import (
	"fmt"

	simplejson "github.com/bitly/go-simplejson"
)

// Extract returns the specified path as the top-level object.
func Extract(spec *Config, data *simplejson.Json) (*simplejson.Json, error) {
	outPath, ok := (*spec.Spec)["path"]
	if !ok {
		return nil, &Error{ErrMsg: fmt.Sprintf("Unable to get path"), ErrType: SpecError}
	}
	outData, err := getJSONPath(data, outPath.(string), spec.Require)
	return outData, err
}