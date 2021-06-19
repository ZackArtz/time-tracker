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
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/zackartz/tt/models"
)

// startCmd represents the start command
var (
	Comment  string
	Project  string
	Category string
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Starts a new timestamp",
		Long: `Starts a new timestamp For example:
	
	tt start --project skill-tracker --category devops --comment "it works!"
	
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			jsonData := models.Timestamp{
				Comment:  Comment,
				Project:  Project,
				Category: Category,
			}
			jsonValue, err := json.Marshal(jsonData)
			if err != nil {
				return err
			}
			response, err := http.Post("http://localhost:6969/api/v1/create", "application/json", bytes.NewBuffer(jsonValue))
			if err != nil {
				return err
			}
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return err
			}
			resp := models.Timestamp{}
			err = json.Unmarshal(body, &resp)
			if err != nil {
				return err
			}
			cmd.Printf("Started timestamp with uuid %s at time %s\n", resp.UUID, resp.StartTime)
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	startCmd.Flags().StringVarP(&Comment, "comment", "", "", "Comment for the Timestamp")
	startCmd.Flags().StringVarP(&Category, "category", "c", "", "Category for the Timestamp")
	startCmd.Flags().StringVarP(&Project, "project", "p", "", "Project for the Timestamp (required)")
	startCmd.MarkFlagRequired("project")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
