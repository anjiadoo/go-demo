package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
var (
	ErrAssertFail           = errors.New("Type Assert Failed.")
	ErrNotFoundTasker       = errors.New("Not Found Tasker")
	ErrNotFoundTaskItem     = errors.New("Not Found Task Item")
	ErrNotDoneTaskItem      = errors.New("Not Done Task Item")
	ErrGrantTypeInvalid     = errors.New("Award Grant Type Invalid")
	ErrAwardIsNone          = errors.New("Award Is None")
	ErrNotFoundAwardPool    = errors.New("Not Found Award Pool")
	ErrNonArriveGrantTime   = errors.New("Non Arrive Grant Time")
	ErrAwardAlreadyAcquired = errors.New("Award Already Aquired")
)

func getAwardAcceptTime() time.Time {
	year, monty, day := time.Now().Date()
	nowStr := fmt.Sprintf("%d-%02d-%02d", year, monty, day)

	curTime0, _ := time.ParseInLocation("2006-01-02", nowStr, time.Local)
	curTime24 := curTime0.AddDate(0, 0, 1)
	return curTime24
}

func main12() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(getAwardAcceptTime())
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
const (
	SERVER_NETWORK = "tcp"
	SERVER_ADDRESS = "127.0.0.1:8085"
	DELIMITER      = '\t'
)

func main123() {
	wg.Add(2)
	go serverGo()
	time.Sleep(500 * time.Millisecond)
	go clientGo(1)
	wg.Wait()
}

var wg sync.WaitGroup

func printLog(role string, sn int, format string, args ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Printf("%s[%d]: %s", role, sn, fmt.Sprintf(format, args...))
}

func printServerLog(format string, args ...interface{}) {
	printLog("Server", 0, format, args...)
}

func printClientLog(sn int, format string, args ...interface{}) {
	printLog("Client", sn, format, args...)
}

func strToInt32(str string) (int32, error) {
	num, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		return 0, fmt.Errorf("\"%s\" is not integer", str)
	}
	if num > math.MaxInt32 || num < math.MinInt32 {
		return 0, fmt.Errorf("%d is not 32-bit integer", num)
	}
	return int32(num), nil
}

func cbrt(param int32) float64 {
	return math.Cbrt(float64(param))
}

// 千万不要使用这个版本的read函数！
//func read(conn net.Conn) (string, error) {
//	reader := bufio.NewReader(conn)
//	readBytes, err := reader.ReadBytes(DELIMITER)
//	if err != nil {
//		return "", err
//	}
//	return string(readBytes[:len(readBytes)-1]), nil
//}

func read(conn net.Conn) (string, error) {
	readBytes := make([]byte, 1)
	var buffer bytes.Buffer
	for {
		_, err := conn.Read(readBytes)
		if err != nil {
			return "", err
		}
		readByte := readBytes[0]
		if readByte == DELIMITER {
			break
		}
		buffer.WriteByte(readByte)
	}
	return buffer.String(), nil
}

func write(conn net.Conn, content string) (int, error) {
	var buffer bytes.Buffer
	buffer.WriteString(content)
	buffer.WriteByte(DELIMITER)
	return conn.Write(buffer.Bytes())
}

