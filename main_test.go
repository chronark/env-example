package main

import (
	"io/ioutil"
	"testing"
)

func Test_main(t *testing.T) {

	t.Run("Removes all values but keeps comments and empty lines", func(t *testing.T) {
		main()
		file, err := ioutil.ReadFile(".env.example")
		if err != nil {
			panic(err)
		}

		expectedContent := `# Comment

# Comment with equal sign: =
HELLO=

# Empty value
NO_VALUE=
`

		if string(file) != expectedContent {
			t.Errorf("Wrong output, want: %s, got: %s", expectedContent, string(file))
		}

	})
}
