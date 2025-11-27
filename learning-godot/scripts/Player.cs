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
		Vector2 inputDirection = Input.GetVector(Keys.Left, Keys.Right, Keys.Up, Keys.Down);
		Velocity = inputDirection * Speed;
		
		var actionEvents = InputMap.HasAction(Keys.Left);
		GD.Print(actionEvents);
	}

	public override void _PhysicsProcess(double delta)
	{
		GetInput();
		MoveAndSlide();
	}
}
