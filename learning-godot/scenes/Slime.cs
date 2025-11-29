using Godot;
using System;

public partial class Slime : CharacterBody2D
{
	[Export]
	private ProgressBar _healthBar;
	
	[Export]
	private Area2D _damageArea;
	
	private double _health = 100;
	private double _maxHealth = 100;
	private float _speed = 100;
	private bool _collidingWithPlayer = false;
	
	public override void _Ready()
	{
		AddToGroup(CollisionGroups.Enemies);
		_healthBar.MaxValue = _maxHealth;
		_healthBar.MinValue = 0;
		_healthBar.Value = _health;
		
		_damageArea.BodyEntered += OnDamageAreaBodyEntered;
		_damageArea.BodyExited += OnDamageAreaBodyExited;
	}
	
	private void OnDamageAreaBodyEntered(Node2D body)
	{
		if (body is Player)
		{
			_collidingWithPlayer = true;
		}
	}
	
	private void OnDamageAreaBodyExited(Node2D body)
	{
		if (body is Player)
		{
			GD.Print("Exited the damage area");
			_collidingWithPlayer = false;
		}
	}

	public override void _PhysicsProcess(double delta)
	{
		var player = GameManager.Instance.Player;
		if (player is null || _collidingWithPlayer) return;

		var direction = GlobalPosition.DirectionTo(player.GlobalPosition);
		Velocity = direction * _speed;
		MoveAndCollide(Velocity * (float)delta);
	}
	
	public void Damage(double amount)
	{
		_health = Math.Clamp(_health - amount, 0, _maxHealth);
		_healthBar.Value = _health;
		
		if (_health <= 0) Die();
	}
	
	private void Die()
	{
		QueueFree();
	}
}
