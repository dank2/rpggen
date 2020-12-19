package utils

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

func PrettyFormat(obj interface{}) (string, error) {
	b, err := yaml.Marshal(obj)
	if err != nil {
		return "", fmt.Errorf("unable to format: %s", err)
	}
	return string(b), nil
}
