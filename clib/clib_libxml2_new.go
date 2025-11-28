//go:build libxml2_new
// +build libxml2_new

package clib

// #include <libxml/parser.h>
import "C"
import "unsafe"

// c14nParseNodeContextPtr handles pointer conversion for xmlParseInNodeContext.
// Newer versions of libxml2 (2.12+) require double pointer cast.
func c14nParseNodeContextPtr(ret *C.xmlNodePtr) unsafe.Pointer {
	return unsafe.Pointer((**C.xmlNode)(unsafe.Pointer(ret)))
}
