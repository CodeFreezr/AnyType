// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package tag

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/ccsafe/tag"
)

type TagChan interface { // bidirectional channel
	TagROnlyChan // aka "<-chan" - receive only
	TagSOnlyChan // aka "chan<-" - send only
}

type TagROnlyChan interface { // receive-only channel
	RequestTag() (dat tag.TagAny)        // the receive function - aka "some-new-Tag-var := <-MyKind"
	TryTag() (dat tag.TagAny, open bool) // the multi-valued comma-ok receive function - aka "some-new-Tag-var, ok := <-MyKind"
}

type TagSOnlyChan interface { // send-only channel
	ProvideTag(dat tag.TagAny) // the send function - aka "MyKind <- some Tag"
}