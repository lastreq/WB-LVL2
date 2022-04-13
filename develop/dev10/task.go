package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123
go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

type connInfo struct {
	host    string
	port    string
	timeout time.Duration
}

func readConnInfo() connInfo {

	flag.Usage = func() {
		fmt.Println("Usage flags: [--timeout] host port")
		flag.PrintDefaults()
	}
	var timeout time.Duration
	flag.DurationVar(&timeout, "timeout", time.Second*10, "timeout")
	flag.Parse()
	args := flag.Args()

	if len(args) != 2 {
		flag.Usage()
		os.Exit(1)
	}

	s := connInfo{
		host:    args[0],
		port:    args[1],
		timeout: timeout,
	}

	return s
}

func readFromServer(conn net.Conn, ErrChan chan<- error) {
	scanner := bufio.NewScanner(conn)
	for {
		if !scanner.Scan() {
			ErrChan <- scanner.Err()
		}
		text := scanner.Text()
		fmt.Printf("%s\n", text)
	}
}

func writeToServer(conn net.Conn, signalChan chan<- os.Signal, ErrChan chan<- error) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			ErrChan <- scanner.Err()
		}
		str := scanner.Text()

		sl := strings.Split(fmt.Sprintf("%x", str), " ")
		for _, u := range sl {
			if u == "04" {
				fmt.Println("отправлен ctrl+D")
				signalChan <- syscall.SIGQUIT

			}

		}
		_, err := conn.Write([]byte(fmt.Sprintln(str)))
		if err != nil {
			//ошибка записи на сервер
			ErrChan <- err
		}
	}
}

func main() {

	s := readConnInfo()
	signalChan := make(chan os.Signal, 1)
	ErrChan := make(chan error, 1)
	signal.Notify(signalChan, syscall.SIGINT)

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(s.host, s.port), s.timeout)
	if err != nil {

		fmt.Println(err)
		os.Exit(0)
	}

	go readFromServer(conn, ErrChan)
	go writeToServer(conn, signalChan, ErrChan)
	//Закрытие клиента либо по сигналу завершения работы, либо при получении ошибки
	select {
	case <-signalChan:
		fmt.Println("telnet клиент остановлен")
	case err = <-ErrChan:

		fmt.Println(err)
		conn.Close()
	}
	fmt.Println("telnet клиент закрыт")
}
