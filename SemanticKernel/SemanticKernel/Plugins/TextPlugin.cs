using Microsoft.SemanticKernel;
using System.ComponentModel;

namespace SemanticKernel.Plugins;

public class TextPlugin
{
    [KernelFunction, Description("Summarize the given text")]
    public async Task<string> Summarize(
        [Description("The text to summarize")] string input,
        Kernel kernel)
    {
        KernelFunction summarizeFunction = kernel.CreateFunctionFromPrompt(
            "Summarize the following text in a concise way: {{$input}}");
        
        var result = await kernel.InvokeAsync(summarizeFunction, new() { ["input"] = input });
        return result.ToString();
    }
    
    [KernelFunction, Description("Translate text to the specified language")]
    public async Task<string> Translate(
        [Description("The text to translate")] string input,
        [Description("The target language")] string language,
        Kernel kernel)
    {
        KernelFunction translateFunction = kernel.CreateFunctionFromPrompt(
            "Translate the following text to {{$language}}: {{$input}}");
        
        var result = await kernel.InvokeAsync(translateFunction, new() 
        { 
            ["input"] = input,
            ["language"] = language
        });
        
        return result.ToString();
    }
} 