func serverGo() {
	//var listener net.Listener
	listener, err := net.Listen(SERVER_NETWORK, SERVER_ADDRESS)
	if err != nil {
		printServerLog("Listen Error: %s", err)
		return
	}
	defer listener.Close()
	printServerLog("Got listener for the server. (local address: %s)", listener.Addr())
	for {
		conn, err := listener.Accept() // 阻塞直至新连接到来。
		if err != nil {
			printServerLog("Accept Error: %s", err)
			continue
		}
		printServerLog("Established a connection with a client application. (remote address: %s)",
			conn.RemoteAddr())
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer func() {
		conn.Close()
		wg.Done()
	}()
	for {
		conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		strReq, err := read(conn)
		if err != nil {
			if err == io.EOF {
				printServerLog("The connection is closed by another side.")
			} else {
				printServerLog("Read Error: %s", err)
			}
			break
		}
		printServerLog("Received request: %s.", strReq)
		intReq, err := strToInt32(strReq)
		if err != nil {
			n, err := write(conn, err.Error())
			printServerLog("Sent error message (written %d bytes): %s.", n, err)
			continue
		}
		floatResp := cbrt(intReq)
		respMsg := fmt.Sprintf("The cube root of %d is %f.", intReq, floatResp)
		n, err := write(conn, respMsg)
		if err != nil {
			printServerLog("Write Error: %s", err)
		}
		printServerLog("Sent response (written %d bytes): %s.", n, respMsg)
	}
}

func clientGo(id int) {
	defer wg.Done()
	conn, err := net.DialTimeout(SERVER_NETWORK, SERVER_ADDRESS, 2*time.Second)
	if err != nil {
		printClientLog(id, "Dial Error: %s", err)
		return
	}
	defer conn.Close()
	printClientLog(id, "Connected to server. (remote address: %s, local address: %s)",
		conn.RemoteAddr(), conn.LocalAddr())
	time.Sleep(200 * time.Millisecond)
	requestNumber := 5
	conn.SetDeadline(time.Now().Add(5 * time.Millisecond))
	for i := 0; i < requestNumber; i++ {
		req := rand.Int31()
		n, err := write(conn, fmt.Sprintf("%d", req))
		if err != nil {
			printClientLog(id, "Write Error: %s", err)
			continue
		}
		printClientLog(id, "Sent request (written %d bytes): %d.", n, req)
	}
	for j := 0; j < requestNumber; j++ {
		strResp, err := read(conn)
		if err != nil {
			if err == io.EOF {
				printClientLog(id, "The connection is closed by another side.")
			} else {
				printClientLog(id, "Read Error: %s", err)
			}
			break
		}
		printClientLog(id, "Received response: %s.", strResp)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func main11() {
	now := time.Now()
	tomorrow := now.AddDate(0, 0, 1)

	fmt.Println(now.Date())
	fmt.Println(tomorrow.Date())
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func main00() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(getAwardGrantTime())
}

func getAwardGrantTime() time.Time {
	year, monty, day := time.Now().Date()
	nowStr := fmt.Sprintf("%d-%02d-%02d", int(year), int(monty), int(day))

	curTime0, _ := time.ParseInLocation("2006-01-02", nowStr, time.Local) // 当前0点：2019-09-26 00:00:00 +0800 CST
	curTime24 := curTime0.AddDate(0, 0, 1)                                // 当前24点：2019-09-27 00:00:00 +0800 CST
	return curTime24
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func main1111() {
	go func() {
		time.Sleep(5 * time.Second)
		sendSignal()
	}()
	handleSignal()
}

func handleSignal() {
	sigRecv1 := make(chan os.Signal, 1)
	sigs1 := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	fmt.Printf("Set notification for %s... [sigRecv1]\n", sigs1)
	signal.Notify(sigRecv1, sigs1...)

	sigRecv2 := make(chan os.Signal, 1)
	sigs2 := []os.Signal{syscall.SIGQUIT}
	fmt.Printf("Set notification for %s... [sigRecv2]\n", sigs2)
	signal.Notify(sigRecv2, sigs2...)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for sig := range sigRecv1 {
			fmt.Printf("Received a signal from sigRecv1: %s\n", sig)
		}
		fmt.Printf("End. [sigRecv1]\n")
		wg.Done()
	}()
	go func() {
		for sig := range sigRecv2 {
			fmt.Printf("Received a signal from sigRecv2: %s\n", sig)
		}
		fmt.Printf("End. [sigRecv2]\n")
		wg.Done()
	}()

	fmt.Println("Wait for 2 seconds... ")
	time.Sleep(2 * time.Second)
	fmt.Printf("Stop notification...")
	signal.Stop(sigRecv1)
	close(sigRecv1)
	fmt.Printf("done. [sigRecv1]\n")
	wg.Wait()
}

func sendSignal() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Fatal Error: %s\n", err)
			debug.PrintStack()
		}
	}()
	// ps aux | grep "signal" | grep -v "grep" | grep -v "go run" | awk '{print $2}'
	cmds := []*exec.Cmd{
		exec.Command("ps", "aux"),
		exec.Command("grep", "anjiadong"),
		exec.Command("grep", "-v", "grep"),
		exec.Command("grep", "-v", "go run"),
		exec.Command("awk", "{print $2}"),
	}
	output, err := runCmds(cmds)
	if err != nil {
		fmt.Printf("%s", output)
		fmt.Printf("Command Execution Error: %s\n", err)
		return
	}

	pids, err := getPids(output)
	if err != nil {
		fmt.Printf("PID Parsing Error: %s\n", err)
		return
	}
	fmt.Printf("Target PID(s):\n%v\n", pids)
	for _, pid := range pids {
		proc, err := os.FindProcess(pid)
		if err != nil {
			fmt.Printf("Process Finding Error: %s\n", err)
			return
		}

		sig := syscall.SIGQUIT
		fmt.Printf("Send signal '%s' to the process (pid=%d)...\n", sig, pid)
		err = proc.Signal(sig)
		if err != nil {
			fmt.Printf("Signal Sending Pid:%d, Error: %s\n", proc.Pid, err)
			continue
		}
	}
}

