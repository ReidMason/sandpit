using Godot;
using System;
using scripts.domain;

public partial class Player : CharacterBody2D
{
	public const float Speed = 100.0f;
	
	public override void _Ready()
	{
		Keys.SetupInputMap();
	}

	public void GetInput()
	{
		Velocity = GetVelocity();
	}
	
	private Vector2 GetVelocity()
	{
		return GetInputVector() * Speed;
	}
	
	private Vector2 GetInputVector()
	{
		return Input.GetVector(Keys.Left, Keys.Right, Keys.Up, Keys.Down);
	}

	public override void _PhysicsProcess(double delta)
	{
		GetInput();
		MoveAndSlide();
	}
}
