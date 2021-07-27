package osrunner
import (
  "fmt"
  "os"
  "os/exec"
  "regexp"
)

func handleTildeExpansion(path string) (string, error) {
  requestedLocationHasTilde, regexMatchErr := regexp.MatchString(`^~/.*`, path)
  if regexMatchErr != nil {
    return "", regexMatchErr
  }

  if requestedLocationHasTilde {
    path = os.Getenv("HOME") + path[1:]
  }

  return path, nil
}

// for exec'ing a vanilla command:
// $> EXECUTABLE ARG1 ARG2 ...
func SetCommand(fsLocation, executableName string, execArgs ...string) (*ExternalCommand, error) {
  extCmd := new(ExternalCommand)

  workingDirectory, tildeExpansionErr := handleTildeExpansion(fsLocation)
  if tildeExpansionErr != nil {
    return nil, tildeExpansionErr
  }

  fmt.Printf("> fsLocation: %s\n", workingDirectory)
  extCmd.osCmd = exec.Command(executableName, execArgs...)
  extCmd.osCmd.Dir = workingDirectory
  extCmd.osCmd.Stdout = &extCmd.outputBuffer
  extCmd.osCmd.Stderr = &extCmd.outputErrBuffer

  return extCmd, nil
}

// for exec'ing something a bit more... ie. piped:
// $> EXECUTABLE_A ARG1 | EXECUTABLE_B ARG1 ARG2
func SetCustomCommand(fsLocation, cmdStr string) *ExternalCommand {
  extCmd := new(ExternalCommand)

  extCmd.osCmd = exec.Command("bash", "-c", cmdStr)
  extCmd.osCmd.Dir = fsLocation
  extCmd.osCmd.Stdout = &extCmd.outputBuffer
  extCmd.osCmd.Stderr = &extCmd.outputErrBuffer

  return extCmd
}
