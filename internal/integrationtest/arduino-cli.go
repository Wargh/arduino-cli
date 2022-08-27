// This file is part of arduino-cli.
//
// Copyright 2022 ARDUINO SA (http://www.arduino.cc/)
//
// This software is released under the GNU General Public License version 3,
// which covers the main part of arduino-cli.
// The terms of this license can be found at:
// https://www.gnu.org/licenses/gpl-3.0.en.html
//
// You can be released from the requirements of the above licenses by purchasing
// a commercial license. Buying such a license is mandatory if you want to
// modify or otherwise use the software for commercial activities involving the
// Arduino software without disclosing the source code of your own applications.
// To purchase a commercial license, send an email to license@arduino.cc.

package integrationtest

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/arduino/arduino-cli/executils"
	"github.com/arduino/arduino-cli/rpc/cc/arduino/cli/commands/v1"
	"github.com/arduino/go-paths-helper"
	"github.com/fatih/color"
	"github.com/stretchr/testify/require"
	"go.bug.st/testsuite"
	"google.golang.org/grpc"
)

// FindRepositoryRootPath returns the repository root path
func FindRepositoryRootPath(t *testing.T) *paths.Path {
	repoRootPath := paths.New(".")
	require.NoError(t, repoRootPath.ToAbs())
	for !repoRootPath.Join(".git").Exist() {
		require.Contains(t, repoRootPath.String(), "arduino-cli", "Error searching for repository root path")
		repoRootPath = repoRootPath.Parent()
	}
	return repoRootPath
}

// CreateArduinoCLIWithEnvironment performs the minimum amount of actions
// to build the default test environment.
func CreateArduinoCLIWithEnvironment(t *testing.T) (*testsuite.Environment, *ArduinoCLI) {
	env := testsuite.NewEnvironment(t)

	cli := NewArduinoCliWithinEnvironment(env, &ArduinoCLIConfig{
		ArduinoCLIPath:         FindRepositoryRootPath(t).Join("arduino-cli"),
		UseSharedStagingFolder: true,
	})

	return env, cli
}

// ArduinoCLI is an Arduino CLI client.
type ArduinoCLI struct {
	path          *paths.Path
	t             *require.Assertions
	proc          *executils.Process
	cliEnvVars    []string
	cliConfigPath *paths.Path
	stagingDir    *paths.Path
	dataDir       *paths.Path
	sketchbookDir *paths.Path
	daemonAddr    string
	daemonConn    *grpc.ClientConn
	daemonClient  commands.ArduinoCoreServiceClient
}

// ArduinoCLIConfig is the configuration of the ArduinoCLI client
type ArduinoCLIConfig struct {
	ArduinoCLIPath         *paths.Path
	UseSharedStagingFolder bool
}

// NewArduinoCliWithinEnvironment creates a new Arduino CLI client inside the given environment.
func NewArduinoCliWithinEnvironment(env *testsuite.Environment, config *ArduinoCLIConfig) *ArduinoCLI {
	color.NoColor = false
	cli := &ArduinoCLI{
		path:          config.ArduinoCLIPath,
		t:             require.New(env.T()),
		dataDir:       env.RootDir().Join("Arduino15"),
		sketchbookDir: env.RootDir().Join("Arduino"),
		stagingDir:    env.RootDir().Join("Arduino15/staging"),
	}
	if config.UseSharedStagingFolder {
		cli.stagingDir = env.SharedDownloadsDir()
	}

	cli.cliEnvVars = []string{
		fmt.Sprintf("ARDUINO_DATA_DIR=%s", cli.dataDir),
		fmt.Sprintf("ARDUINO_DOWNLOADS_DIR=%s", cli.stagingDir),
		fmt.Sprintf("ARDUINO_SKETCHBOOK_DIR=%s", cli.sketchbookDir),
	}
	env.RegisterCleanUpCallback(cli.CleanUp)
	return cli
}

// CleanUp closes the Arduino CLI client.
func (cli *ArduinoCLI) CleanUp() {
	if cli.proc != nil {
		cli.daemonConn.Close()
		cli.proc.Kill()
		cli.proc.Wait()
	}
}

// DataDir returns the data directory
func (cli *ArduinoCLI) DataDir() *paths.Path {
	return cli.dataDir
}

// SketchbookDir returns the sketchbook directory
func (cli *ArduinoCLI) SketchbookDir() *paths.Path {
	return cli.sketchbookDir
}

// Run executes the given arduino-cli command and returns the output.
func (cli *ArduinoCLI) Run(args ...string) ([]byte, []byte, error) {
	if cli.cliConfigPath != nil {
		args = append([]string{"--config-file", cli.cliConfigPath.String()}, args...)
	}
	fmt.Println(color.HiBlackString(">>> Running: ") + color.HiYellowString("%s %s", cli.path, strings.Join(args, " ")))
	cliProc, err := executils.NewProcessFromPath(cli.cliEnvVars, cli.path, args...)
	cli.t.NoError(err)
	stdout, err := cliProc.StdoutPipe()
	cli.t.NoError(err)
	stderr, err := cliProc.StderrPipe()
	cli.t.NoError(err)
	_, err = cliProc.StdinPipe()
	cli.t.NoError(err)

	cli.t.NoError(cliProc.Start())

	var stdoutBuf, stderrBuf bytes.Buffer
	stdoutCtx, stdoutCancel := context.WithCancel(context.Background())
	stderrCtx, stderrCancel := context.WithCancel(context.Background())
	go func() {
		io.Copy(&stdoutBuf, io.TeeReader(stdout, os.Stdout))
		stdoutCancel()
	}()
	go func() {
		io.Copy(&stderrBuf, io.TeeReader(stderr, os.Stderr))
		stderrCancel()
	}()
	cliErr := cliProc.Wait()
	<-stdoutCtx.Done()
	<-stderrCtx.Done()
	fmt.Println(color.HiBlackString("<<< Run completed (err = %v)", cliErr))

	return stdoutBuf.Bytes(), stderrBuf.Bytes(), cliErr
}