func getPids(strs []string) ([]int, error) {
	var pids []int
	for _, str := range strs {
		pid, err := strconv.Atoi(strings.TrimSpace(str))
		if err != nil {
			return nil, err
		}
		pids = append(pids, pid)
	}
	return pids, nil
}

func runCmds(cmds []*exec.Cmd) ([]string, error) {
	if cmds == nil || len(cmds) == 0 {
		return nil, errors.New("The cmd slice is invalid!")
	}
	first := true
	var output []byte
	var err error
	for _, cmd := range cmds {
		fmt.Printf("Run command: %v\n", getCmdPlaintext(cmd))
		if !first {
			var stdinBuf bytes.Buffer
			stdinBuf.Write(output)
			cmd.Stdin = &stdinBuf
		}
		var stdoutBuf bytes.Buffer
		cmd.Stdout = &stdoutBuf
		if err = cmd.Start(); err != nil {
			fmt.Println("==", getError(err, cmd).Error())
			return nil, getError(err, cmd)
		}
		if err = cmd.Wait(); err != nil {
			fmt.Println("--", err.Error())
			return nil, getError(err, cmd)
		}
		output = stdoutBuf.Bytes()
		if first {
			first = false
		}
	}
	var lines []string
	var outputBuf bytes.Buffer
	outputBuf.Write(output)
	for {
		line, err := outputBuf.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err.Error())
				return nil, getError(err, nil)
			}
		}
		lines = append(lines, string(line))
	}
	return lines, nil
}

func getCmdPlaintext(cmd *exec.Cmd) string {
	var buf bytes.Buffer
	buf.WriteString(cmd.Path)
	for _, arg := range cmd.Args[1:] {
		buf.WriteRune(' ')
		buf.WriteString(arg)
	}
	return buf.String()
}

func getError(err error, cmd *exec.Cmd, extraInfo ...string) error {
	var errMsg string
	if cmd != nil {
		errMsg = fmt.Sprintf("%s  [%s %v]", err, (*cmd).Path, (*cmd).Args)
	} else {
		errMsg = fmt.Sprintf("%s", err)
	}
	if len(extraInfo) > 0 {
		errMsg = fmt.Sprintf("%s (%v)", errMsg, extraInfo)
	}
	return errors.New(errMsg)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func mainTest1() {
	sigRecv1 := make(chan os.Signal, 1)
	sigs := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	fmt.Printf("set notification for %s...[sigRecv1]\n", sigs)
	signal.Notify(sigRecv1, sigs...)

	sigRecv2 := make(chan os.Signal, 1)
	sigs1 := []os.Signal{syscall.SIGQUIT}
	fmt.Printf("set notification for %s...[sigRecv2]\n", sigs1)
	signal.Notify(sigRecv2, sigs1...)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for sig := range sigRecv1 {
			fmt.Printf("Received a signal form sigRecv1: %s\n", sig)
		}
		fmt.Printf("End. [sigRecv1]\n")
		wg.Done()
	}()

	go func() {
		for sig := range sigRecv2 {
			fmt.Printf("Received a signal form sigRecv2: %s\n", sig)
		}
		fmt.Printf("End. [sigRecv2]\n")
		wg.Done()
	}()

	fmt.Println("Wait for 2 seconds...")
	time.Sleep(2 * time.Second)
	fmt.Printf("Stop notification...")
	signal.Stop(sigRecv1)
	close(sigRecv1)
	fmt.Printf("done [sigRecv1]\n")

	wg.Wait()
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func mainTest() {
	reader, writer, err := os.Pipe()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)

	input := getProcessInfo()

	go func() {
		output := make([]byte, 100000)
		n, err := reader.Read(output)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("Read %d byte(s)\n", n)
		wg.Done()
	}()

	go func() {
		time.Sleep(3 * time.Second)
		n, err := writer.Write(input)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("written %d byte(s)\n", n)
		wg.Done()
	}()

	wg.Wait()
}

func getProcessInfo() []byte {
	ps := exec.Command("ps", "aux")
	grep := exec.Command("grep", "anjiadong")

	var outPutBufPS bytes.Buffer
	ps.Stdout = &outPutBufPS

	if err := ps.Start(); err != nil {
		fmt.Println(err.Error())
		return nil
	}
	if err := ps.Wait(); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	grep.Stdin = &outPutBufPS
	var outPutBufGrep bytes.Buffer
	grep.Stdout = &outPutBufGrep
	if err := grep.Start(); err != nil {
		fmt.Println(err.Error())
		return nil
	}
	if err := grep.Wait(); err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return outPutBufGrep.Bytes()
}
