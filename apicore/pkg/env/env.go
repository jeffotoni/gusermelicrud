package env

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var debug = !GetProd()

func GetInt(name string, defaultValue int) int {
	if stringValue, ok := os.LookupEnv(name); ok {
		if val, ok := strconv.Atoi(stringValue); ok == nil {
			if debug {
				fmt.Printf("Looking for ENV: [%s] found value %v\n", name, val)
			}
			return val
		}
		if debug {
			fmt.Printf("Looking for ENV: [%s] found value %s, but could not parse to int, therefore returning default value [%v]\n", name, stringValue, defaultValue)
		}
		return defaultValue
	}
	if debug {
		fmt.Printf("Looking for ENV: [%s], but could not find the value, therefore returning default value [%v]\n", name, defaultValue)
	}
	return defaultValue
}

func GetDuration(name string, defaultValue time.Duration) time.Duration {
	if stringValue, ok := os.LookupEnv(name); ok {
		if val, ok := strconv.Atoi(stringValue); ok == nil {
			if debug {
				fmt.Printf("Looking for ENV: [%s] found value %v\n", name, val)
			}
			return time.Duration(val)
		}
		if debug {
			fmt.Printf("Looking for ENV: [%s] found value %s, but could not parse to duration, therefore returning default value [%v]\n", name, stringValue, defaultValue)
		}
		return defaultValue
	}
	if debug {
		fmt.Printf("Looking for ENV: [%s], but could not find the value, therefore returning default value [%v]\n", name, defaultValue)
	}
	return defaultValue
}

func GetString(name string, defaultValue string) string {
	if val, ok := os.LookupEnv(name); ok {
		if debug {
			fmt.Printf("Looking for ENV: [%s] found value %v\n", name, val)
		}
		return val
	}
	if debug {
		fmt.Printf("Looking for ENV: [%s], but could not find the value, therefore returning default value [%v]\n", name, defaultValue)
	}
	return defaultValue
}

func GetBool(name string, defaultValue bool) bool {
	if stringValue, ok := os.LookupEnv(name); ok {
		if val, ok := strconv.ParseBool(stringValue); ok == nil {
			if debug {
				fmt.Printf("Looking for ENV: [%s] found value %v\n", name, val)
			}
			return val
		}
		if debug {
			fmt.Printf("Looking for ENV: [%s] found value %s, but could not parse to boolean, therefore returning default value [%v]\n", name, stringValue, defaultValue)
		}
		return defaultValue
	}
	if debug {
		fmt.Printf("Looking for ENV: [%s], but could not find the value, therefore returning default value [%v]\n", name, defaultValue)
	}
	return defaultValue
}

func GetProd() bool {
	if val, ok := os.LookupEnv("ENV"); ok {
		return "prd" == val
	}
	return false
}
