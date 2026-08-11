package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"syscall"
)

var pwd string
var ctx map[string]int

func parseToken(token string) int {
	if token[0] == '$' {
		val, ok := ctx[token[1:]]
		if !ok {
			return 0
		}
		return val
	}

	val, err := strconv.Atoi(token)
	if err != nil {
		return 0
	}

	return val
}

func cat(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func echo(args []string) {
	output := os.Stdout

	if len(args) >= 3 && args[len(args)-2] == ">" {
		output, _ = os.OpenFile(args[len(args)-1], os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		defer func(output *os.File) {
			err := output.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(output)
		args = args[:len(args)-2]
	}

	for _, arg := range args {
		if arg[0] == '$' {
			if arg[1:] == "PWD" {
				cwd, _ := os.Getwd()
				_, _ = fmt.Fprintln(output, cwd)
				return
			}
			_, _ = fmt.Fprint(output, strconv.Itoa(ctx[arg[1:]]))
			_, _ = fmt.Fprint(output, " ")
		} else {
			_, _ = fmt.Fprint(output, arg)
			_, _ = fmt.Fprint(output, " ")
		}
	}

	_, _ = fmt.Fprintln(output)
}

func basename(filename string) string {
	slash := strings.LastIndex(filename, "/")
	return filename[slash+1:]
}

func cd(filename string) {
	pwd = filename
	err := os.Chdir(pwd)
	if err != nil {
		log.Fatal(err)
	}

	os.Setenv("GPWD", pwd)
}

func runCmd(cmd string) {
	for _, subcmd := range strings.Split(cmd, "&&") {
		run(strings.TrimSpace(subcmd))
	}
}

func run(cmd string) {
	sl := strings.Split(cmd, " ")

	switch sl[0] {
	case "echo":
		echo(sl[1:])
	case "exit":
		if len(sl) > 1 {
			code, _ := strconv.Atoi(sl[1])
			os.Exit(code)
		}

		os.Exit(0)
	case "var":
		ctx[sl[1]], _ = strconv.Atoi(sl[2])
	case "cd":
		cd(sl[1])
	case "sum":
		switch len(sl) {
		case 3:
			fmt.Println(parseToken(sl[1]) + parseToken(sl[2]))
		case 4:
			ctx[sl[3]] = parseToken(sl[1]) + parseToken(sl[2])
		default:
			log.Fatal("Malformed 'sum' command")
		}
	default:
		path := os.Getenv("GPATH")
		success := false

	OUTER:
		for _, dir := range strings.Split(path, ";") {
			files, err := os.OpenFile(dir, os.O_RDONLY, os.ModeDir)
			defer func(files *os.File) {
				_ = files.Close()
			}(files)

			if err != nil {
				log.Fatal(err)
			}

			ffiles, err := files.Readdirnames(0)
			if err != nil {
				log.Fatal(err)
			}

			for _, file := range ffiles {
				if basename(file) == sl[0] {
					file = dir + "/" + file
					pid, err := syscall.ForkExec(file, append([]string{file}, sl[1:]...), &syscall.ProcAttr{
						Dir:   dir,
						Env:   os.Environ(),
						Files: []uintptr{0, 1, 2},
					})
					if err != nil {
						fmt.Println(err)
					}

					_, err = syscall.Wait4(pid, nil, 0, nil)
					if err != nil {
						return
					}
					success = true
					break OUTER
				}
			}
		}

		if !success {
			log.Fatalf("Unknown command: %s", sl[0])
		}
	}
}

func gsh() {
	if len(os.Args) > 1 && os.Args[1] == "-c" {
		runCmd(strings.Join(os.Args[2:], " "))
		return
	}

	os.Setenv("GPWD", pwd)

	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	_, _ = fmt.Fprintf(writer, "%s $ ", pwd)
	_ = writer.Flush()

	for scanner.Scan() {
		line := scanner.Text()
		runCmd(line)
		_, _ = fmt.Fprintf(writer, "%s $ ", pwd)
		_ = writer.Flush()
	}
}

func ls() {
	pwd := os.Getenv("GPWD")
	files, err := os.OpenFile(pwd, os.O_RDONLY, os.ModeDir)
	defer func(files *os.File) {
		_ = files.Close()
	}(files)

	if err != nil {
		log.Fatal(err)
	}

	ffiles, err := files.Readdirnames(0)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range ffiles {
		fmt.Println(file)
	}
}

func touch() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: touch <file>")
	}

	file := os.Args[1]
	f, err := os.Create(file)
	defer func(files *os.File) {
		_ = files.Close()
	}(f)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	pwd, _ = os.Getwd()
	log.SetFlags(0)
	log.SetPrefix("GSH Runtime >> ")

	ctx = make(map[string]int)

	cmd := basename(os.Args[0])
	switch cmd {
	case "ls":
		if os.Args[0] != "/usr/bin/ls" {
			log.Fatal("Invalid placement")
		}
		ls()
	case "touch":
		if os.Args[0] != "/bin/touch" {
			log.Fatal("Invalid placement")
		}
		touch()
	case "gsh":
		if os.Args[0] != "/usr/bin/gsh" {
			log.Fatal("Invalid placement")
		}
		gsh()
	case "cat":
		if os.Args[0] != "/usr/local/bin/cat" {
			log.Fatal("Invalid placement")
		}
		cat(os.Args[1])
	default:
		fmt.Println("GSH Runtime >>")
		fmt.Printf("Unknown command: %s\n", cmd)
		fmt.Println("<< GSH Runtime")
		os.Exit(3)
	}
}
