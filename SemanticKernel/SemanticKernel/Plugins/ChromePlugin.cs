using Microsoft.SemanticKernel;
using OpenQA.Selenium;
using OpenQA.Selenium.Chrome;
using System.ComponentModel;
using WebDriverManager;
using WebDriverManager.DriverConfigs.Impl;
using System.Text.Json;
using System.Collections.Generic;

namespace SemanticKernel.Plugins;

// Model classes for page elements
public class PageElements
{
    public Metadata? Metadata { get; set; }
    public List<Form>? Forms { get; set; }
    public Navigation? Navigation { get; set; }
    public List<Link>? ContentLinks { get; set; }
    public List<Button>? Buttons { get; set; }
    public List<Input>? StandaloneInputs { get; set; }
    public PageAnalysis? PageAnalysis { get; set; }
}

public class Metadata
{
    public string? Title { get; set; }
    public string? Url { get; set; }
    public string? Description { get; set; }
    public string? H1 { get; set; }
}

public class Form
{
    public string? Id { get; set; }
    public string? Name { get; set; }
    public string? Action { get; set; }
    public string? Method { get; set; }
    public List<Input>? Inputs { get; set; }
    public List<Button>? Buttons { get; set; }
    public string? Purpose { get; set; }
}

public class Input
{
    public string? Type { get; set; }
    public string? Id { get; set; }
    public string? Name { get; set; }
    public string? Placeholder { get; set; }
    public string? Value { get; set; }
    public string? Label { get; set; }
    public bool Required { get; set; }
    public string? CssSelector { get; set; }
}

public class Button
{
    public string? Type { get; set; }
    public string? Id { get; set; }
    public string? Text { get; set; }
    public string? CssSelector { get; set; }
}

public class Navigation
{
    public List<Link>? Links { get; set; }
}

public class Link
{
    public string? Text { get; set; }
    public string? Url { get; set; }
    public string? CssSelector { get; set; }
}

public class PageAnalysis
{
    public bool HasLogin { get; set; }
    public bool HasSearch { get; set; }
    public bool HasRegistration { get; set; }
    public bool HasNewsletter { get; set; }
    public bool HasContactForm { get; set; }
}

/// <summary>
/// ChromePlugin for Semantic Kernel that enables browser automation with Selenium WebDriver
/// </summary>
public class ChromePlugin : IDisposable
{
    private IWebDriver? _driver;
    private bool _isInitialized = false;
    private readonly object _lock = new();

    /// <summary>
    /// Initializes the Chrome WebDriver
    /// </summary>
    public void EnsureDriverInitialized()
    {
        if (_isInitialized && _driver != null)
            return;

        lock (_lock)
        {
            if (_isInitialized && _driver != null)
                return;

            try
            {
                // Setup Chrome WebDriver using WebDriverManager to handle driver binaries
                new DriverManager().SetUpDriver(new ChromeConfig());

                var options = new ChromeOptions();
                options.AddArgument("--start-maximized");
                // Uncomment the following for headless operation if needed
                // options.AddArgument("--headless=new");

                _driver = new ChromeDriver(options);
                _driver.Manage().Timeouts().ImplicitWait = TimeSpan.FromSeconds(10);
                _isInitialized = true;

                Console.WriteLine("Chrome WebDriver initialized successfully.");
            }
            catch (Exception ex)
            {
                Console.WriteLine($"Error initializing Chrome WebDriver: {ex.Message}");
                throw;
            }
        }
    }

    [KernelFunction, Description("Navigate to a URL")]
    public void Navigate(string url)
    {
        Console.WriteLine($"Navigating to URL: {url}");
        try
        {
            EnsureDriverInitialized();

            if (_driver == null)
                throw new InvalidOperationException("Chrome driver is not initialized.");

            _driver.Navigate().GoToUrl(url);
        }
        catch (Exception ex)
        {
            Console.WriteLine($"Error navigating to {url}: {ex.Message}");
            throw;
        }
    }

