package helpers

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/kataras/iris"
)

var FailOnError = failOnError

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func GetRequiredEnvString(name string) string {
	envVar := os.Getenv(name)

	if len(envVar) == 0 {
		err := fmt.Errorf("not defined %s in environ", name)
		FailOnError(err, "ERROR")
	}

	return envVar
}

func GetEnvStringWithDefaultValue(name string, defValue string) string {
	envVar := os.Getenv(name)

	if len(envVar) == 0 {
		envVar = defValue
	}

	return envVar
}

func GetEnvIntWithDefaultValue(name string, defValue int) int {
	var res int

	value, err := strconv.Atoi(os.Getenv(name))
	if err != nil {
		res = defValue
	} else {
		res = value
	}

	return res
}

func SetIrisCtxError(ctx iris.Context, err error, statusCode int) {
	ctx.Application().Logger().Error(err)

	ctx.Values().Set("error", err.Error())
	ctx.StatusCode(statusCode)
}

func LogInfoError(method string, err error) {
	if err != nil {
		log.Printf("%s failed. Error: %s.\n", method, err)
	}
}
