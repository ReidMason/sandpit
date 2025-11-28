using Godot;
using System;

public partial class Knife : Area2D
{
	private float _startAngle;
	private float _currentRotation = 0f;
	private float _sweepDuration = 0.5f; // Seconds
	private float _sweepRange = Mathf.Pi / 2; // 90 degrees in radians
	private float _radius;
	private float _elapsed = 0f;
	private Player _player;
	
	public override void _Ready()
	{
		_startAngle = 0;
		_currentRotation = _startAngle;
		Visible = false;
		
		BodyEntered += OnBodyEntered;
	}
	
	public override void _Process(double delta)
	{
		_elapsed += (float)delta;
		
		float progress = Mathf.Clamp(_elapsed / _sweepDuration, 0f, 1f);
		
		float easedProgress = 1f - Mathf.Pow(1f - progress, 3f); // Cubic ease-out
		_currentRotation = _startAngle + (_sweepRange * easedProgress);
		
		UpdatePosition();
		
		Rotation = _currentRotation + Mathf.Pi / 2;
		
		if (progress >= 1f) QueueFree();
	}
	
	private void UpdatePosition()
	{
		var playerPosition = _player.GlobalPosition;
		float x = playerPosition.X + Mathf.Cos(_currentRotation) * _radius;
		float y = playerPosition.Y + Mathf.Sin(_currentRotation) * _radius;
		GlobalPosition = new Vector2(x, y);
	}
	
	private void OnBodyEntered(Node2D body)
	{
		GD.Print($"Knife hit {body.Name}!");
		if (body.IsInGroup("enemies"))
		{
			GD.Print($"Knife hit {body.Name}!");
		}
	}
	
	public void Initialize(Player player, float startAngle, float xOffset = 0f, float yOffset = 0f, float radius = 5f, float duration = 0.5f)
	{
		_player = player;
		_startAngle = startAngle;
		GlobalPosition = new Vector2(xOffset, yOffset);
		_radius = radius;
		_sweepDuration = duration;
		_currentRotation = _startAngle;

		Rotation = _currentRotation + Mathf.Pi / 2;
		UpdatePosition();
		Visible = true;
	}
}
