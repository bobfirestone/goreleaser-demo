package main

func main() {
	Foo("v0.1.4")
}

func Foo(version string) {
	println("Hello goreleaser bump to " + version)
}
