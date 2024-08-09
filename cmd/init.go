package cmd

import (
	"github.com/spf13/cobra"
	"github.com/youssefsiam38/fasi7/utils"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Fasi7",
	Long:  `Initialize Fasi7 by creating a config file`,
	Run: func(cmd *cobra.Command, args []string) {
		err := utils.WriteFile(".fasi7.yaml", []byte("dir: ./locales # Required\r\nbusinessDescription: | # This will be ignored if you specify a systemPrompt\r\n  <Business description here>\r\nopenai:\r\n  apiKey: ${OPENAI_API_KEY} # Required\r\n  model: gpt-4o-mini # Required\r\n  # systemPrompt: | # Optional\r\ninputLocale: en # Required\r\noutputLocales: # Required\r\n  - ar\r\n  - de\r\n  - es\r\n  - fr\r\n  - ru\r\nignoreFilesWithContent: '{}' # Optional"))
		if err != nil {
			panic(err)
		}
	},
}
