using Godot;
using System;

public partial class Coin : Area2D
{ 
	public override void _Ready()
	{
		this.BodyEntered += OnBodyEntered;
	}
	
	private void OnBodyEntered(Node2D body)
	{
		QueueFree();
		GameManager.Instance.AddScore(1);
	}
}
