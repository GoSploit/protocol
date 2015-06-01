package protocol

import (
	"encoding/json"
)

type Packet struct {
	Type string
	ID   int64
	Msg  Message
}

type RawPacket struct {
	Type string
	ID   int64
	Msg  *json.RawMessage
}

type Message interface{}

type ChDirCommand struct {
	NewDir string
}

type ChDirResponse struct {
	Error string
}

type ListCommand struct {
}

type File struct {
	Name  string
	IsDir bool
	Size  int64
}

type ListResponse struct {
	Files []File
}

type GetCommand struct {
	File string
}

type GetResponse struct {
	Data []byte
}

type PutCommand struct {
	File string
	Data []byte
}

type PutResponse struct {
}

type GetSessionsRequest struct {
}

type SessionInfo struct {
	ID int64
}

type GetSessionsResponse struct {
	Sessions []SessionInfo
}

type SelectSessionRequest struct {
	ID int64
}

type SelectSessionResponse struct {
	Error string
}

func (p *Packet) UnmarshalJSON(data []byte) error {
	var rp RawPacket
	json.Unmarshal(data, &rp)
	p.ID = rp.ID
	p.Type = rp.Type
	switch rp.Type {
	case "ChDirCommand":
		var msg ChDirCommand
		json.Unmarshal(*rp.Msg, &msg)
		p.Msg = msg
	case "ChDirResponse":
		var msg ChDirResponse
		json.Unmarshal(*rp.Msg, &msg)
		p.Msg = msg
	case "ListCommand":
		var msg ListCommand
		json.Unmarshal(*rp.Msg, &msg)
		p.Msg = msg
	case "ListResponse":
		var msg ListResponse
		json.Unmarshal(*rp.Msg, &msg)
		p.Msg = msg
	case "GetCommand":
		var msg GetCommand
		json.Unmarshal(*rp.Msg, &msg)
		p.Msg = msg
	case "GetResponse":
		var msg GetResponse
		json.Unmarshal(*rp.Msg, &msg)
		p.Msg = msg
	case "PutCommand":
		var msg PutCommand
		json.Unmarshal(*rp.Msg, &msg)
		p.Msg = msg
	case "PutResponse":
		var msg PutResponse
		json.Unmarshal(*rp.Msg, &msg)
		p.Msg = msg
	case "GetSessionsRequest":
		var msg GetSessionsRequest
		json.Unmarshal(*rp.Msg, &msg)
		p.Msg = msg
	case "GetSessionsResponse":
		var msg GetSessionsResponse
		json.Unmarshal(*rp.Msg, &msg)
		p.Msg = msg
	case "SelectSessionRequest":
		var msg SelectSessionRequest
		json.Unmarshal(*rp.Msg, &msg)
		p.Msg = msg
	case "SelectSessionResponse":
		var msg SelectSessionResponse
		json.Unmarshal(*rp.Msg, &msg)
		p.Msg = msg
	}
	return nil
}

func (p *Packet) MarshalJSON() ([]byte, error) {
	var rp RawPacket
	rp.ID = p.ID
	switch msg := p.Msg.(type) {
	case ChDirCommand:
		p.Type = "ChDirCommand"
		buf, _ := json.Marshal(msg)
		jbuf := json.RawMessage(buf)
		rp.Msg = &jbuf
	case ChDirResponse:
		p.Type = "ChDirResponse"
		buf, _ := json.Marshal(msg)
		jbuf := json.RawMessage(buf)
		rp.Msg = &jbuf
	case ListCommand:
		p.Type = "ListCommand"
		buf, _ := json.Marshal(msg)
		jbuf := json.RawMessage(buf)
		rp.Msg = &jbuf
	case ListResponse:
		p.Type = "ListResponse"
		buf, _ := json.Marshal(msg)
		jbuf := json.RawMessage(buf)
		rp.Msg = &jbuf
	case GetCommand:
		p.Type = "GetCommand"
		buf, _ := json.Marshal(msg)
		jbuf := json.RawMessage(buf)
		rp.Msg = &jbuf
	case GetResponse:
		p.Type = "GetResponse"
		buf, _ := json.Marshal(msg)
		jbuf := json.RawMessage(buf)
		rp.Msg = &jbuf
	case PutCommand:
		p.Type = "PutCommand"
		buf, _ := json.Marshal(msg)
		jbuf := json.RawMessage(buf)
		rp.Msg = &jbuf
	case PutResponse:
		p.Type = "PutResponse"
		buf, _ := json.Marshal(msg)
		jbuf := json.RawMessage(buf)
		rp.Msg = &jbuf
	case GetSessionsRequest:
		p.Type = "GetSessionsRequest"
		buf, _ := json.Marshal(msg)
		jbuf := json.RawMessage(buf)
		rp.Msg = &jbuf
	case GetSessionsResponse:
		p.Type = "GetSessionsResponse"
		buf, _ := json.Marshal(msg)
		jbuf := json.RawMessage(buf)
		rp.Msg = &jbuf
	case SelectSessionRequest:
		p.Type = "SelectSessionRequest"
		buf, _ := json.Marshal(msg)
		jbuf := json.RawMessage(buf)
		rp.Msg = &jbuf
	case SelectSessionResponse:
		p.Type = "SelectSessionResponse"
		buf, _ := json.Marshal(msg)
		jbuf := json.RawMessage(buf)
		rp.Msg = &jbuf
	}
	rp.Type = p.Type
	data, err := json.Marshal(rp)
	return data, err
}
