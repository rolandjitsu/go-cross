package hello

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	stdout := os.Stdout // Keep a backup of the current stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Greet("John")

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = stdout

	assert.Equal(t, "Hello, John!\n", string(out))
}

func TestGreetMsgPack(t *testing.T) {
	stdout := os.Stdout // Keep a backup of the current stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := GreetMsgPack("John")
	if !assert.NoError(t, err) {
		return
	}

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = stdout

	assert.Equal(t, "Hello, John! from msgpack\n", string(out))
}
