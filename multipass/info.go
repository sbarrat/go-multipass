package multipass

import (
	"errors"
	"os/exec"
	"strings"

	"github.com/oleiade/reflections"
)

type InfoRequest struct {
	Name string
}

func Info(req *InfoRequest) (*Instance, error) {
	args := []string{"info"}
	args = append(args, req.Name)

	result := exec.Command("multipass", args...)
	out, err := result.CombinedOutput()
	if err != nil {
		return nil, errors.New(string(out) + " " + err.Error())
	}

	return parseInfo(string(out)), nil
}

var fields = map[string]string{
	"Name":      "Name:",
	"State":     "State:",
	"Snapshots": "Snapshots:",
	"IPv4":      "IPv4:",
	"Release":   "Release:",
	"ImageHash": "Image hash:",
	"CPUs":      "CPU(s):",
	"Load":      "Load:",
	"DiskUsage": "Disk usage:",
	"Mounts":    "Mounts:",
}

func parseInfo(out string) *Instance {
	var instance Instance
	for line := range strings.SplitSeq(out, "\n") {
		for key, value := range fields {
			if strings.Contains(line, value) && !strings.HasSuffix(line, "--") {
				reflections.SetField(&instance, key, strings.TrimSpace(strings.ReplaceAll(line, value, "")))
				continue
			}
		}
	}

	return &instance
}
