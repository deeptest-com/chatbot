package _file

import (
	"bytes"
	"errors"
	"github.com/deeptest-com/deeptest-next/internal/pkg/consts"
	_consts "github.com/deeptest-com/deeptest-next/pkg/libs/consts"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

// GetExecDir 当前执行目录
func GetExecDir() (ret string) {
	exePath, err := os.Executable()
	if err != nil {
		return
	}

	if strings.Index(strings.ToLower(exePath), "goland") > -1 { // idea debug
		ret, _ = os.Getwd()
	} else {
		ret = filepath.Dir(exePath)
	}

	ret = AddSepIfNeeded(ret)

	return
}

// GetWorkDir 当前工作目录
func GetWorkDir() (dir string) {
	if consts.WorkDir != "" {
		return consts.WorkDir
	}

	home, _ := GetUserHome()

	dir = filepath.Join(home, consts.App)
	dir = AddSepIfNeeded(dir)

	InsureDir(dir)

	consts.WorkDir = dir

	return
}

func GetUserHome() (dir string, err error) {
	user, err := user.Current()
	if err == nil {
		dir = user.HomeDir

	} else { // cross compile support

		if "windows" == runtime.GOOS { // windows
			dir, err = homeWindows()
		} else { // Unix-like system, so just assume Unix
			dir, err = homeUnix()
		}
	}

	dir = AddSepIfNeeded(dir)

	return
}
func homeUnix() (string, error) {
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// If failed, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path

	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}

	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}

func AddSepIfNeeded(pth string) string {
	if strings.LastIndex(pth, _consts.FilePthSep) < len(pth)-1 {
		pth += _consts.FilePthSep
	}
	return pth
}
