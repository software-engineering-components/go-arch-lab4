package commands

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/software-engineering-components/go-arch-lab4/engine"
)

func setParameters(reflectionObject reflect.Value, args []string) error {
	refrectE := reflectionObject.Elem()
	if refrectE.NumField() != len(args) {
		return fmt.Errorf("error: wrong number of args")
	}

	for index, value := range args {
		refrectE.Field(index).Set(reflect.ValueOf(value))
	}

	return nil

}

func dublicateCommand(oldObj interface{}) interface{} {
	next := reflect.New(reflect.TypeOf(oldObj).Elem())
	prev := reflect.ValueOf(oldObj).Elem()
	current := next.Elem()
	for i := 0; i < prev.NumField(); i++ {
		newValField := current.Field(i)
		if newValField.CanSet() {
			newValField.Set(prev.Field(i))
		}
	}
	return next.Interface()
}

func Parse(line string) engine.Command {
	splittedLine := strings.Fields(line)

	cmd, params := splittedLine[0], splittedLine[1:]

	var command engine.Command

	for _, value := range []engine.Command{
		&print{},
		&reverse{},
	} {
		val := reflect.ValueOf(value)
		name := val.Elem().Type().Name()
		if cmd == name {
			setParameters(val, params)
			command = dublicateCommand(value).(engine.Command)
			break
		}
	}
	return command
}
