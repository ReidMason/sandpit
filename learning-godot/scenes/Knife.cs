using Godot;
using System;

public record Offset {
	public float X { get; set; }
	public float Y { get; set; }
}

public partial class Knife : Area2D
{
	private float _startAngle = 0;
	private float _currentRotation = 0;
	private float _sweepDuration = 0.5f; // Seconds
	private float _sweepRange = Mathf.Pi / 2; // 90 degrees in radians
	private float _radius;
	private float _elapsed = 0f;
	private Player _player;
	private Offset _offset;
	
	public override void _Ready()
	{
		Visible = false;
		BodyEntered += OnBodyEntered;
	}
	
	public override void _Process(double delta)
	{
		_elapsed += (float)delta;
		
		var progress = Mathf.Clamp(_elapsed / _sweepDuration, 0f, 1f);
		
		var easedProgress = 1 - Mathf.Pow(1 - progress, 3); // Cubic ease-out
		_currentRotation = _startAngle + (_sweepRange * easedProgress);
		
		UpdatePosition();
			
		if (progress >= 1f) QueueFree();
	}
	
	private void UpdatePosition()
	{
		Rotation = _currentRotation + _sweepRange;

		var playerPosition = _player.GlobalPosition;
		float x = playerPosition.X + _offset.X;
		float y = playerPosition.Y + _offset.Y;
		GlobalPosition = new Vector2(x, y);
	}
	
	private void OnBodyEntered(Node2D body)
	{
		if (body.IsInGroup(CollisionGroups.Enemies))
		{
			body.QueueFree();
			GameManager.Instance.AddScore(1);
		}
	}
	
	public void Initialize(Player player, float startAngle, float xOffset = 0f, float yOffset = 0f, float radius = 5f, float duration = 0.5f)
	{
		_player = player;
		_startAngle = startAngle;
		_radius = radius;
		_sweepDuration = duration;
		_offset = new Offset { X = xOffset, Y = yOffset };
		_currentRotation = _startAngle;

		Rotation = _currentRotation + Mathf.Pi / 2;
		UpdatePosition();
		Visible = true;
	}
}
