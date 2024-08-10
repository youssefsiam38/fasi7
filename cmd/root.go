package cmd

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/youssefsiam38/fasi7/utils"
)

var cfgFile string
var noConfig bool

func init() {
	cobra.OnInitialize(initConfig)
	// config
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", ".fasi7.yaml", "config file path")
}

func initConfig() {
	viper.SetConfigFile(cfgFile)
	if err := viper.ReadInConfig(); err != nil {
		noConfig = true
	}
}

var rootCmd = &cobra.Command{
	Use:   "fasi7",
	Short: "Fasi7 is a localization files translation tool using AI",
	Long:  `Fasi7 was made to help developers to translate their localization files using AI. for more information visit github.com/youssefsiam38/fasi7`,
	Run: func(cmd *cobra.Command, args []string) {
		if noConfig {
			fmt.Println("No config file found, run 'fasi7 init' to create one")
			os.Exit(1)
		}
		ctx := cmd.Context()
		openaiClient := openai.NewClient(utils.GetConfigString("openai.apiKey"))
		outputLocales := viper.GetStringSlice("outputLocales")
		if len(outputLocales) == 0 {
			fmt.Println("No output locales found in config file")
			os.Exit(1)
		}

		inputLocale := utils.GetConfigString("inputLocale")
		if inputLocale == "" {
			fmt.Println("No input locale found in config file")
			os.Exit(1)
		}

		businessDescription := utils.GetConfigString("businessDescription")
		if businessDescription == "" {
			fmt.Println("No business description found in config file")
			os.Exit(1)
		}

		files, err := utils.GetFilesRecursive(utils.GetConfigString("dir"))
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		wg := sync.WaitGroup{}

		inputFilesCount := 0
		translatedFilesCount := 0
		for _, file := range files {
			if !strings.Contains(file, fmt.Sprintf("/%s/", inputLocale)) {
				continue
			}

			fileContent, err := os.ReadFile(file)
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}

			if utils.GetConfigString("ignoreFilesWithContent") != "" && strings.Contains(string(fileContent), utils.GetConfigString("ignoreFilesWithContent")) {
				continue
			}

			inputFilesCount++

			for _, outputLocale := range outputLocales {
				wg.Add(1)
				go func(fileContent []byte, file, outputLocale string) {
					defer wg.Done()
					systemPrompt := utils.GetConfigString("openai.systemPrompt")
					if systemPrompt == "" {
						systemPrompt = fmt.Sprintf(`
					You are an expert language translator. Your task is to translate the text provided to you from %s to %s, while preserving any placeholders that may be present in the source text.

					The input text will be the original file content, and the output text should be parsable content, do not include any comments or unnecessary whitespace.
					
					If the input text contains any placeholders (e.g. {name}, {count}, etc.), ensure that these placeholders remain unchanged in the Arabic translation.
					
					Provide a high-quality translation that conveys the original meaning accurately and idiomatically. Do not simply perform a literal word-for-word translation, but adapt the phrasing and grammar to produce natural-sounding text that is appropriate for the target audience for the business described below.
					
					Business Description:
					%s
					`, utils.IsoToLanguage(inputLocale), utils.IsoToLanguage(outputLocale), businessDescription)
					}

					seed := 1

					res, err := openaiClient.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
						Model: utils.GetConfigString("openai.model"),
						Messages: []openai.ChatCompletionMessage{
							{
								Role:    "system",
								Content: systemPrompt,
							},
							{
								Role:    "user",
								Content: string(fileContent),
							},
						},
						Temperature: 0.2,
						Seed:        &seed,
					})
					if err != nil {
						fmt.Println("Error:", err)
						os.Exit(1)
					}
					err = utils.WriteFile(strings.ReplaceAll(file, fmt.Sprintf("/%s/", inputLocale), fmt.Sprintf("/%s/", outputLocale)), []byte(res.Choices[0].Message.Content))
					if err != nil {
						fmt.Println("Error:", err)
						os.Exit(1)
					}
					translatedFilesCount++
				}(fileContent, file, outputLocale)
			}
		}

		wg.Wait()
		fmt.Printf("Translated %d files from %d input files\n", translatedFilesCount, inputFilesCount)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