// StartDaemon starts the Arduino CLI daemon. It returns the address of the daemon.
func (cli *ArduinoCLI) StartDaemon(verbose bool) string {
	args := []string{"daemon", "--format", "json"}
	if cli.cliConfigPath != nil {
		args = append([]string{"--config-file", cli.cliConfigPath.String()}, args...)
	}
	if verbose {
		args = append(args, "-v", "--log-level", "debug")
	}
	cliProc, err := executils.NewProcessFromPath(cli.cliEnvVars, cli.path, args...)
	cli.t.NoError(err)
	stdout, err := cliProc.StdoutPipe()
	cli.t.NoError(err)
	stderr, err := cliProc.StderrPipe()
	cli.t.NoError(err)
	_, err = cliProc.StdinPipe()
	cli.t.NoError(err)

	cli.t.NoError(cliProc.Start())
	cli.proc = cliProc
	cli.daemonAddr = "127.0.0.1:50051"

	copy := func(dst io.Writer, src io.Reader) {
		buff := make([]byte, 1024)
		for {
			n, err := src.Read(buff)
			if err != nil {
				return
			}
			dst.Write([]byte(color.YellowString(string(buff[:n]))))
		}
	}
	go copy(os.Stdout, stdout)
	go copy(os.Stderr, stderr)
	conn, err := grpc.Dial(cli.daemonAddr, grpc.WithInsecure(), grpc.WithBlock())
	cli.t.NoError(err)
	cli.daemonConn = conn
	cli.daemonClient = commands.NewArduinoCoreServiceClient(conn)

	return cli.daemonAddr
}

// ArduinoCLIInstance is an Arduino CLI gRPC instance.
type ArduinoCLIInstance struct {
	cli      *ArduinoCLI
	instance *commands.Instance
}

var logCallfMutex sync.Mutex

func logCallf(format string, a ...interface{}) {
	logCallfMutex.Lock()
	fmt.Print(color.HiRedString(format, a...))
	logCallfMutex.Unlock()
}

// Create calls the "Create" gRPC method.
func (cli *ArduinoCLI) Create() *ArduinoCLIInstance {
	logCallf(">>> Create()")
	resp, err := cli.daemonClient.Create(context.Background(), &commands.CreateRequest{})
	cli.t.NoError(err)
	logCallf(" -> %v\n", resp)
	return &ArduinoCLIInstance{
		cli:      cli,
		instance: resp.Instance,
	}
}

// Init calls the "Init" gRPC method.
func (inst *ArduinoCLIInstance) Init(profile string, sketchPath string, respCB func(*commands.InitResponse)) error {
	initReq := &commands.InitRequest{
		Instance:   inst.instance,
		Profile:    profile,
		SketchPath: sketchPath,
	}
	logCallf(">>> Init(%v)\n", initReq)
	initClient, err := inst.cli.daemonClient.Init(context.Background(), initReq)
	if err != nil {
		return err
	}
	for {
		msg, err := initClient.Recv()
		if err == io.EOF {
			logCallf("<<< Init EOF\n")
			return nil
		}
		if err != nil {
			return err
		}
		if respCB != nil {
			respCB(msg)
		}
	}
}

// BoardList calls the "BoardList" gRPC method.
func (inst *ArduinoCLIInstance) BoardList(timeout time.Duration) (*commands.BoardListResponse, error) {
	boardListReq := &commands.BoardListRequest{
		Instance: inst.instance,
		Timeout:  timeout.Milliseconds(),
	}
	logCallf(">>> BoardList(%v) -> ", boardListReq)
	resp, err := inst.cli.daemonClient.BoardList(context.Background(), boardListReq)
	logCallf("err=%v\n", err)
	return resp, err
}

// BoardListWatch calls the "BoardListWatch" gRPC method.
func (inst *ArduinoCLIInstance) BoardListWatch() (commands.ArduinoCoreService_BoardListWatchClient, error) {
	boardListWatchReq := &commands.BoardListWatchRequest{
		Instance: inst.instance,
	}
	logCallf(">>> BoardListWatch(%v)\n", boardListWatchReq)
	watcher, err := inst.cli.daemonClient.BoardListWatch(context.Background())
	if err != nil {
		return watcher, err
	}
	return watcher, watcher.Send(boardListWatchReq)
}

// PlatformInstall calls the "PlatformInstall" gRPC method.
func (inst *ArduinoCLIInstance) PlatformInstall(ctx context.Context, packager, arch, version string, skipPostInst bool) (commands.ArduinoCoreService_PlatformInstallClient, error) {
	installCl, err := inst.cli.daemonClient.PlatformInstall(ctx, &commands.PlatformInstallRequest{
		Instance:        inst.instance,
		PlatformPackage: packager,
		Architecture:    arch,
		Version:         version,
		SkipPostInstall: skipPostInst,
	})
	logCallf(">>> PlatformInstall(%v:%v %v)\n", packager, arch, version)
	return installCl, err
}

// Compile calls the "Compile" gRPC method.
func (inst *ArduinoCLIInstance) Compile(ctx context.Context, fqbn, sketchPath string) (commands.ArduinoCoreService_CompileClient, error) {
	compileCl, err := inst.cli.daemonClient.Compile(ctx, &commands.CompileRequest{
		Instance:   inst.instance,
		Fqbn:       fqbn,
		SketchPath: sketchPath,
		Verbose:    true,
	})
	logCallf(">>> Compile(%v %v)\n", fqbn, sketchPath)
	return compileCl, err
}