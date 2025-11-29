using Godot;
using System;

public partial class Slime : CharacterBody2D
{
	[Export]
	private ProgressBar _healthBar;
	
	private double _maxHealth = 100;
	private double _health = 100;
	
	public override void _Ready()
	{
		AddToGroup(CollisionGroups.Enemies);
		_healthBar.MaxValue = _maxHealth;
		_healthBar.MinValue = 0;
		_healthBar.Value = _health; 
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
