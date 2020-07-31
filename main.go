package main

import (
    "os"
    "lunarOS-imagebuilder/build"
)

func main() {
    codeType := os.Getenv("TYPE")
    build.RunOnType(codeType)
}
