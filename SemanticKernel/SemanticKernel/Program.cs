using Microsoft.Extensions.Configuration;
using Microsoft.SemanticKernel;
using Microsoft.SemanticKernel.ChatCompletion;
using Microsoft.SemanticKernel.Connectors.OpenAI;
using OpenQA.Selenium;
using SemanticKernel.Plugins;

// Load configuration
var configuration = new ConfigurationBuilder()
    .AddJsonFile("appsettings.json", optional: true, reloadOnChange: true)
    .Build();

// Get Ollama configuration
string ollamaModelId = configuration["Ollama:ModelId"] ?? "ERROR";
string ollamaEndpoint = configuration["Ollama:ApiEndpoint"] ?? "http://localhost:11434";

Console.WriteLine($"Using Ollama with model: {ollamaModelId}");
Console.WriteLine($"API endpoint: {ollamaEndpoint}");

// Create a kernel builder and add Ollama chat completion service
var builder = Kernel.CreateBuilder();
#pragma warning disable SKEXP0070
builder.AddOllamaChatCompletion(
    modelId: ollamaModelId,
    endpoint: new Uri(ollamaEndpoint),
    serviceId: "ollama"
    );
#pragma warning restore SKEXP0070

// Enable planning
OpenAIPromptExecutionSettings openAIPromptExecutionSettings = new()
{
    FunctionChoiceBehavior = FunctionChoiceBehavior.Auto()
};

Kernel kernel = builder.Build();
var chatCompletionService = kernel.GetRequiredService<IChatCompletionService>();

var browserPlugin = new ChromePlugin();
browserPlugin.Navigate("https://hub.youroverseashome.com/login?redirect=/");
Thread.Sleep(1000);
var content = browserPlugin.GetPageElements();
// browserPlugin.ClickElement(".dark\\:inset-ring-white\\/15");
// Thread.Sleep(1000);
// browserPlugin.TypeText("#docsearch-input", "Margin");
// browserPlugin.TypeText("#docsearch-input", Keys.Return);
kernel.Plugins.AddFromObject(browserPlugin, "Browser");

// Create a history store the conversation
var history = new ChatHistory();

// Initiate a back-and-forth chat
string? userInput;
do {
    // Collect user input
    Console.Write("User > ");
    userInput = Console.ReadLine();

    // Add user input
    history.AddUserMessage(userInput);

    // Get the response from the AI
    var result = await chatCompletionService.GetChatMessageContentAsync(
        history,
        executionSettings: openAIPromptExecutionSettings,
        kernel: kernel);

    // Print the results
    Console.WriteLine("Assistant > " + result);

    // Add the message from the agent to the chat history
    history.AddMessage(result.Role, result.Content ?? string.Empty);
} while (userInput is not null);
