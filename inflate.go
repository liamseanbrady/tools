package main

import (
	"io"
	"log"
	"os"
	"bytes"
	"compress/zlib"
)

func main() {
  // Grab data from STDIN, most likely via a pipe to this program
  compressedData, stdinReadErr := io.ReadAll(os.Stdin)
  if stdinReadErr != nil {
    log.Fatalf("Error reading from stdin: %v", stdinReadErr)
  }

  // The first byte of a file in zlib compressed format should be 120 in decimal, so we can use this to discount files from being in zlib compressed format. This is an exlusive check for the zlib compressed format - it doesn't necessarily mean that the given input is in zlib compressed format.
  const zlibMagicByteValue = 120
  if compressedData[0] != zlibMagicByteValue {
    log.Fatalf("This file is not in the zlib compressed file format")
  }

  // Create a reader from the compressed data we got from STDIN
  var compressedDataBuffer io.Reader = bytes.NewBuffer(compressedData)

  // Use the `zlib` package decompress the data from the zlib compression file format
  r, zlibErr := zlib.NewReader(compressedDataBuffer)
  if zlibErr != nil {
    log.Fatalf("Error decompressing data using the zlib package: %v", zlibErr)
  }
  defer r.Close()

  _, stdoutCopyErr := io.Copy(os.Stdout, r)
  if stdoutCopyErr != nil {
    log.Fatalf("Error copying data to STDOUT: %v", stdoutCopyErr)
  }
}
