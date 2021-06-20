/*
Copyright © 2020 Zachary Myers <zackmyers@lavabit.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/zackartz/tt/models"
)

// stopCmd represents the stop command
var (
	uuid    string
	stopCmd = &cobra.Command{
		Use:   "stop",
		Short: "End a running timestamp",
		Long: `End a currently running timestamp, note you do not need the full uuid
For example:
	
	tt end --uuid 6dfe75f0
	// full uuid is 6dfe75f0-f215-41de-b0e2-5f057da7fd30

You do not need the full uuid, even though providing 6 charactes is recommended.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := http.Get("http://localhost:6969/api/v1/active-timestamps")
			if err != nil {
				return err
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			tss := []models.Timestamp{}
			err = json.Unmarshal(body, &tss)

			count := 0
			index := 0
			for i, ts := range tss {
				if strings.Contains(ts.UUID, args[0]) {
					count += 1
					index = i
				}
			}

			var (
				startTime time.Time
				endTime   time.Time
			)
			if count == 1 {
				uuid = tss[index].UUID
				startTime = tss[index].StartTime
				endTime = time.Now()
			} else {
				return errors.New("multiple matching uuids found please be more specific")
			}

			response, err := http.Get(fmt.Sprintf("http://localhost:6969/api/v1/end/%s", uuid))
			if err != nil {
				return err
			}
			if response.StatusCode != 200 {
				return errors.New("something went wrong")
			}
			cmd.Printf("Ended timestamp with uuid %s with a length of %s\n", uuid, endTime.Sub(startTime))
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(stopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
