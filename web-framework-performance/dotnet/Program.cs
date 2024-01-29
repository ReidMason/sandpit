using Serilog;
using Serilog.Sinks.Grafana.Loki;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
Log.Logger = new LoggerConfiguration()
    .WriteTo.Console()
    .WriteTo.GrafanaLoki("http://localhost:3100", new[] { new LokiLabel { Key = "job", Value = "dotnet" }})
    .CreateLogger();

// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();

app.MapGet("/api/hello", () =>
{
    // Create an object to log
    var myObject = new { Name = "John Doe", Age = 30 };

    Log.Information("Hello, World! {@myObject}", myObject);

    return "Hello, World!";
})
.WithName("hello")
.WithOpenApi();

app.Run();

Log.CloseAndFlush();
