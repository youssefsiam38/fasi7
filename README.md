# Fasi7

## Overview

Fasi7 is a command-line interface (CLI) tool designed to help developers translate their localization files using AI. It automates the process of translating content from one language to multiple target languages, preserving placeholders and adapting to the context of your business.

## Installation

Download the latest release from the [releases page](https://github.com/youssefsiam38/fasi7/releases)

## Configuration

Fasi7 uses a configuration file to store settings. By default, it looks for a file named `.fasi7.yaml` in the current directory.

### Creating a Configuration File

To create a new configuration file, run:

```
fasi7 init
```

### Configuration Options

The following options can be set in the configuration file:

- `openai.apiKey`: Your OpenAI API key
- `openai.model`: The OpenAI model to use for translations
- `openai.systemPrompt`: Custom system prompt for the AI (optional)
- `outputLocales`: List of target locales for translation
- `inputLocale`: The source locale to translate from
- `businessDescription`: A description of your business to provide context for translations
- `dir`: The directory containing localization files
- `ignoreFilesWithContent`: (Optional) If specified, files containing this string will be ignored

## Usage

To run Fasi7 and translate your localization files:

```
fasi7
```

This command will:

1. Read the configuration file
2. Recursively search for files in the specified directory
3. Identify files matching the input locale
4. Translate the content to each specified output locale
5. Write the translated content to new files, preserving the directory structure

## Command-line Options

- `--config`: Specify a custom path for the configuration file (default: `.fasi7.yaml`)

Example:
```
fasi7 --config /path/to/custom/config.yaml
```

## How It Works

1. Fasi7 uses the OpenAI API to perform translations.
2. It sends each file's content to the AI along with a system prompt that includes:
   - The source and target languages
   - Instructions to preserve placeholders
   - Your business description for context
3. The AI generates a translation, which is then saved to a new file in the appropriate locale directory.

## Error Handling

Fasi7 will exit with an error message in the following cases:

- No configuration file found
- Missing required configuration options
- File read/write errors
- OpenAI API errors

## Concurrent Processing

Fasi7 translates files concurrently to improve performance when dealing with multiple output locales or many files.

## Best Practices

1. Provide a detailed business description in your configuration to improve translation accuracy and context.
2. Use meaningful placeholder names in your localization files to ensure they are correctly preserved.
3. Choose an appropriate model for your needs.

## Limitations

- Fasi7 relies on the OpenAI API, which may have usage limits and costs associated with it.
- The quality of translations depends on the AI model used and the context provided.

## Contributing

Contributions are welcome! Please submit issues or pull requests to help improve Fasi7.

## License
MIT License