using Serilog;
using Serilog.Sinks.Grafana.Loki;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

// Set up Serilog
Log.Logger = new LoggerConfiguration()
    .ReadFrom.Configuration(builder.Configuration)
    .Enrich.FromLogContext()
    .Enrich.WithCorrelationId(addValueIfHeaderAbsence: true)
    .WriteTo.Console()
    .WriteTo.GrafanaLoki("http://localhost:3100", labels: new[] { new LokiLabel { Key = "app", Value = "dotnet-1" } })
    .WriteTo.Seq("http://localhost:5341", apiKey: "HWzCLFw0ctMiMHcTSkGx")
    .CreateLogger();

builder.Host.UseSerilog();
builder.Services.AddHttpContextAccessor();

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();

var summaries = new[]
{
    "Freezing", "Bracing", "Chilly", "Cool", "Mild", "Warm", "Balmy", "Hot", "Sweltering", "Scorching"
};

app.MapGet("/weatherforecast", (ILogger<Program> logger) => app.MapGet("/weatherforecast", async (ILogger<Program> logger, IHttpContextAccessor httpContextAccessor) =>
{
    try
    {
        Random rnd = new Random();
        int num  = rnd.Next(1, 10);
        logger.LogInformation("Weather forecast endpoint was called {number}", num);
        if (num > 5)
            throw new Exception("Something went wrong");

        var forecast =  Enumerable.Range(1, 5).Select(index =>
                new WeatherForecast
                (
                    DateOnly.FromDateTime(DateTime.Now.AddDays(index)),
                    Random.Shared.Next(-20, 55),
                    summaries[Random.Shared.Next(summaries.Length)]
                ))
            .ToArray();

        logger.LogInformation("Got weather forecast {@forecast}", forecast);
        return forecast;
    }
    catch (Exception ex)
    {
        logger.LogError(ex, "Error in weather forecast endpoint");
        return new WeatherForecast[] { };
    }
})
.WithName("GetWeatherForecast")
.WithOpenApi();

app.Run();

// Ensure to flush and stop Serilog on shutdown
Log.CloseAndFlush();

record WeatherForecast(DateOnly Date, int TemperatureC, string? Summary)
{
    public int TemperatureF => 32 + (int)(TemperatureC / 0.5556);
}
