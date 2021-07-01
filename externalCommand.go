package osrunner
import (
  "bytes"
  "fmt"
  "os"
  "os/exec"
  "sync"
)

type ExternalCommand struct {
  associatedGroup *sync.WaitGroup
  osCmd *exec.Cmd
  outputBuffer bytes.Buffer
  outputErrBuffer bytes.Buffer
}

func (ec *ExternalCommand) AssociateWith(collection *sync.WaitGroup) *ExternalCommand {
  ec.associatedGroup = collection

  return ec
}

func (ec *ExternalCommand) Do(withDebug bool) {
  // attempt to execute specified command in the os
  osCmdErr := ec.osCmd.Run()
  //_, isExitError := osCmdErr.(*exec.ExitError)

  if ec.associatedGroup != nil {
    defer ec.associatedGroup.Done()
  }

  if withDebug {
    ec.showCommandExecutionData()
  }

  // determine if we've encountered a command failure
  // this program fully depends on successful execution of these commands...
  //if ec.outputErrBuffer.Len() > 0 && osCmdErr != nil && isExitError {
  if osCmdErr != nil {
    // if we've seen a problem, let's not continue
    fmt.Printf("failed to execute: %s (enable debugging to see more)\n", ec.osCmd.String())
    fmt.Printf("osCmdErr> %+v\n", osCmdErr)
    fmt.Printf("ec> %+v\n", ec)
    os.Exit(1)
  }
}

func (ec *ExternalCommand) GetError() string {
  return ec.outputErrBuffer.String()
}

func (ec *ExternalCommand) GetOutput() string {
  return ec.outputBuffer.String()
}

func (ec *ExternalCommand) showCommandExecutionData() {
  fmt.Printf(`
  ----------
  |  command: %s
  |  bytes returned: length, capacity
  |  =========================
  |  out: %d, %d
  |  err: %d, %d
  ----------
  out >>
%s
  err >>
%s
  `,
    ec.osCmd.String(),
    ec.outputBuffer.Len(), ec.outputBuffer.Cap(),
    ec.outputErrBuffer.Len(), ec.outputErrBuffer.Cap(),
    ec.outputBuffer.String(),
    ec.outputErrBuffer.String(),
  )
  fmt.Println()
}
