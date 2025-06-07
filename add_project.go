package main

import (
	"context"
	"fmt"

	todoist "github.com/sachaos/todoist/lib"
	"github.com/urfave/cli/v2"
)

func AddProject(c *cli.Context) error {
	client := GetClient(c)

	project := todoist.Project{}
	if !c.Args().Present() {
		return CommandFailed
	}

	project.Name = c.Args().First()
	parentName := c.String("parent-name") 
	if parentName != "" {
		parentID := client.Store.Projects.GetIDByName(parentName)
		if parentID == "" {
			return fmt.Errorf("Did not find a project named '%v'", parentName)
		}
		project.ParentID = &parentID
	} else {
		parentID := c.String("parent-id")
		project.ParentID = &parentID
	}

	if c.String("color") != "0" {
		project.Color = c.String("color")
	}
	project.ItemOrder = c.Int("item-order")

	if err := client.AddProject(context.Background(), project); err != nil {
		return err
	}

	return Sync(c)
}
