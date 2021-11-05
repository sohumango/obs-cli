package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/andreykaipov/goobs/api/requests/scenes"
	"github.com/spf13/cobra"
)

var (
	sceneCmd = &cobra.Command{
		Use:   "scene",
		Short: "manage scenes",
		Long:  `The scene command manages scenes`,
		RunE:  nil,
	}

	switchSceneCmd = &cobra.Command{
		Use:   "switch",
		Short: "Switch to a different scene",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("switch requires a scene name as argument")
			}
			return switchScene(strings.Join(args, " "))
		},
	}
	listSceneCmd = &cobra.Command{
		Use:   "list",
		Short: "List all scene",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listScenes()
		},
	}
)

func switchScene(scene string) error {
	r := scenes.SetCurrentSceneParams{
		SceneName: scene,
	}
	_, err := client.Scenes.SetCurrentScene(&r)
	return err
}

func listScenes() error {
	{
		resp, err := client.Scenes.GetSceneList()
		if err != nil {
			return err
		}

		fmt.Println("Scene List")
		fmt.Println("===============")
		for _, v := range resp.Scenes {
			fmt.Println(v.Name)
		}
	}
	return nil
}

func init() {
	sceneCmd.AddCommand(switchSceneCmd)
	sceneCmd.AddCommand(listSceneCmd)
	rootCmd.AddCommand(sceneCmd)
}
