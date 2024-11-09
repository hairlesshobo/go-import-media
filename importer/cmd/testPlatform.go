// =================================================================================
//
//		ccmm - https://www.foxhollow.cc/projects/ccmm/
//
//	 Connection Church Media Manager, aka ccmm, is a tool for managing all
//   aspects of produced media- initial import from removable media,
//   synchronization with clients and automatic data replication and backup
//
//		Copyright (c) 2024 Steve Cross <flip@foxhollow.cc>
//
//		Licensed under the Apache License, Version 2.0 (the "License");
//		you may not use this file except in compliance with the License.
//		You may obtain a copy of the License at
//
//		     http://www.apache.org/licenses/LICENSE-2.0
//
//		Unless required by applicable law or agreed to in writing, software
//		distributed under the License is distributed on an "AS IS" BASIS,
//		WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//		See the License for the specific language governing permissions and
//		limitations under the License.
//
// =================================================================================

package cmd

import (
	"ccmm/util"

	"github.com/spf13/cobra"
)

var (
	testPlatformCmd = &cobra.Command{
		Use:   "test_platform",
		Short: "Test the platform you are running on for support",
		Long:  `Quick check to see if the running platform is supported. Only used for debugging while building`,

		Run: func(cmd *cobra.Command, args []string) {
			util.TestPlatform()
		},
	}
)

func init() {
	rootCmd.AddCommand(testPlatformCmd)
}
