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

// Package devline has structures, interfaces and methods related to devline
// manipulation and use.  It is primarily focused on devline versions since that
// tends to be what is always used when it comes to devlines (and if no version
// is given the "default" VCS version would be used to grab the appropriate
// devline.
package devline

// FIXME: erik: may need a base devline Defn type struct but perhaps not,
//    leave it out for now
