package greeter

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGreeter(t *testing.T) {
	// TODO
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	Greet()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	s := string(out[:])

	expected := `                                        ` + `
      _                 _               ` + `
 ___ | |_   __ _  _ __ | |_   ___  _ __ ` + `
/ __|| __| / _` + "`" + ` || '__|| __| / _ \| '__|
\__ \| |_ | (_| || |   | |_ |  __/| |   ` + `
|___/ \__| \__,_||_|    \__| \___||_|   ` + `
`

	if s != expected {
		t.Errorf("\n %s \n not equal, expected: \n %s", s, expected)
	}

}
