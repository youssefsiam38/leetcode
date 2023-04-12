func simplifyPath(path string) string {
	dirs := strings.Split(path, "/")
	stack := []string{}
	out := ""
	for i := 0; i < len(dirs); i++ {
		switch dirs[i] {
		case "..":
			removeLast(&stack)
		case "":
			continue
		case ".":
			continue
		default:
			stack = append(stack, dirs[i])
		}
	}
	for _, s := range stack {
		out += "/" + s
	}
	if out == "" {
		return "/"
	}
	return out
}

func removeLast(_slice *[]string) {
	if len(*_slice) > 0 {
		temp := *_slice
		*_slice = temp[:len(temp)-1]
	}
}
