package fsio

import (
	"fmt"
	"runtime"
	"strings"
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
				if l < i+2 || strings.HasPrefix(args[i+2], "-") {
					return nil, "", fmt.Errorf("not enough arguments passed for the output")
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

	_platform, err := Platforms.Validate(platform)
	return _platform, out, err
}

var Platforms = platformStruct{
	Windows: windowsPlatform,
	Linux:   linuxPlatform,
}
