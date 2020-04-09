package core

func pToStrings(e []*Extract) []string {
	var objs []string
	for _, el := range e {
		objs = append(objs, el.Object)
	}
	return objs
}

func toStrings(e []Extract) []string {
	var objs []string
	for _, el := range e {
		objs = append(objs, el.Object)
	}
	return objs
}

