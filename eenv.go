package main

import (
	"errors"
	"os"
	"regexp"
	"strings"
)

const ENV = ".env"
const envVarRegex = `\w*[\s+]*=([\s+]*\w).+`

type envVar struct {
	Key   string
	Value string
}

// get environment variables from .env file
func GetEnvVars() ([]envVar, error) {
	// compile regex
	regex, err := regexp.Compile(envVarRegex)
	if err != nil {
		return []envVar{}, err
	}

	// create return value
	var result []envVar

	// read .env file
	envByte, err := os.ReadFile(ENV)
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

		newEnvVar.Key = newEnvVarKey
		newEnvVar.Value = newEnvVarValue

		*result = append(*result, newEnvVar)
	}
	return nil
}
