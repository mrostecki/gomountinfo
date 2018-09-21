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

package gomountinfo

// ParseMountTable parses the mount info available for the current process.
// Not implemented for Windows!
func ParseMountTable(f FilterFunc) ([]*MountInfo, error) {
	// Do NOT return an error!
	return nil, nil
}

// ParseMountTablePid parses the mount info available for the given process
// Not implemented for Windows!
func ParseMountTablePid(pid int, f FilterFunc) ([]*MountInfo, error) {
	// Do NOT return an error!
	return nil, nil
}
