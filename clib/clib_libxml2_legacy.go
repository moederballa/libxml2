//go:build !libxml2_new
// +build !libxml2_new

package clib

// #include <libxml/parser.h>
import "C"
import "unsafe"

// c14nParseNodeContextPtr handles pointer conversion for xmlParseInNodeContext.
// Older versions of libxml2 use direct pointer conversion.
func c14nParseNodeContextPtr(ret *C.xmlNodePtr) unsafe.Pointer {
	return unsafe.Pointer(ret)
}
