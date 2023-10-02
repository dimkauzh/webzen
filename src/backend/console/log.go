package console

import (
  "fmt"
  "syscall/js"
)

func Log(args ...interface{}) {
  logArgs := make([]interface{}, len(args))
  for i, arg := range args {
    // Convert arrays to strings before logging
    if arr, ok := arg.([]interface{}); ok {
      logArgs[i] = fmt.Sprintf("%v", arr)
    } else {
      logArgs[i] = arg
    }
  }

  js.Global().Get("console").Call("log", logArgs...)
}
