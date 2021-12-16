package commands

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/software-engineering-components/go-arch-lab4/engine"
)

var commandsArr = []engine.Command{
	&print{},
	&reverse{},
}

func defineField(field reflect.Value, str string) error {
	var val interface{}
	var err error

	switch field.Type().Name() {
	case "int":
		val, err = strconv.Atoi(str)
	case "float":
		val, err = strconv.ParseFloat(str, 32)
	default:
		val = str
	}

	if err != nil {
		return err
	}

	if field.Type() != reflect.ValueOf(val).Type() {
		return fmt.Errorf("error: wrong arg")
	}

	field.Set(reflect.ValueOf(val))

	return nil
}

func setParameters(cmdReflection reflect.Value, args []string) error {
	cmdReflectionElem := cmdReflection.Elem()
	if cmdReflectionElem.NumField() != len(args) {
		return fmt.Errorf("error: wrong number of args")
	}

	for i, v := range args {
		field := cmdReflectionElem.Field(i)
		err := defineField(field, v)
		if err != nil {
			return err
		}
	}

	return nil

}

func dublicateCommand(oldObj interface{}) interface{} {
	newObj := reflect.New(reflect.TypeOf(oldObj).Elem())
	oldVal := reflect.ValueOf(oldObj).Elem()
	newVal := newObj.Elem()
	for i := 0; i < oldVal.NumField(); i++ {
		newValField := newVal.Field(i)
		if newValField.CanSet() {
			newValField.Set(oldVal.Field(i))
		}
	}
	return newObj.Interface()
}

func pipe(cmdName string, args []string) engine.Command {
	var command engine.Command

	for _, v := range commandsArr {
		commandValue := reflect.ValueOf(v)
		name := commandValue.Elem().Type().Name()
		if cmdName == name {
			setParameters(commandValue, args)
			command = dublicateCommand(v).(engine.Command)
			break
		}
	}
	return command
}

func Parse(line string) engine.Command {
	splittedLine := strings.Fields(line)

	command, params := splittedLine[0], splittedLine[1:]
	return pipe(command, params)
}
