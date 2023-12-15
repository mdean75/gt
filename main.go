package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/huh"
)

func main() {
	var flag string
	var path string
	var testRegex string
	var runCmd bool
	var writeToFile string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select gotests action").
				Options(
					huh.NewOption("all", "-all"),
					huh.NewOption("only", "-only"),
					huh.NewOption("excluded", "-excl"),
					huh.NewOption("exported", "-exported"),
				).Value(&flag),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("Regex pattern match for tests to generate").
				Value(&testRegex),
			huh.NewInput().
				Title("Please specify a file or directory containing the source").
				Value(&path),
		).WithHideFunc(func() bool {
			return flag == "-all" || flag == "-exported"
		}),
		huh.NewGroup(
			huh.NewSelect[string]().Title("Where to output tests").
				Options(
					huh.NewOption("stdout", ""),
					huh.NewOption("write to file", "-w"),
				).Value(&writeToFile),
		),
	)

	form.Run()

	huh.NewConfirm().
		Title(fmt.Sprintf("issue command: `gotests %s %s %s %s`", flag, writeToFile, testRegex, path)).
		Affirmative("yes").
		Negative("no").Value(&runCmd).Run()

	if runCmd {
		var cmdx strings.Builder

		cmdx.WriteString("gotests")
		cmdx.WriteRune(' ')
		cmdx.WriteString(flag)
		if testRegex != "" {
			cmdx.WriteRune(' ')
			cmdx.WriteString(testRegex)
		}

		if path != "" {
			cmdx.WriteRune(' ')
			cmdx.WriteString(path)
		}

		cmdArgs := strings.Split(cmdx.String(), " ")
		fmt.Println("runnning command:", cmdArgs)
		out, err := exec.Command(cmdArgs[0], cmdArgs[1:]...).CombinedOutput()
		if err != nil {
			fmt.Println(err, string(out))
			os.Exit(1)
		}
		fmt.Println(string(out))
	}
}

func Add(a, b int) int {
	return a + b
}
