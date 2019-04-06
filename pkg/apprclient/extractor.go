package apprclient

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

type extractor struct {
}

func (e *extractor) Extract(encoded []byte) ([]byte, error) {
	gzipReader, err := gzip.NewReader(bytes.NewBuffer(encoded))
	if err != nil {
		return nil, err
	}
	defer gzipReader.Close()

	decoded, err := e.extract(gzipReader)
	if err != nil {
		return nil, err
	}

	return decoded, nil
}

func (e *extractor) extract(r io.Reader) ([]byte, error) {
	reader := tar.NewReader(r)

	writer := &bytes.Buffer{}
	for true {
		header, err := reader.Next()
		if err != nil && err != io.EOF {
			return nil, errors.New(fmt.Sprintf("extraction of tar ball failed - %s", err.Error()))
		}

		if err == io.EOF {
			break
		}

		switch header.Typeflag {
		case tar.TypeReg:
			// io.Copy(writer, reader)
			fmt.Printf("%s\n", header.Name)
			continue

		case tar.TypeDir:
			fmt.Printf("%s\n", header.Name)
			continue
		}
	}

	return writer.Bytes(), nil
}

func Decode(encoded []byte) ([]byte, error) {
	maxlength := base64.StdEncoding.DecodedLen(len(encoded))
	decoded := make([]byte, maxlength)

	n, err := base64.StdEncoding.Decode(decoded, encoded)
	if err != nil {
		return nil, err
	}

	return decoded[:n], nil
}
