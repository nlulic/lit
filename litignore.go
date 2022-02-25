package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func (lit *Lit) Ignorefiles() ([]string, error) {

	r, err := lit.Root()

	if err != nil {
		return nil, err
	}

	f, err := os.Open(filepath.Join(r, lit.config.IgnoreFileName))

	if err != nil {
		if os.IsNotExist(err) {
			return matches(r, lit.config.RootDir), nil
		}

		return nil, err
	}

	defer f.Close()

	globs := []string{
		lit.config.RootDir, //.lit
	}

	s := bufio.NewScanner(f)

	for s.Scan() {
		line := s.Text()

		// skip comments prefixed with '#'
		if strings.HasPrefix(line, "#") {
			continue
		}

		// skip empty lines
		if strings.TrimSpace(line) == "" {
			continue
		}

		globs = append(globs, line)
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return matches(r, globs...), nil
}

// TODO: Test
func matches(base string, globs ...string) []string {
	var matches []string

	for _, glob := range globs {
		m, _ := filepath.Glob(filepath.Join(base, glob))
		matches = append(matches, m...)
	}

	return matches
}
