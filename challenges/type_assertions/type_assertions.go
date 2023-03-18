package typeassertions

import "errors"

type Developer struct {
	Name string
	Age  int
}

func GetDeveloper(name interface{}, age interface{}) (Developer, error) {
	var d Developer
	switch a := age.(type) {
	case int:
		d.Age = a
	default:
		return Developer{}, errors.New("Developer age must be a number")
	}

	switch n := name.(type) {
	case string:
		d.Name = n
	default:
		return Developer{}, errors.New("Developer age must be a number")
	}

	return d, nil
}
