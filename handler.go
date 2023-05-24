package lab2

import (
  "io"
  "strings"
)

type ComputeHandler struct {
  Input  io.Reader
  Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
  bufRead, err := io.ReadAll(ch.Input)
  if err != nil {
    return err
  }
  expression := string(bufRead)
  trimmed := strings.Trim(expression, " \n")
  res, err := PrefixToPostfix(trimmed)
  if err != nil {
    return err
  }
  ch.Output.Write([]byte(res))
  return nil
}