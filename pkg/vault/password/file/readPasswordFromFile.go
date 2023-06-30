package file

import (
	"bufio"
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

// OptionsFunc is a function used to configure ReadPasswordFromFile
type OptionsFunc func(*ReadPasswordFromFile)

// ReadPasswordFromFile allows you to read a passowrd from a file
type ReadPasswordFromFile struct {
	file string
	fs   afero.Fs
}

// NewReadPasswordFromFile returns a ReadPasswordFromFile
func NewReadPasswordFromFile(options ...OptionsFunc) *ReadPasswordFromFile {
	secret := &ReadPasswordFromFile{}
	secret.Options(options...)

	return secret
}

// WithFile sets the a file where to look for a password
func WithFile(file string) OptionsFunc {
	return func(s *ReadPasswordFromFile) {
		s.file = file
	}
}

// WithFs set the filesystem
func WithFs(fs afero.Fs) OptionsFunc {
	return func(s *ReadPasswordFromFile) {
		s.fs = fs
	}
}

// Options configure the ReadPasswordFromFile
func (s *ReadPasswordFromFile) Options(opts ...OptionsFunc) {
	for _, opt := range opts {
		opt(s)
	}
}

// Read returns a password from a file. It return an error when the file does not exist or the content of the file can not be read.
func (s *ReadPasswordFromFile) Read() (string, error) {
	var password string
	var err error
	var passwordFile afero.File

	if s == nil {
		return "", errors.New("Read password from a file component has not been initialized.")
	}

	if len(s.file) <= 0 {
		return "", errors.New("File path must be specified to read the password from a file.")
	}

	if s.fs == nil {
		s.fs = afero.NewOsFs()
	}

	_, err = s.fs.Stat(s.file)
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("Error describing the file '%s'.", s.file))
	}

	passwordFile, err = s.fs.Open(s.file)
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("Error opening the file '%s'.", s.file))
	}
	defer passwordFile.Close()

	scanner := bufio.NewScanner(passwordFile)

	// Read the first line
	scanner.Scan()
	password = scanner.Text()

	err = scanner.Err()
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("Error reading the file '%s' content.", s.file))
	}

	return password, nil
}
