using Godot;
using System;

public partial class Knife : Area2D
{
	private float _startAngle;
	private float _currentRotation = 0f;
	private float _sweepDuration = 0.5f; // Time to complete the sweep
	private float _sweepRange = Mathf.Pi / 2; // 90 degrees in radians
	private float _radius;
	private float _elapsed = 0f;
	private Player _player;
	
	public override void _Ready()
	{
		// Start at a random angle (or you can make it face mouse/enemy)
		_startAngle = 0;
		_currentRotation = _startAngle;
		
		// Set initial position
		UpdatePosition();
		
		// Connect body entered signal for damage
		BodyEntered += OnBodyEntered;
	}
	
	public override void _Process(double delta)
	{
		_elapsed += (float)delta;
		
		// Calculate sweep progress (0 to 1)
		float progress = Mathf.Clamp(_elapsed / _sweepDuration, 0f, 1f);
		
		// Smoothly interpolate the angle using ease-out
		float easedProgress = 1f - Mathf.Pow(1f - progress, 3f); // Cubic ease-out
		_currentRotation = _startAngle + (_sweepRange * easedProgress);
		
		// Update position in arc
		UpdatePosition();
		
		// Update sprite rotation to point outward
		Rotation = _currentRotation + Mathf.Pi / 2;
		
		// Despawn when sweep is complete
		if (progress >= 1f)
		{
			QueueFree();
		}
	}
	
	private void UpdatePosition()
	{
		// Calculate position in circular arc around player
		var playerPosition = _player.GlobalPosition;
		float x = playerPosition.X + Mathf.Cos(_currentRotation) * _radius;
		float y = playerPosition.Y + Mathf.Sin(_currentRotation) * _radius;
		GlobalPosition = new Vector2(x, y);
	}
	
	private void OnBodyEntered(Node2D body)
	{
		// Handle damage to enemies
		if (body.IsInGroup("enemies"))
		{
			GD.Print($"Knife hit {body.Name}!");
			// TODO: Apply damage to enemy
			// body.TakeDamage(10);
		}
	}
	
	// Helper method to set parameters from spawner
	public void Initialize(Player player, float startAngle, float xOffset = 0f, float yOffset = 0f, float radius = 5f, float duration = 0.5f)
	{
		_player = player;
		_startAngle = startAngle;
		GlobalPosition = new Vector2(xOffset, yOffset);
		_radius = radius;
		_sweepDuration = duration;
		_currentRotation = _startAngle;

		UpdatePosition();
		Rotation = _currentRotation + Mathf.Pi / 2;
	}
}
