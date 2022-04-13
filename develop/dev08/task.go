package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/process"
	"os"
	"strconv"
	"strings"
)

/*
Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*


Так же требуется поддерживать функционал fork/exec-команд

Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).

*/

func main() {
	fmt.Println("Go Shell")
	reader := bufio.NewScanner(os.Stdin)
	for {
		dir, _ := os.Getwd()
		fmt.Printf("%s$> ", dir)
		reader.Scan()
		input := reader.Text()
		err := execCommands(input)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execCommand(command string) error {

	command = strings.TrimSuffix(command, "\n")
	args := strings.Split(command, " ")
	var ErrNoPath = errors.New("arg required")
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return ErrNoPath
		}
		return os.Chdir(args[1])

	case "pwd":
		dir, pwdErr := os.Getwd()
		fmt.Println("PWD result:", dir)
		return pwdErr

	case "echo":
		if len(args) < 2 {
			return ErrNoPath
		}
		_, err := fmt.Fprintln(os.Stdout, args[1:])
		return err
	case "kill":

		var processId int
		var process *os.Process
		var err error
		if len(args) < 2 {
			return ErrNoPath
		}

		if processId, err = strconv.Atoi(args[1]); err != nil {
			return err
		}
		if process, err = os.FindProcess(processId); err != nil {
			return err
		}
		if err = process.Kill(); err != nil {
			return err
		}
		return err

	case "ps":
		infoStat, err := host.Info()
		if err != nil {
			return err
		}
		fmt.Printf("Total processes: %d\n", infoStat.Procs)
		processSlice, err := process.Processes()
		for _, val := range processSlice {
			name, _ := val.Name()
			status, _ := val.Status()
			fmt.Println(val.Pid, name, status)

		}
	case "quit":
		os.Exit(0)
	}

	return nil
}

func execCommands(input string) error {

	input = strings.TrimSuffix(input, "\n")
	commands := strings.Split(input, "|")
	var err error
	for _, val := range commands {
		err = execCommand(val)

	}
	return err
}
