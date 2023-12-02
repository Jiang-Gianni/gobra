package scan

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"sort"
)

var (
	ErrExists    = errors.New("host already in the list")
	ErrNotExists = errors.New("host not in the list")
)

// HostsList represents a list of hosts to run port scan
type HostsList struct {
	Hosts []string
}

func (hl *HostsList) Search(host string) (bool, int) {
	slices.Sort(hl.Hosts)

	i := sort.SearchStrings(hl.Hosts, host)
	if i < len(hl.Hosts) && hl.Hosts[i] == host {
		return true, i
	}

	return false, -1
}

func (hl *HostsList) Add(host string) error {
	if found, _ := hl.Search(host); found {
		return ErrExists
	}

	hl.Hosts = append(hl.Hosts, host)

	return nil
}

func (hl *HostsList) Remove(host string) error {
	if found, i := hl.Search(host); found {
		hl.Hosts = append(hl.Hosts[:i], hl.Hosts[i+1:]...)
		return nil
	}

	return fmt.Errorf("%w:%s", ErrNotExists, host)
}

func (hl *HostsList) Load(hostsFile string) (err error) {
	f, err := os.Open(hostsFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return err
	}

	defer func() { err = f.Close() }()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		hl.Hosts = append(hl.Hosts, scanner.Text())
	}

	return nil
}

func (hl *HostsList) Save(hostsFile string) error {
	output := ""
	for _, h := range hl.Hosts {
		output += fmt.Sprintln(h)
	}

	return os.WriteFile(hostsFile, []byte(output), 0o644)
}
