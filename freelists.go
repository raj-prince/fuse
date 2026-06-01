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
	"github.com/jacobsa/fuse/internal/buffer"
)

////////////////////////////////////////////////////////////////////////
// buffer.InMessage
////////////////////////////////////////////////////////////////////////

func (c *Connection) getInMessage() *buffer.InMessage {
	select {
	case x := <-c.inMessages:
		return x
	default:
		return buffer.NewInMessage()
	}
}

func (c *Connection) putInMessage(x *buffer.InMessage) {
	select {
	case c.inMessages <- x:
	default:
		// Pool is full, discard the buffer.
	}
}

////////////////////////////////////////////////////////////////////////
// buffer.OutMessage
////////////////////////////////////////////////////////////////////////

func (c *Connection) getOutMessage() *buffer.OutMessage {
	var x *buffer.OutMessage
	select {
	case x = <-c.outMessages:
	default:
		x = new(buffer.OutMessage)
	}
	x.Reset()

	return x
}

func (c *Connection) putOutMessage(x *buffer.OutMessage) {
	select {
	case c.outMessages <- x:
	default:
		// Pool is full, discard the buffer.
	}
}
