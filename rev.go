// Copyright © 2015 Erik Brady <brady@dvln.org>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// The rev.go file in package devline concerns itself with managing and
// working with revisions of development lines for the 'dvln' tool.

package devline

import (
	"io"

	"github.com/dvln/pkg"
)

// RevType is an indicator if a revision is static or dynamic (or unknown)
type RevType string

// devline content revision type, is it dynamic or static (or unknown)?
const (
	UnknownRevType RevType = ""
	Dynamic        RevType = "dynamic"
	Static         RevType = "static"
)

// Rev identifies the data associated with a revision of a devline.
// To operate on it one would use the devline.Revision interface which
// also allows one to better test/mock things/etc.
type Rev struct {
	Name     string                       `json:"name"`
	Desc     string                       `json:"desc,omitempty"`
    pkg      pkg.Revision
	Attrs    map[string]string            `json:"attrs,omitempty"`
	Contacts []string                     `json:"contacts,omitempty"`
	Access   map[string]string            `json:"access,omitempty"`
	Type     RevType                      `json:"type,omitempty"`
	Pkgs     []pkg.Revision               `json:"pkgs,omitempty"`
}

type Revision interface {
	Read(r io.Reader) error
	Write(w io.Writer) error
	//Commit() error
	//Push() error
    //Type() error
	//PkgRev() error
	//Contains(pkgSel string) error
}
