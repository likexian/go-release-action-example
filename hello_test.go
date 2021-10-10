/*
 * Copyright 2021 Li Kexian
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Automate publishing Go build binary artifacts to GitHub releases through GitHub Actions.
 * https://www.likexian.com/
 */

package goreleaseactionexample

import "testing"

func TestHello(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"", "Hello !"},
		{"world", "Hello world!"},
	}
	for _, tt := range tests {
		if got := Hello(tt.in); got != tt.out {
			t.Errorf("Hello() = %v, want %v", got, tt.out)
		}
	}
}
