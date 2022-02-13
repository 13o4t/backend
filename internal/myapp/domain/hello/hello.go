package hello

func Hello(name string) string {
	if name == "" {
		return "Hello!\n"
	}
	return "Hello " + name + "!\n"
}
