variable "DEST" {
	default = "./bin"
}

target "default" {
  dockerfile  = "Dockerfile.hello"
  output     = ["${DEST}"]
}
