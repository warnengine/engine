package main

import (
	"strings"
	"unsafe"
)

func ogldebugcb(source uint32, gltype uint32, id uint32, severity uint32, length int32, message string, userParam unsafe.Pointer) {
	if !strings.Contains(message, "GL_STATIC_DRAW") {
		println(message)
	}
}
