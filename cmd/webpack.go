package cmd

import (
	"github.com/armour/jarvis/utils"
	"github.com/spf13/cobra"
)

var (
	coverallToken      string
	projectName        string
	projectShortname   string
	projectDescription string
	projectBgColor     string
	projectThemeColor  string
	license            string
	ci                 bool
	docker             bool
	materialize        bool
	postgres           bool
	react              bool
	redis              bool
	typescript         bool
)

var webpackCmd = &cobra.Command{
	Use:   "webpack",
	Short: "Start new project using webpack template",
	Long:  "Start new project using webpack template",
	Run: func(cmd *cobra.Command, args []string) {
		templatePath := "../new/webpack"
		requireMap := map[string]interface{}{
			"ci":          ci,
			"coverage":    coverallToken,
			"docker":      docker,
			"materialize": materialize,
			"postgres":    postgres,
			"react":       react,
			"redis":       redis,
			"typescript":  typescript,
		}
		replaceMap := map[string]interface{}{
			"ci":                 ci,
			"coverallToken":      coverallToken,
			"docker":             docker,
			"materialize":        materialize,
			"postgres":           postgres,
			"projectName":        projectName,
			"projectShortname":   projectShortname,
			"projectDescription": projectDescription,
			"projectBgColor":     projectBgColor,
			"projectThemeColor":  projectThemeColor,
			"react":              react,
			"redis":              redis,
			"typescript":         typescript,
			"license":            license,
		}
		if projectShortname == "" {
			projectShortname = projectName
		}
		utils.GenerateFile(templatePath, "", requireMap, replaceMap)
	},
}

func init() {
	newCmd.AddCommand(webpackCmd)

	webpackCmd.Flags().StringVarP(&coverallToken, "coverallToken", "c", "", "The token for coverall code coverage. (optional)")
	webpackCmd.Flags().StringVarP(&projectName, "projectName", "n", "", "The name for this project. (required)")
	webpackCmd.Flags().StringVarP(&projectShortname, "projectShortname", "s", "", "The short name for this project. (optional)")
	webpackCmd.Flags().StringVarP(&projectDescription, "projectDescription", "d", "project description here", "The description for this project. (optional)")
	webpackCmd.Flags().StringVarP(&projectBgColor, "projectBgColor", "b", "#2196f3", "The background color for this project. (optional)")
	webpackCmd.Flags().StringVarP(&projectThemeColor, "projectThemeColor", "t", "#2196f3", "The theme color for this project. (optional)")
	webpackCmd.Flags().StringVarP(&license, "license", "l", "MIT", "The license for this project, default is \"MIT\". (\"MIT\", \"GPL\", \"Apache\"")
	webpackCmd.Flags().BoolVar(&ci, "ci", true, "The flag to enable continuous integration support.")
	webpackCmd.Flags().BoolVar(&docker, "docker", true, "The flag to enable docker support.")
	webpackCmd.Flags().BoolVar(&materialize, "materialize", true, "The flag to enable materialize support.")
	webpackCmd.Flags().BoolVar(&postgres, "postgres", true, "The flag to enable postgres support.")
	webpackCmd.Flags().BoolVar(&react, "react", true, "The flag to enable react support.")
	webpackCmd.Flags().BoolVar(&redis, "redis", true, "The flag to enable redis support.")
	webpackCmd.Flags().BoolVar(&typescript, "typescript", true, "The flag to enable typescript support.")

	webpackCmd.MarkFlagRequired("projectName")
}