    [KernelFunction, Description("Gets the current page elements including CSS selectors and text")]
    public PageElements GetPageElements()
    {
        Console.WriteLine("Getting page elements in AI-optimized format");
        try
        {
            EnsureDriverInitialized();

            if (_driver == null)
                throw new InvalidOperationException("Chrome driver is not initialized.");

            // Use JavaScript to extract various elements and structure them for AI analysis
            string jsonResult = (string)((IJavaScriptExecutor)_driver).ExecuteScript(@"
                function getPageElementsForAI() {
                    // Get page metadata
                    const metadata = {
                        title: document.title,
                        url: window.location.href,
                        description: document.querySelector('meta[name=""description""]')?.content || '',
                        h1: Array.from(document.querySelectorAll('h1')).map(el => el.textContent.trim()).join(' | ')
                    };
                    
                    // Extract form information
                    const forms = Array.from(document.querySelectorAll('form')).slice(0, 5).map(form => {
                        // Get inputs in this form
                        const formInputs = Array.from(form.querySelectorAll('input, select, textarea')).map(input => {
                            return {
                                type: input.type || input.tagName.toLowerCase(),
                                id: input.id || '',
                                name: input.name || '',
                                placeholder: input.placeholder || '',
                                value: input.value || '',
                                label: document.querySelector(`label[for=""${input.id}""]`)?.textContent.trim() || '',
                                required: input.required || false,
                                cssSelector: input.id ? `#${input.id}` : `${input.tagName.toLowerCase()}[name=""${input.name}""]`
                            };
                        });
                        
                        // Get submit buttons in this form
                        const formButtons = Array.from(form.querySelectorAll('button, input[type=""submit""], input[type=""button""]')).map(button => {
                            return {
                                type: button.type || button.tagName.toLowerCase(),
                                id: button.id || '',
                                text: button.textContent.trim() || button.value || '',
                                cssSelector: button.id ? `#${button.id}` : `${button.tagName.toLowerCase()}[type=""${button.type}""]`
                            };
                        });
                        
                        return {
                            id: form.id || '',
                            name: form.getAttribute('name') || '',
                            action: form.action || '',
                            method: form.method || 'get',
                            inputs: formInputs,
                            buttons: formButtons,
                            purpose: detectFormPurpose(form, formInputs)
                        };
                    });
                    
                    // Extract primary navigation
                    const mainNav = {
                        links: Array.from(document.querySelectorAll('nav a, header a, .navigation a, .menu a')).slice(0, 15).map(link => ({
                            text: link.textContent.trim(),
                            url: link.href,
                            cssSelector: link.id ? `#${link.id}` : createSelector(link)
                        }))
                    };
                    
                    // Extract main content links
                    const mainContentLinks = Array.from(document.querySelectorAll('main a, #content a, .content a, article a')).slice(0, 10).map(link => ({
                        text: link.textContent.trim(),
                        url: link.href,
                        cssSelector: link.id ? `#${link.id}` : createSelector(link)
                    }));
                    
                    // Extract main buttons (not in forms)
                    const mainButtons = Array.from(document.querySelectorAll('button:not(form button), [role=""button""]:not(form [role=""button""])')).slice(0, 10).map(button => ({
                        text: button.textContent.trim(),
                        type: button.type || '',
                        id: button.id || '',
                        cssSelector: button.id ? `#${button.id}` : createSelector(button)
                    }));
                    
                    // Extract inputs not in forms
                    const standaloneInputs = Array.from(document.querySelectorAll('input:not(form input), select:not(form select), textarea:not(form textarea)')).slice(0, 10).map(input => ({
                        type: input.type || input.tagName.toLowerCase(),
                        id: input.id || '',
                        name: input.name || '',
                        placeholder: input.placeholder || '',
                        label: document.querySelector(`label[for=""${input.id}""]`)?.textContent.trim() || '',
                        cssSelector: input.id ? `#${input.id}` : createSelector(input)
                    }));
                    
                    // Helper function to create a selector
                    function createSelector(element) {
                        if (element.id) return `#${element.id}`;
                        if (element.name) return `${element.tagName.toLowerCase()}[name=""${element.name}""]`;
                        
                        // Try classes
                        if (element.classList && element.classList.length > 0) {
                            const mainClass = element.classList[0];
                            return `${element.tagName.toLowerCase()}.${mainClass}`;
                        }
                        
                        // Fallback to attribute selectors
                        if (element.getAttribute('type')) {
                            return `${element.tagName.toLowerCase()}[type=""${element.getAttribute('type')}""]`;
                        }
                        
                        if (element.getAttribute('role')) {
                            return `${element.tagName.toLowerCase()}[role=""${element.getAttribute('role')}""]`;
                        }
                        
                        return element.tagName.toLowerCase();
                    }
                    
                    // Helper function to detect form purpose
                    function detectFormPurpose(form, inputs) {
                        const formText = form.textContent.toLowerCase();
                        const inputTypes = inputs.map(i => i.type);
                        const inputNames = inputs.map(i => i.name.toLowerCase());
                        const inputLabels = inputs.map(i => i.label.toLowerCase());
                        const allText = [...inputNames, ...inputLabels, formText].join(' ').toLowerCase();
                        
                        if (allText.includes('login') || allText.includes('sign in') || 
                            (inputTypes.includes('password') && (inputTypes.includes('email') || inputTypes.includes('text')))) {
                            return 'login';
                        }
                        
                        if (allText.includes('register') || allText.includes('sign up') || allText.includes('create account')) {
                            return 'registration';
                        }
                        
                        if (allText.includes('search') || form.querySelector('input[type=""search""]') !== null) {
                            return 'search';
                        }
                        
                        if (allText.includes('subscribe') || allText.includes('newsletter')) {
                            return 'newsletter';
                        }
                        
                        if (allText.includes('contact') || allText.includes('message') || allText.includes('feedback')) {
                            return 'contact';
                        }
                        
                        return 'unknown';
                    }
                    
                    // Combine all data in a structured way
                    return {
                        metadata: metadata,
                        forms: forms,
                        navigation: mainNav,
                        contentLinks: mainContentLinks,
                        buttons: mainButtons,
                        standaloneInputs: standaloneInputs,
                        pageAnalysis: {
                            hasLogin: forms.some(f => f.purpose === 'login'),
                            hasSearch: forms.some(f => f.purpose === 'search'),
                            hasRegistration: forms.some(f => f.purpose === 'registration'),
                            hasNewsletter: forms.some(f => f.purpose === 'newsletter'),
                            hasContactForm: forms.some(f => f.purpose === 'contact')
                        }
                    };
                }
                
                return JSON.stringify(getPageElementsForAI(), null, 2);
            ");
            
            // Deserialize the JSON result into the PageElements object
            return System.Text.Json.JsonSerializer.Deserialize<PageElements>(jsonResult, new System.Text.Json.JsonSerializerOptions 
            { 
                PropertyNameCaseInsensitive = true 
            }) ?? new PageElements();
        }
        catch (Exception ex)
        {
            Console.WriteLine($"Error getting AI-optimized page elements: {ex.Message}");
            throw;
        }
    }
    
    [KernelFunction, Description("Click on an element using a CSS selector")]
    public void ClickElement(string cssSelector)
    {
        Console.WriteLine($"Clicking element: {cssSelector}");
        try
        {
            EnsureDriverInitialized();

            if (_driver == null)
                throw new InvalidOperationException("Chrome driver is not initialized.");

            var element = _driver.FindElement(By.CssSelector(cssSelector));
            element.Click();
        }
        catch (NoSuchElementException ex)
        {
            throw new NoSuchElementException($"Element with selector '{cssSelector}' was not found.", ex);
        }
        catch (Exception ex)
        {
            Console.WriteLine($"Error clicking element '{cssSelector}': {ex.Message}");
            throw;
        }
    }
    
    [KernelFunction, Description("Type text into an element using a CSS selector")]
    public void TypeText(string cssSelector, string text)
    {
        Console.WriteLine($"Typing text into element {cssSelector}: {text}");
        try
        {
            EnsureDriverInitialized();

            if (_driver == null)
                throw new InvalidOperationException("Chrome driver is not initialized.");

            var element = _driver.FindElement(By.CssSelector(cssSelector));
            element.Clear();
            element.SendKeys(text);
        }
        catch (NoSuchElementException ex)
        {
            throw new NoSuchElementException($"Element with selector '{cssSelector}' was not found.", ex);
        }
        catch (ElementNotInteractableException ex)
        {
            throw new ElementNotInteractableException($"Element with selector '{cssSelector}' is not interactable. It may be hidden or disabled.", ex);
        }
        catch (Exception ex)
        {
            Console.WriteLine($"Error typing text into element '{cssSelector}': {ex.Message}");
            throw;
        }
    }
    
    /// <summary>
    /// Dispose of the WebDriver when done
    /// </summary>
    public void Dispose()
    {
        if (_driver != null)
        {
            try
            {
                _driver.Quit();
                _driver.Dispose();
            }
            catch (Exception ex)
            {
                Console.WriteLine($"Error disposing Chrome driver: {ex.Message}");
            }
            finally
            {
                _driver = null;
                _isInitialized = false;
            }
        }
    }
} 