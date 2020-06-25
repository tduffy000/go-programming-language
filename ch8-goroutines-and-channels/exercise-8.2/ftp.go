// TODO: still several commands to implement
// and general clean up required

// based on RFC 959 --> https://tools.ietf.org/html/rfc959
//
// 5.  DECLARATIVE SPECIFICATIONS
//
//    5.1.  MINIMUM IMPLEMENTATION
//
//       In order to make FTP workable without needless error messages, the
//       following minimum implementation is required for all servers:
//
//          TYPE - ASCII Non-print
//          MODE - Stream
//          STRUCTURE - File, Record
//          COMMANDS - USER, QUIT, PORT,
//                     TYPE, MODE, STRU,
//                       for the default values
//                     RETR, STOR,
//                     NOOP.
//
//       The default values for transfer parameters are:
//
//          TYPE - ASCII Non-print
//          MODE - Stream
//          STRU - File
//
//       All hosts must accept the above as the standard defaults.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8000, "The port to bind the FTP Server to.")
	flag.Parse()

	server := new(FtpServer)
	server.users = make(map[string]bool)
	server.port = port
	listener, err := net.Listen("tcp4", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal("Could not start server: ", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print("Error accepting new connection: %v\n", err)
		}
		go OpenConn(server, conn).serve()
	}
}

func OpenConn(server *FtpServer, pi net.Conn) *connection {
	return &connection{pi, "", nil, server, false, nil, ""}
}

type connection struct {
	protocolInterpreter net.Conn
	dataPort            string
	dataConnection      net.Listener
	server              *FtpServer
	loggedIn            bool
	errorState          error
	previousCmd         string
}

// TODO: update to accept {}interface
func (conn *connection) println(s string) {
	if conn.errorState != nil {
		return
	}
	_, conn.errorState = fmt.Fprintf(conn.protocolInterpreter, s+"\n\r")
}

func (conn *connection) getDataConnection() (io.ReadWriteCloser, error) {
	if conn.previousCmd == "PORT" {
		c, err := net.Dial("tcp", conn.dataPort)
		if err != nil {
			return nil, err
		}
		return c, err
	} else if conn.previousCmd == "PASV" {
		c, err := conn.dataConnection.Accept()
		if err != nil {
			return nil, err
		}
		return c, err
	} else {
		return nil, fmt.Errorf("Need previous command == PORT | PASV")
	}
}

// TODO: make sure to clean up stuff
func (conn *connection) serve() {
	conn.println(Response{220, "Ready."}.String())
	scanner := bufio.NewScanner(conn.protocolInterpreter)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) == 0 {
			continue
		}
		cmd := strings.ToUpper(fields[0])
		var args []string
		var hasArgs bool
		if len(fields) > 1 {
			args = fields[1:]
			hasArgs = true
		}
		var response *Response
		switch cmd {
		case "USER":
			if !hasArgs {
				response = &Response{503, "Expected arguments, got none."}
			} else {
				response = conn.server.User(args[0])
			}
			if response.code == 230 {
				conn.loggedIn = true
			}
		case "QUIT":
			if conn.loggedIn {
				response = conn.server.Quit()
			} else {
				response = &Response{503, "User is already logged out."}
			}
			if response.code == 221 {
				conn.loggedIn = false
			}
		case "PORT":
			if !hasArgs {
				response = &Response{503, "Expected arguments, got none."}
			} else {
				if conn.loggedIn {
					conn.dataPort, response = conn.server.Port(args[0])
				} else {
					response = &Response{530, "Please login first"}
				}
			}
		case "STOR":
			if !hasArgs {
				response = &Response{}
			} else {
				response, file := conn.server.Stor(args[0])
				conn.println(response.String())
				w, err := conn.getDataConnection()
				if err != nil {
					response = &Response{425, "Cannot open the data connections"}
				}
				defer w.Close()
				_, err = io.Copy(file, w)
				if err != nil {
					response = &Response{450, "File unavailable."}
				} else {
					response = &Response{226, "File transfer complete"}
				}
			}
		case "PASV":
			if conn.dataPort == "" {
				response = &Response{425, "Please define a port first."}
			} else {
				var err error
				conn.dataConnection, err = net.Listen("tcp4", "")
				if err != nil {
					response = &Response{}
				}
				_, port, err := net.SplitHostPort(conn.dataConnection.Addr().String())
				if err != nil {
					response = &Response{}
				}
				ip, _, err := net.SplitHostPort(conn.protocolInterpreter.LocalAddr().String())
				if err != nil {
					response = &Response{}
				}
				response = conn.server.Pasv(ip, port)
			}
		case "NOOP":
			response = conn.server.Noop()
		default:
			response = &Response{502, "Command not implemented."}
		}
		conn.println(response.String())
	}
}

