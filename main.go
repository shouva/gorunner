package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"

	helper "github.com/shouva/dailyhelper"
)

var setting Setting

func main() {

	currentdir := helper.GetCurrentPath(false)

	err := helper.ReadConfig(currentdir+"/gorunner.json", &setting)
	if err != nil {
		setting.Path = "/$HOME/go/src/github.com/username/repository"
		setting.Delay = 10 //second
		setting.Branch = "master"
		setting.Output = "output"
		file, _ := json.MarshalIndent(setting, "", " ")

		_ = ioutil.WriteFile("gorunner.json", file, 0644)
		fmt.Println("file gorunner.json tidak ditemukan.")
		fmt.Println("kami telah membuatkan 1 untuk anda. Silahkan edit, lalu jalankan kembali.")

		os.Exit(1)
	}
	fmt.Println(currentdir)
	runProcess(currentdir)
	for {
		runUpdate(currentdir)
		time.Sleep(time.Duration(setting.Delay * 1000000000))
	}
}

func runUpdate(currentdir string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover :", r)
		}
	}()
	var cmd *exec.Cmd
	if len(setting.Branch) > 0 {
		cmd = exec.Command("git", "pull", "origin", setting.Branch)
	} else {
		cmd = exec.Command("git", "pull")
	}
	cmd.Dir = setting.Path
	b, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(string(b))
	if !strings.Contains(string(b), "Already") {
		runProcess(currentdir)
	}
}

func runProcess(currentdir string) {
	if setting.Output == "" {
		setting.Output = "output"
	}
	output := currentdir + "/" + setting.Output
	fmt.Println("kill the proses!")
	exec.Command("killall", setting.Output).Start()

	// compile
	cmd := exec.Command("go", "build", "-o", output, setting.Path)
	cmd.Dir = setting.Path
	fmt.Println("Build the source!")
	fmt.Println(cmd.Args)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("proses build success!")
	cmd = exec.Command(output)
	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
}

// Setting :
type Setting struct {
	Path   string `json:"path"`
	Delay  int    `json:"delay"`
	Branch string `json:"branch"`
	Output string `json:"output"`
}
