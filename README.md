# EENV

It stands for 'Easy ENV'.

---

If you have a .env file in your project, and you want to get variables in that file:
Use this format in .env:
    SOME_VARIABLE=someVariablesValue
    ANOTHER_VARIABLE=ANOTHER_variables_VAluE

Do not use empty variables because I couldn't implement the regex in 'Golang'. Golang (regexp) unfortunately doesn't support lookaheads and lookbehinds.
I am not such a regex wizard, so...

---

The package is extremely easy to use

```go
    package main

    import "github.com/betelgeuse-7/eenv"

    func main() {
        envVars, err := eenv.GetEnvVars()
        if err != nil {
            ...
        }

        // envVars is of []eenv.EnvVar type. 
        // eenv.EnvVar type is a simple struct consisting of two values:
        //              Key string
        //              Value string
        // you can use this type to store your individual environment variables.

        SOME_VARIABLE := envVars[0]
        ANOTHER_VARIABLE := envVars[1]

        SOME_VARIABLE_VALUE := SOME_VARIABLE.Value

        fmt.Println(SOME_VARIABLE)         // {SOME_VARIABLE someVariablesValue}

        fmt.Println(SOME_VARIABLE.Value) // someVariablesValue
        fmt.Println(SOME_VARIABLE.Key)  // SOME_VARIABLE
        
        fmt.Println(ANOTHER_VARIABLE) // {ANOTHER_VARIABLE ANOTHER_variables_VAluE}
    }
```

Good coding...