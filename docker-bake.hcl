variable "DEST" {
  default = "./bin"
}

target "hello-arm64" {
  dockerfile = "Dockerfile.hello"
  args = {
    TARGETARCH = "arm64"
  }
  output = ["${DEST}/arm64"]
}

target "hello-arm" {
  dockerfile = "Dockerfile.hello"
  args = {
    TARGETARCH = "arm"
  }
  output = ["${DEST}/arm"]
}

target "default" {
  dockerfile = "Dockerfile.hello"
  args = {
    TARGETARCH = "arm64"
  }
  output = ["${DEST}"]
}
