using Godot;
using System;

public record Offset {
	public float X { get; set; }
	public float Y { get; set; }
}

public partial class Knife : Area2D
{
	private float _startAngle = 90;
	private float _currentRotation = 0;
	private float _sweepDuration = 0.5f; // Seconds
	private float _sweepAngle = Mathf.Pi; // 180 degrees
	private float _radius;
	private float _elapsed = 0f;
	private Player _player;
	private Offset _offset;
	private bool _flipped;
	private double _damage = 1;
	
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
		
		_currentRotation = _startAngle + (_sweepAngle * easedProgress * (_flipped ? -1 : 1));
		
		UpdatePosition();
			
		if (progress >= 1f) QueueFree();
	}
	
	private void UpdatePosition()
	{
		Rotation = _currentRotation + _sweepAngle;

		var playerPosition = _player.GlobalPosition;
		float x = playerPosition.X + _offset.X;
		float y = playerPosition.Y + _offset.Y;
		GlobalPosition = new Vector2(x, y);
	}
	
	private void OnBodyEntered(Node2D body)
	{
		if (body.IsInGroup(CollisionGroups.Enemies) && body is Slime slime)
		{
			slime.Damage(CalculateDamage());
		}
	}
	
	private double CalculateDamage()
	{
		return _damage + GameManager.Instance.GetScore();
	}
	
	public void Initialize(Player player, bool flipped, float xOffset = 0f, float yOffset = 0f, float radius = 5f, float duration = 0.5f)
	{
		_player = player;
		_flipped = flipped;
		_radius = radius;
		_sweepDuration = duration;
		_offset = new Offset { X = xOffset, Y = yOffset };
		_currentRotation = _startAngle;

		Rotation = _currentRotation + Mathf.Pi / 2;
		UpdatePosition();
		Visible = true;
	}
}
