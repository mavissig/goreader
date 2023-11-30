package reader

import (
	"goreader_cli/internal/domain"
	"log"
	"path/filepath"
)

type DBReader interface {
	Read(string) domain.MainStruct
	Mock()
}

type Handler interface {
	handleRequest(request string) (string, domain.MainStruct)
	setNext(handler Handler)
}

type Reader struct {
	next Handler
}

func (r *Reader) setNext(handler Handler) {
	r.next = handler
}

func (r *Xml) handleRequest(request string) (string, domain.MainStruct) {
	suffix := filepath.Ext(request)
	if ".xml" == suffix {
		data := r.Read(request)
		return suffix, data
	} else if r.next != nil {
		return r.next.handleRequest(request)
	}

	log.Printf("[ERROR]: Unsupported Format \"%s\"\n", filepath.Ext(request))
	return "", domain.MainStruct{}
}

func (r *Json) handleRequest(request string) (string, domain.MainStruct) {
	suffix := filepath.Ext(request)
	if ".json" == suffix {
		data := r.Read(request)
		return suffix, data
	} else if r.next != nil {
		return r.next.handleRequest(request)
	}

	log.Printf("[ERROR]: Unsupported Format \"%s\"\n", filepath.Ext(request))
	return "", domain.MainStruct{}
}

func (r *Reader) Read(path string) (string, domain.MainStruct) {
	jsonHandler := &Json{}
	xmlHandler := &Xml{}

	jsonHandler.setNext(xmlHandler)

	return jsonHandler.handleRequest(path)
}

func New() *Reader {
	return &Reader{}
}
