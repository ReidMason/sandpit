namespace ClientLib.Models;

public class ContactDtoResponse
{
    public int Id { get; set; }
    public string Name { get; set; } = "";
    public string Email { get; set; } = "";
    public PhoneNumber PhoneNumber { get; set; } = new PhoneNumber();
}

public class PhoneNumber
{
    public string Number { get; set; } = "";
   public string CountryCode { get; set; } = "";
}