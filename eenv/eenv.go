package eenv

import (
	"errors"
	"os"
	"regexp"
	"strings"
)

/*
* Spaces are allowed after both sides of =
* No value is not allowed.
* at least 2 characters after = sign.
*
* (I could do better but unfotunately golang doesn't support lookbehinds and lookforwards)
 */

const ENV = ".env"
const envVarRegex = `\w*[\s+]*=([\s+]*\w).+`

/*
* a struct representing an environment variable
 */
type envVar struct {
	Key   string
	Value string
}

/*
* specify the path to .env file using this struct.
* and call GetEnvVars method on it.
 */
type EnvFileLoc struct {
	Path string
}

// TODO Find a solution to paths !
/*
* set EnvFileLoc's Path to project root
*
 */
func (efl *EnvFileLoc) SetEnvPath(path string) EnvFileLoc {
	// if user included .env at the end of the path:
	// delete it.
	path = strings.Replace(path, ".env", "", -1)

	if string(path[len(path)-1]) == "." {
		path[len(path)-1] = ""
	}

	efl.Path = path

	return *efl
}

/*
* Parse environment variables in the .env file in the current directory.
* and return all the variables with their keys as a slice of envVars.
* envVar struct comprises of a Key (string) and a Value (string).
 */
func (efl EnvFileLoc) GetEnvVars() ([]envVar, error) {
	// compile regex
	regex, err := regexp.Compile(envVarRegex)
	if err != nil {
		return []envVar{}, err
	}

	// create return value
	var result []envVar

	// read .env file
	envByte, err := os.ReadFile(efl.Path)
	if err != nil {
		return []envVar{}, err
	}

	// get all the matches
	// -1 -> we want all the matches (no limit)
	matches := regex.FindAll(envByte, -1)

	if err := makeEnvVarSlice(matches, &result); err != nil {
		return []envVar{}, err
	}

	return result, nil
}

/*
* parse env config variable to get its key or value
* variable -> a .env config variable
* key -> whether parseEnv should return the key of variable or not
 */
func parseEnv(variable string, key bool) (string, error) {
	if len(variable) == 0 {
		return "", errors.New("env variable empty")
	}
	if key {
		return strings.Split(variable, "=")[0], nil
	}
	return strings.Split(variable, "=")[1], nil
}

/*
* populate the return value with matches
 */
func makeEnvVarSlice(matches [][]byte, result *[]envVar) error {
	for _, v := range matches {
		var newEnvVar envVar

		newEnvVarKey, err := parseEnv(string(v), true)
		if err != nil {
			return err
		}

		newEnvVarValue, err := parseEnv(string(v), false)
		if err != nil {
			return err
		}

		newEnvVar.Key = strings.TrimSpace(newEnvVarKey)
		newEnvVar.Value = strings.TrimSpace(newEnvVarValue)

		*result = append(*result, newEnvVar)
	}
	return nil
}
