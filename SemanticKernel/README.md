# Semantic Kernel Sample Project with Ollama and Chrome Automation

This is a sample project demonstrating the usage of Microsoft's Semantic Kernel SDK with a local Ollama instance and Chrome browser automation.

## Prerequisites

- .NET 8.0 SDK
- [Ollama](https://ollama.ai/) installed and running locally
- Gemma 3B model available in your Ollama instance
- Google Chrome browser installed
- Chrome WebDriver (automatically managed by the WebDriverManager package)

## Setup

1. Clone this repository
2. Make sure Ollama is running locally
3. If you haven't already, pull the Gemma model:
   ```bash
   ollama pull gemma3:latest
   ```
4. Update the `appsettings.json` file if needed with your Ollama model and endpoint configuration.

## Running the Project

```bash
cd SemanticKernel
dotnet run
```

## Features

This sample demonstrates:

- Basic Semantic Kernel setup with Ollama
- Chrome browser automation using Selenium WebDriver
- Natural language commands for browser interaction
- Configurable system prompts to guide the AI's behavior

## Chrome Automation

The project includes a ChromePlugin with the following capabilities:

- **Navigate to URLs**: Open any website by URL
- **Click Elements**: Click on any element using CSS selectors
- **Type Text**: Enter text into form fields using CSS selectors
- **Get Current URL**: Retrieve the current page URL
- **Get Element Text**: Extract text from elements using CSS selectors
- **Get Page Content**: Retrieve the entire text content of the current page

Example commands:

```
Navigate to google.com
Search for "Semantic Kernel" on Google
Go to github.com and click on the Sign in button
Get the content of this page
```

## Project Structure

- `Program.cs` - Main entry point, interactive console, and system prompt definition
- `Plugins/TextPlugin.cs` - Text processing plugin with semantic functions
- `Plugins/ChromePlugin.cs` - Chrome automation plugin using Selenium
- `appsettings.json` - Configuration for Ollama settings

## System Prompt

The system prompt that guides the AI on how to use the Chrome plugin is defined directly in the `Program.cs` file. You can modify it there to change how the AI responds to browser automation requests.

## Learn More

- [Semantic Kernel Documentation](https://learn.microsoft.com/en-us/semantic-kernel/overview/)
- [GitHub Repository](https://github.com/microsoft/semantic-kernel)
- [Ollama](https://ollama.ai/)
- [Selenium WebDriver](https://www.selenium.dev/documentation/webdriver/) 