// Copyright 2015,2016,2017,2018,2019 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// int position in 2d space
package ivector2d

import "testing"

func TestVt_Eq(t *testing.T) {
	v1 := Vt{0, 0}
	v2 := Vt{0, 1}
	t.Logf("v1 %v v2 %v eq %v", v1, v2, v1.Eq(v2))
}