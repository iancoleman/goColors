package main

import "fmt"

func main() {
    fmt.Println(green("hi") + " " + yellow("test") + " " + red("bye") + " default")
}

func green(text string) string {
    return color(text, 32)
}

func red(text string) string {
    return color(text, 31)
}

func yellow(text string) string {
    return color(text, 33)
}

func color(text string, color int) string {
    return fmt.Sprintf("\033[%dm%v\033[0m", color, text)
}

