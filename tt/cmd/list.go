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
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/zackartz/tt/models"
)

// listCmd represents the list command
var (
	active   bool
	project  string
	category string
	listCmd  = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "Shows all timestamps",
		Long: `Shows all of the timestamps available, for example:

		tt ls

You can use the -a or --active to show only active timestamps`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if active {
				var err error
				resp, err := http.Get("http://localhost:6969/api/v1/active-timestamps")
				if err != nil {
					return err
				}
				tss := []models.Timestamp{}
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					return err
				}
				err = json.Unmarshal(body, &tss)
				if err != nil {
					return err
				}
				data := [][]string{}
				for _, ts := range tss {
					data = append(data, []string{ts.UUID, ts.StartTime.Format("2006-01-02 3:4:5pm"), ts.Project, ts.Category, ts.Comment})
				}
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"UUID", "Start Time", "Project", "Category", "Comment"})
				table.SetBorder(false)
				table.AppendBulk(data)
				table.Render()
				return nil
			} else {
				var err error
				resp := &http.Response{}
				if project != "" {
					resp, err = http.Get(fmt.Sprintf("http://localhost:6969/api/v1/timestamps/%s", project))
					if err != nil {
						return err
					}
				} else {
					resp, err = http.Get("http://localhost:6969/api/v1/timestamps")
					if err != nil {
						return err
					}
				}

				tss := []models.Timestamp{}
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					return err
				}
				err = json.Unmarshal(body, &tss)
				if err != nil {
					return err
				}
				data := [][]string{}
				var total time.Time
				for _, ts := range tss {
					if category != "" {
						if ts.Category == category {
							data = append(data, []string{fmt.Sprintf("%t", ts.Active), ts.StartTime.Format("2006-01-02 3:4:5pm"), ts.Project, ts.Category, ts.Comment})
							total = total.Add(ts.EndTime.Sub(ts.StartTime))
						}
					} else {
						data = append(data, []string{fmt.Sprintf("%t", ts.Active), ts.StartTime.Format("2006-01-02 3:4:5pm"), ts.Project, ts.Category, ts.Comment})
						total = total.Add(ts.EndTime.Sub(ts.StartTime))
					}
				}
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"Active", "Started At", "Project", "Category", "Comment"})
				table.SetBorder(false)
				table.SetFooter([]string{"", "", "", "Total", fmt.Sprintf("%vhr(s)%dm", total.Hour(), total.Minute())})
				table.AppendBulk(data)
				table.Render()
				return nil
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")
	listCmd.Flags().BoolVarP(&active, "active", "a", false, "Easily toggle between showing active and inactive timestamps")
	listCmd.Flags().StringVarP(&project, "project", "p", "", "Get timestamps for a specific project")
	listCmd.Flags().StringVarP(&category, "category", "c", "", "Get timestamps for a specific category")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
