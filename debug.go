// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fuse

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/jacobsa/fuse/fuseops"
)

// Decide on the name of the given op.
func opName(op interface{}) string {
	// We expect all ops to be pointers.
	t := reflect.TypeOf(op).Elem()

	// Strip the "Op" from "FooOp".
	return strings.TrimSuffix(t.Name(), "Op")
}

func describeRequest(op interface{}) string {
	var components []string
	addComponent := func(format string, v ...interface{}) {
		components = append(components, fmt.Sprintf(format, v...))
	}

	switch typedOp := op.(type) {
	case *interruptOp:
		addComponent("fuseid 0x%08x", typedOp.FuseID)

	case *unknownOp:
		if typedOp.Inode != 0 {
			addComponent("inode %v", typedOp.Inode)
		}
		addComponent("opcode %d", typedOp.OpCode)

	case *fuseops.SetInodeAttributesOp:
		if typedOp.Inode != 0 {
			addComponent("inode %v", typedOp.Inode)
		}
		if typedOp.OpContext.Pid != 0 {
			addComponent("PID %+v", typedOp.OpContext.Pid)
		}
		if typedOp.Size != nil {
			addComponent("size %d", *typedOp.Size)
		}
		if typedOp.Mode != nil {
			addComponent("mode %v", *typedOp.Mode)
		}
		if typedOp.Atime != nil {
			addComponent("atime %v", *typedOp.Atime)
		}
		if typedOp.Mtime != nil {
			addComponent("mtime %v", *typedOp.Mtime)
		}

	case *fuseops.RenameOp:
		if typedOp.OpContext.Pid != 0 {
			addComponent("PID %+v", typedOp.OpContext.Pid)
		}
		addComponent("old_parent %v", typedOp.OldParent)
		addComponent("old_name %q", typedOp.OldName)
		addComponent("new_parent %v", typedOp.NewParent)
		addComponent("new_name %q", typedOp.NewName)

	case *fuseops.ReadFileOp:
		if typedOp.Inode != 0 {
			addComponent("inode %v", typedOp.Inode)
		}
		if typedOp.OpContext.Pid != 0 {
			addComponent("PID %+v", typedOp.OpContext.Pid)
		}
		addComponent("handle %d", typedOp.Handle)
		addComponent("offset %d", typedOp.Offset)
		addComponent("%d bytes", typedOp.Size)

	case *fuseops.WriteFileOp:
		if typedOp.Inode != 0 {
			addComponent("inode %v", typedOp.Inode)
		}
		if typedOp.OpContext.Pid != 0 {
			addComponent("PID %+v", typedOp.OpContext.Pid)
		}
		addComponent("handle %d", typedOp.Handle)
		addComponent("offset %d", typedOp.Offset)
		addComponent("%d bytes", len(typedOp.Data))

	case *fuseops.RemoveXattrOp:
		if typedOp.Inode != 0 {
			addComponent("inode %v", typedOp.Inode)
		}
		if typedOp.Name != "" {
			addComponent("name %q", typedOp.Name)
		}
		if typedOp.OpContext.Pid != 0 {
			addComponent("PID %+v", typedOp.OpContext.Pid)
		}
		addComponent("name %s", typedOp.Name)

	case *fuseops.GetXattrOp:
		if typedOp.Inode != 0 {
			addComponent("inode %v", typedOp.Inode)
		}
		if typedOp.Name != "" {
			addComponent("name %q", typedOp.Name)
		}
		if typedOp.OpContext.Pid != 0 {
			addComponent("PID %+v", typedOp.OpContext.Pid)
		}
		addComponent("name %s", typedOp.Name)

	case *fuseops.SetXattrOp:
		if typedOp.Inode != 0 {
			addComponent("inode %v", typedOp.Inode)
		}
		if typedOp.Name != "" {
			addComponent("name %q", typedOp.Name)
		}
		if typedOp.OpContext.Pid != 0 {
			addComponent("PID %+v", typedOp.OpContext.Pid)
		}
		addComponent("name %s", typedOp.Name)

	case *fuseops.FallocateOp:
		if typedOp.Inode != 0 {
			addComponent("inode %v", typedOp.Inode)
		}
		if typedOp.OpContext.Pid != 0 {
			addComponent("PID %+v", typedOp.OpContext.Pid)
		}
		addComponent("offset %d", typedOp.Offset)
		addComponent("length %d", typedOp.Length)
		addComponent("mode %d", typedOp.Mode)

	case *fuseops.ReleaseFileHandleOp:
		if typedOp.OpContext.Pid != 0 {
			addComponent("PID %+v", typedOp.OpContext.Pid)
		}
		addComponent("handle %d", typedOp.Handle)
	default:
		return opName(op) // Fallback if type is not handled
	}

	if len(components) == 0 {
		return opName(op)
	}
	return fmt.Sprintf("%s (%s)", opName(op), strings.Join(components, ", "))
}

func describeResponse(op interface{}) string {
	var components []string
	addComponent := func(format string, v ...interface{}) {
		components = append(components, fmt.Sprintf(format, v...))
	}

	switch typedOp := op.(type) {
	case *fuseops.OpenFileOp:
		addComponent("handle %d", typedOp.Handle)
	default:
		return opName(op) // Fallback if type is not handled
	}

	return fmt.Sprintf("%s (%s)", opName(op), strings.Join(components, ", "))
}
