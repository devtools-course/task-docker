# Mini alpine

You have to create a miniature "alpine-like" image from the [main.go](./main.go) file

> To build a standalone binary from Go sources, use `CGO_ENABLED=0 GOOS=linux go build -o OUT FROM`

The resulting binary implements these utilities:
```go
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
```

Where `gsh` is a custom shell. It uses `GPATH` instead of `PATH` with a semicolon as the separator.

`docker run --rm -it gpine:latest` must open a `gsh` session.

**Important!** The image size must be less than 6.5 MB. 

---
* What is `scratch`?
* How can the file duplication be reduced?
