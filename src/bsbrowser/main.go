package main

import (
	"fmt"
	"html/template"
	"net"
	"net/http"
	"sort"
	"time"
)

type Server struct {
	Address     net.IP
	Port        uint16
	Name        string
	Map         string
	Capacity    uint8
	PlayerCount uint8
}

type Servers []Server

func (s Servers) Len() int           { return len(s) }
func (s Servers) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Servers) Less(i, j int) bool { return s[i].PlayerCount > s[j].PlayerCount }

var tpl *template.Template

func getString(buf []byte, index int) string {
	i := index
	slen := 0
	for buf[i] != 0 {
		slen++
		i++
	}
	return string(buf[index : index+slen])
}

func ServeServerList(rw http.ResponseWriter, req *http.Request) {
	addr, _ := net.ResolveUDPAddr("udp", "hl2master.steampowered.com:27011")
	conn, _ := net.DialUDP("udp", nil, addr)
	defer conn.Close()
	conn.Write([]byte("\x31\xFF0.0.0.0:0\x00\\appid\\225600\\gamedir\\berimbau\\empty\\1\x00"))

	buf := make([]byte, 512)
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	readBytes, err := conn.Read(buf)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	var srvs []Server
	i := 6
	for i+6 < readBytes {
		srv := Server{}
		srv.Address = net.IPv4(buf[i], buf[i+1], buf[i+2], buf[i+3])
		srv.Port = uint16(buf[i+4])<<8 + uint16(buf[i+5])

		addr, _ = net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", srv.Address.String(), srv.Port))
		conn, _ = net.DialUDP("udp", nil, addr)
		defer conn.Close()
		conn.Write([]byte("\xFF\xFF\xFF\xFF\x54Source Engine Query\x00"))

		srvbuf := make([]byte, 512)
		conn.SetReadDeadline(time.Now().Add(500 * time.Millisecond))
		srvReadBytes, err := conn.Read(srvbuf)
		_ = srvReadBytes
		if err != nil {
			fmt.Println(err)
			i += 6
			continue
		}
		curIndex := 6
		srv.Name = getString(srvbuf, curIndex)
		curIndex += len(srv.Name) + 1
		srv.Map = getString(srvbuf, curIndex)
		curIndex += len(srv.Map) + 1
		curIndex += len(getString(srvbuf, curIndex)) + 1 //skip folder
		curIndex += len(getString(srvbuf, curIndex)) + 1 //skip game
		curIndex += 2                                    //skip gameid
		srv.PlayerCount = srvbuf[curIndex]
		srv.Capacity = srvbuf[curIndex+1]
		srvs = append(srvs, srv)
		i += 6
	}
	sort.Sort(Servers(srvs))
	rw.Header().Add("content-type", "text/html; charset=utf-8")
	tpl.Execute(rw, srvs)
	return
}

func ServeStaticFiles(rw http.ResponseWriter, req *http.Request) {
	http.ServeFile(rw, req, "."+req.URL.Path)
}

func main() {
	tpl = template.Must(template.ParseFiles("./templates/index.tpl"))
	http.HandleFunc("/", ServeServerList)
	http.HandleFunc("/static/", ServeStaticFiles)
	http.ListenAndServe(":3000", nil)
}
