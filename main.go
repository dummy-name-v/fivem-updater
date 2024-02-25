package main

import (
	"fivem-updater/fsio"
	"fivem-updater/github"
	"log"
	"os"
)

const REPO = "citizenfx/fivem"

func main() {
	platform, out, err := fsio.ParseArguments(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := fsio.GetConfig("updater.json")
	if err != nil {
		log.Fatal(err)
	}

	tag, err := github.GetLatestRepositoryTag(REPO)
	if err != nil {
		log.Fatal(err)
	}

	hash := fsio.FormatConfigHash(platform, tag)
	if cfg.Hash == hash {
		log.Println("Up to date, skipping download")
		return
	}

	url, archiveName := fsio.GetFileAssociation(platform, tag)
	err = fsio.DownloadFile(url, archiveName)
	if err != nil {
		log.Fatal(err)
	}

	err = fsio.UnzipArchive(platform, archiveName, out)
	if err != nil {
		log.Fatal(err)
	}

	if err := cfg.UpdateHash(hash); err != nil {
		log.Fatal(err)
	}

	log.Println("> Update finished")
}