type Response struct {
	code int
	body string
}

func (r Response) String() string {
	return fmt.Sprintf("%v: %v", r.code, r.body)
}

type FtpServer struct {
	users map[string]bool
	port  int
}

// The argument field is a Telnet string identifying the user.
// The user identification is that which is required by the
// server for access to its file system. (4.1.1)
func (f *FtpServer) User(login string) *Response {
	if _, ok := f.users[login]; ok {
		return &Response{503, "User already logged in."}
	}
	f.users[login] = true
	return &Response{230, "User logged in, proceed."}
}

// The argument is a HOST-PORT specification for the data port
// to be used in data connection. The argument is the concatenation of a
// 32-bit internet host address and a 16-bit TCP port address.
// This address information is broken into 8-bit fields and the
// value of each field is transmitted as a decimal number (in
// character string representation).  The fields are separated
// by commas. (4.1.2)
func (f *FtpServer) Port(hostport string) (string, *Response) {
	address, err := convertFTPFromIP(hostport)
	if err != nil {
		return "", &Response{501, fmt.Sprintf("Got an error: %v", err)}
	}
	return address, &Response{225, "Opened data connection on " + address}
}

func convertFTPFromIP(hostport string) (string, error) {
	host, portStr, err := net.SplitHostPort(hostport)
	if err != nil {
		return "", err
	}
	addr, err := net.ResolveIPAddr("ip4", host)
	if err != nil {
		return "", err
	}
	port, err := strconv.ParseInt(portStr, 10, 64)
	if err != nil {
		return "", err
	}
	ip := addr.IP.To4()
	return fmt.Sprintf("%d,%d,%d,%d,%d,%d", ip[0], ip[1], ip[2], ip[3], port/256, port%256), nil
}

// This command requests the server-DTP to "listen" on a data
// port (which is not its default data port) and to wait for a
// connection rather than initiate one upon receipt of a
// transfer command.  The response to this command includes the
// host and port address this server is listening on. (4.1.2)
func (f *FtpServer) Pasv(ip, port string) *Response {
	var a, b, c, d byte
	var p1, p2 int
	addr, err := convertFTPFromIP(fmt.Sprintf("%s:%s", ip, port))
	if err != nil {
		return &Response{425, fmt.Sprintf("Got an error: %v", err)}
	}
	_, err = fmt.Sscanf(addr, "%d,%d,%d,%d,%d,%d", &a, &b, &c, &d, &p1, &p2)
	if err != nil {
		return &Response{425, "Got an improperly formatted port."}
	}
	return &Response{227, "Listening on " + addr}
}

// This command terminates a USER and if file transfer is not
// in progress, the server closes the control connection. (4.1.1)
func (f *FtpServer) Quit() *Response {
	return &Response{221, "Closing connection, adios."}
}

func (f *FtpServer) Stor(filename string) (*Response, *os.File) {
	file, err := os.Create(filename)
	if err != nil {
		return &Response{550, "Problem creating file."}, nil
	}
	return &Response{150, "Ok to send data."}, file
}

func (f *FtpServer) Noop() *Response {
	return &Response{200, "Okay."}
}
