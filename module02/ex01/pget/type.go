package pget

import (
	"os"
)

type openfile interface {
	open(name string) (*os.File, error)
}

type dstfile struct {
	name  string
	exist bool
}

func (d *dstfile) open(name string) (*os.File, error) {
	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	d.name = file.Name()
	d.exist = true
	return file, nil
}

type tmpfile struct {
	name  string
	exist bool
}

func (t *tmpfile) open(pattern string) (*os.File, error) {
	file, err := os.CreateTemp("", pattern)
	if err != nil {
		return nil, err
	}
	t.name = file.Name()
	t.exist = true
	return file, nil
}

func (t *tmpfile) remove() {
	if t.exist {
		os.Remove(t.name)
	}
}
