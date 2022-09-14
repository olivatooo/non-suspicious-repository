package actions

import (
	"encoding/json"
	"errors"
	"io"
)

func GetJSONParametersFromBody(body *io.ReadCloser, jsonStructure interface{}) (interface{}, error) {
	var unmarshalErr *json.UnmarshalTypeError
	decoder := json.NewDecoder(*body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(jsonStructure)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			return nil, err
		}
	}
	return jsonStructure, nil
}
