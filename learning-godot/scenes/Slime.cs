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
	private float _speed = 50;
	
	public override void _Ready()
	{
		AddToGroup(CollisionGroups.Enemies);
		_healthBar.MaxValue = _maxHealth;
		_healthBar.MinValue = 0;
		_healthBar.Value = _health;
	}

	public override void _PhysicsProcess(double delta)
	{
		var player = GameManager.Instance.Player;
		if (player is null) return;

		var direction = GlobalPosition.DirectionTo(player.GlobalPosition);
		Velocity = direction * _speed;
		
		var collision = MoveAndCollide(Velocity * (float)delta);
		if (collision is not null && collision.GetCollider() is Player)
		{
			GD.Print("There was a collision");
		} else {
			GD.Print("No more collision");
		}
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
