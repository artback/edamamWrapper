package read

import (
	"io"
	"io/ioutil"
)

func ToString(reader io.Reader) (*string, error) {
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	str := string(b)
	return &str, err

}

