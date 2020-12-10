package utils

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

func PrettyPrint(obj interface{}) {
	b, err := yaml.Marshal(obj)
	if err != nil {
		fmt.Printf("unable to prettyprint: %s\n", err)
	}

	fmt.Println(string(b))
}
