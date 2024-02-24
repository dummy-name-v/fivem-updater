package fsio

import (
	"fmt"
	"log"
	"runtime"
)

type Platform string

const (
	windowsPlatform Platform = "windows"
	linuxPlatform   Platform = "linux"
)

type platformStruct struct {
	Windows Platform
	Linux   Platform
}

func (p *platformStruct) Validate(value string) (*Platform, error) {
	if value != string(windowsPlatform) && value != string(linuxPlatform) {
		return nil, fmt.Errorf("unsupported platform %s", value)
	}

	platform := Platform(value)
	return &platform, nil
}

func ParseArguments(args []string) (*Platform, string, error) {
	platform := runtime.GOOS
	out := "binaries"

	l := len(args)
	for i := 0; i < l; i++ {
		switch args[i] {
		case "-o":
			{
				if i+1 < l {
					log.Fatal("NOP")
				}

				out = args[i+1]
				break
			}

		case "-linux", "-windows":
			{
				platform = args[i][1:]
				break
			}
		}
	}

	a, err := Platforms.Validate(platform)
	return a, out, err
}

var Platforms = platformStruct{
	Windows: windowsPlatform,
	Linux:   linuxPlatform,
}
