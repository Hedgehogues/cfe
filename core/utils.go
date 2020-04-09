package core

func PToStrings(e []*Extract) []string {
	var objs []string
	for _, el := range e {
		objs = append(objs, el.Object)
	}
	return objs
}

func ToStrings(e []Extract) []string {
	var objs []string
	for _, el := range e {
		objs = append(objs, el.Object)
	}
	return objs
}

