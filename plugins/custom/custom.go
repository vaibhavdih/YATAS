package custom

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/stangirard/yatas/internal/yatas"
)

func findPluginWithName(c *yatas.Config, name string) *yatas.Plugin {
	for _, plugin := range c.Plugins {
		if plugin.Name == name {
			return &plugin
		}
	}
	return nil
}

func Run(c *yatas.Config, name string) (yatas.Tests, error) {
	plugin := findPluginWithName(c, name)
	checks, err := ExecuteCommand(c, plugin)
	return checks, err

}

func ExecuteCommand(c *yatas.Config, plugin *yatas.Plugin) (yatas.Tests, error) {
	checks := []yatas.Check{}
	check := yatas.Check{}
	check.Name = plugin.Name
	check.Description = plugin.Description
	check.Status = "OK"

	cmd := exec.Command(plugin.Command, plugin.Args...)
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	result := yatas.Result{}
	if strings.TrimRight(outb.String(), "\n") == plugin.ExpectedOutput {
		result.Message = fmt.Sprint("Output matched: ", plugin.ExpectedOutput)
		result.Status = "OK"
	} else {
		result.Message = fmt.Sprint("Output did not match: ", plugin.ExpectedOutput, " instead got: ", outb.String())
		result.Status = "FAIL"
	}
	check.Results = append(check.Results, result)
	checks = append(checks, check)
	test := yatas.Tests{}
	test.Checks = checks
	test.Account = plugin.Name
	return test, nil

}
