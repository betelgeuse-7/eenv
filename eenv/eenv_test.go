package eenv

import (
	"testing"
)

func TestParseEnv(t *testing.T) {
	want := map[string]string{
		"key":   "ACCESS_TOKEN_SECRET",
		"value": "someSecret_Phrase",
	}

	var got map[string]string = map[string]string{
		"key":   "",
		"value": "",
	}

	key, err := parseEnv("ACCESS_TOKEN_SECRET=someSecret_Phrase", true)
	if err != nil {
		t.Fail()
	}

	value, err := parseEnv("ACCESS_TOKEN_SECRET=someSecret_Phrase", false)
	if err != nil {
		t.Fail()
	}

	got["key"] = key
	got["value"] = value

	if got["key"] != want["key"] || got["value"] != want["value"] {
		t.Errorf("expected %v got %v\n", want, got)
	}
}

/*
func TestMakeEnvVarSlice(t *testing.T) {
	matches := [][]byte{
		{68, 65, 6, 20, 77, 6, 72},
		{69, 20, 68, 70, 65, 20},
		{61, 72, 65, 20, 67, 64},
	}
	var toPopulate []envVar

	err := makeEnvVarSlice(matches, &toPopulate)
	if err != nil {
		t.Fail()
	}

	if len(toPopulate) == 0 {
		t.Errorf("couldnt populate")
	}

	fmt.Println("populated", toPopulate)
}
*/
