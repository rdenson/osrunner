package osrunner
import (
  "os"
  "testing"

  "github.com/stretchr/testify/suite"
)

type ExternalCommandSuite struct {
  suite.Suite
}

func TestExternalCommand(t *testing.T) {
  suite.Run(t, new(ExternalCommandSuite))
}

func (suite *ExternalCommandSuite) TestIt() {
  /*genv := os.Environ()
  suite.T().Logf("found %d environment variables", len(genv))
  for _, kp := range genv {
    suite.T().Logf("%s\n", kp)
  }
  os.Exit(1)*/

  listFsContent, setCommandErr := SetCommand("~/config", "ls", "-ahl")
  if setCommandErr != nil {
    suite.T().Logf("%+v\n", setCommandErr)
    os.Exit(1)
  }
  listFsContent.Do(true)
  suite.T().Logf("stdout> %s\n", listFsContent.GetOutput())
  suite.T().Logf("stderr> %s\n", listFsContent.GetError())

  suite.Equal(1, 1)
}
