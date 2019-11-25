/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"gitlabctl/handlers"
	"gitlabctl/model"
	"net/http"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// groupCmd represents the group command
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "manipulation for groups",
	Long:  `Manipulate groups, list/copy/create/delete.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := &http.Client{
			Timeout: time.Second * 30,
		}

		if len(args) != 0 {
			runList(args[0], target, client)
		}

	},
}

//func init() {
//	lsCmd.AddCommand(groupCmd)
//}

//groupPages brings model.Groups to this package
type groupPages model.Groups

//Groups is the appended pagesGroup
type Groups struct {
	Group []groupPages
}

// list groups on gitlab
func (pg groupPages) list(client *http.Client, url, token string) (box Groups, err error) {

	items := []groupPages{}
	box = Groups{items}

	get := handlers.Requester{
		Client: client,
		Url:    url + token}

	opts := "&per_page=5"
	totalpages := handlers.ScanTotalPages(client, get.Url+opts)
	opts = opts + "&page="

	for page := 1; page <= totalpages; page++ {
		get.Url = url + token + opts + strconv.Itoa(page)
		_, pages := get.Req()
		err = json.Unmarshal(pages, &pg)
		if err != nil {
			return box, err
		}

		for _, g := range pg {
			item := groupPages{g}
			box.Group = append(box.Group, item)
		}

	}

	return box, nil

}

//func runList(arg, token string, client *http.Client) {
//	g := groupPages{}
//	groupSearch := 0
//	groups, _ := g.list(client, getGroups, token)
//	for _, grp := range groups.Group {
//		if name != "" {
//			if name == grp[0].Path {
//				groupSearch = grp[0].ID
//			}
//			if grp[0].ID == groupSearch || grp[0].ParentID == groupSearch {
//				fmt.Println(grp[0].FullPath + "\t\t" + grp[0].Path)
//			}
//			continue
//		}
//		fmt.Println(grp[0].FullPath + "\t\t" + grp[0].Path)
//	}
//}

func runGroup(arg, token string, client *http.Client) {
	fmt.Println(arg)
}
