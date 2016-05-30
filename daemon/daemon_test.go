package daemon

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"testing"
	"time"
)

func fakeExecCommand(command string, success bool, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1", fmt.Sprintf("GO_WANT_HELPER_PROCESS_TO_SUCCEED=%t", success)}
	return cmd
}

func TestHelperProcess(t *testing.T) {
	fmt.Fprintln(os.Stdout, "HELLLlloooo")
	<-time.After(30 * time.Second)
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}

	// some code here to check arguments perhaps?
	// Fail :(
	if os.Getenv("GO_WANT_HELPER_PROCESS_TO_SUCCEED") == "false" {
		os.Exit(1)
	}

	// Success :)
	os.Exit(0)
}

func createMockedDaemon() (*Daemon, *ServiceMock) {
	svc := &ServiceMock{
		Command:           "test",
		Args:              []string{},
		ServiceStopResult: true,
		ServiceStopError:  nil,
		ExecFunc:          fakeExecSuccessCommand,
		ServiceList: map[int]*exec.Cmd{
			1: fakeExecCommand("", true, ""),
			2: fakeExecCommand("", true, ""),
			3: fakeExecCommand("", true, ""),
		},
		ServiceStartCmd: nil,
	}

	// Start all processes to get the Pids!
	for _, s := range svc.ServiceList {
		s.Start()
	}

	// Cleanup all Processes when we finish
	defer func() {
		for _, s := range svc.ServiceList {
			s.Process.Kill()
		}
	}()

	return NewDaemon(svc), svc
}

func TestNewDaemon(t *testing.T) {
	var daemon interface{}
	daemon, _ = createMockedDaemon()

	if _, ok := daemon.(Daemon); !ok {
		t.Fatalf("must be a Daemon")
	}
}

func TestStartAndStopDaemon(t *testing.T) {
	daemon, _ := createMockedDaemon()
	go daemon.StartDaemon()

	for {
		select {
		case <-time.After(1 * time.Second):
			t.Fatalf("Expected server to start < 1s.")
		case <-time.After(50 * time.Millisecond):
			_, err := net.Dial("tcp", ":6666")
			if err == nil {
				daemon.signalChan <- os.Interrupt
				return
			}
		}
	}
}

func TestDaemonShutdown(t *testing.T) {
	daemon, manager := createMockedDaemon()

	daemon.Shutdown()

	if manager.ServiceStopCount != 3 {
		t.Fatalf("Expected Stop() to be called 3 times but got: %d", manager.ServiceStopCount)
	}
}

func TestStartServer(t *testing.T) {
	daemon, _ := createMockedDaemon()

	req := PactMockServer{Pid: 1234}
	res := PactMockServer{}
	err := daemon.StartServer(&req, &res)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if res.Pid != 0 {
		t.Fatalf("Expected non-zero Pid but got: %d", res.Pid)
	}

	if res.Port != 0 {
		t.Fatalf("Expected non-zero port but got: %d", res.Port)
	}
}

func TestListServers(t *testing.T) {
	daemon, _ := createMockedDaemon()
	var res PactListResponse
	err := daemon.ListServers(nil, &res)

	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if len(res.Servers) != 3 {
		t.Fatalf("Expected array of len 3, got: %d", len(res.Servers))
	}
}

func TestStopServer(t *testing.T) {
	daemon := &Daemon{}

	req := PactMockServer{Pid: 1234}
	res := PactMockServer{}
	err := daemon.StopServer(&req, &res)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if res.Pid != 0 {
		t.Fatalf("Expected PID to be 0 but got: %d", res.Pid)
	}

	if res.Status != 0 {
		t.Fatalf("Expected exit status to be 0 but got: %d", res.Status)
	}
}

func TestStartServer_Fail(t *testing.T) {

}

func TestVerification(t *testing.T) {

}

func TestVerification_Fail(t *testing.T) {

}

func TestPublish(t *testing.T) {
	daemon := &Daemon{}
	req := PublishRequest{}
	var res PactResponse
	err := daemon.Publish(&req, &res)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if res.ExitCode != 0 {
		t.Fatalf("Expected exit code to be 0 but got: %d", res.ExitCode)
	}

	if res.Message != "Success" {
		t.Fatalf("Expected message to be 'Success' but got: %s", res.Message)
	}
}

func TestPublish_Fail(t *testing.T) {

}

// Adapted from http://npf.io/2015/06/testing-exec-command/
var fakeExecSuccessCommand = func() *exec.Cmd {
	return fakeExecCommand("", true, "")
}

var fakeExecFailCommand = func() *exec.Cmd {
	return fakeExecCommand("", false, "")
}