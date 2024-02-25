package fsio

import (
    "fivem-updater/github"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "os/exec"
    "path/filepath"
)

const BASE_URL = "https://runtime.fivem.net/artifacts/fivem"

var platformMapping = map[Platform][2]string{
    Platforms.Windows: {"build_server_windows", "server.7z"},
    Platforms.Linux:   {"build_proot_linux", "fx.tar.xz"},
}

func GetFileAssociation(platform *Platform, tag *github.Tag) (string, string) {
    mapping, ok := platformMapping[*platform]
    if !ok {
        log.Fatalf("platform %s is not supported", platform)
    }

    return fmt.Sprintf("%s/%s/master/%s-%s/%s", BASE_URL, mapping[0], tag.Version, tag.Sha, mapping[1]), mapping[1]
}

func DownloadFile(url string, filePath string) error {
    fmt.Println("> Downloading binaries..")

    out, err := os.Create(filePath)
    if err != nil {
        return err
    }
    defer out.Close()

    response, err := http.Get(url)
    if err != nil {
        return err
    }
    defer response.Body.Close()

    if response.StatusCode != http.StatusOK {
        return fmt.Errorf("bad status: %s", response.Status)
    }

    if _, err := io.Copy(out, response.Body); err != nil {
        return err
    }

    return nil
}

func removeDirectoryContent(dir string) error {
    files, err := filepath.Glob(filepath.Join(dir, "*"))
    if err != nil {
        return err
    }

    for _, file := range files {
        if err := os.RemoveAll(file); err != nil {
            return err
        }
    }

    return nil
}

func UnzipArchive(platform *Platform, in string, out string) error {
    fmt.Println("> Unzipping binaries..")

    if err := removeDirectoryContent("./binaries"); err != nil {
        return err
    }

    var cmd *exec.Cmd
    if *platform == "windows" {
        cmd = exec.Command("7z", "x", in, fmt.Sprintf("-o%s/", out))
    } else {
        cmd = exec.Command("tar", "-xf", in, "-C", out)
    }

    if err := cmd.Run(); *platform == "window" && err != nil {
        return err
    }

    for _, mapping := range platformMapping {
        if _, err := os.Stat(mapping[1]); err != nil {
            continue
        }

        if err := os.Remove(mapping[1]); err != nil {
            return err
        }
    }

    return nil
}
