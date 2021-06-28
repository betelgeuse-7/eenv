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
const EnvVarRegex = `\w*[\s+]*=([\s+]*\w).+`

/*
* a struct representing an environment variable
* This is exported so that users can store their 
* environment variables using this struct.
*/
type EnvVar struct {
	Key   string
	Value string
}

/*
* Parse environment variables in the .env file in the current directory.
* and return all the variables with their keys as a slice of EnvVars.
* EnvVar struct comprises of a Key (string) and a Value (string). 
 */
func GetEnvVars() ([]EnvVar, error) {
	// compile regex
	regex, err := regexp.Compile(EnvVarRegex)
	if err != nil {
		return []EnvVar{}, err
	}

	// create return value
	var result []EnvVar

	// read .env file
	envByte, err := os.ReadFile(ENV)
	if err != nil {
		return []EnvVar{}, err
	}

	// get all the matches
	// -1 -> we want all the matches (no limit)
	matches := regex.FindAll(envByte, -1)

	if err := makeEnvVarSlice(matches, &result); err != nil {
		return []EnvVar{}, err
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
func makeEnvVarSlice(matches [][]byte, result *[]EnvVar) error {
	for _, v := range matches {
		var newEnvVar EnvVar

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
