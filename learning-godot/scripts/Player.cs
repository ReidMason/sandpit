using Godot;
using System;

public partial class Player : CharacterBody2D
{
	[Export]
	private Timer _attackTimer;
	[Export]
	private AnimatedSprite2D _animatedSprite;
	
	public const float Speed = 100.0f;
	private float attackSpeed = 2f;

	public override void _Ready()
	{
		Keys.SetupInputMap();
		this.ZIndex = (int)ZIndexes.Player;
		
		_attackTimer.Timeout += OnAttackTimerTimeout;
		_attackTimer.WaitTime = attackSpeed;
	}
	
	private void OnAttackTimerTimeout()
	{
		PerformAttack();
	}
	
	private void PerformAttack()
	{
		var knifeScene = GD.Load<PackedScene>("res://scenes/knife.tscn");
		var knife = knifeScene.Instantiate<Knife>();
		GetParent().AddChild(knife);

		float startAngle = -90 + (_animatedSprite.FlipH ? 180 : 0);
		knife.Initialize(this, startAngle, 5, -5);
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

	private void Animate()
	{
		// TODO: Should probably use a state machine for this
		if (Velocity.X > 0)
		{
			_animatedSprite.FlipH = false;
		}
		else if (Velocity.X < 0)
		{
			_animatedSprite.FlipH = true;
		}
		
		if (Velocity.X != 0)
		{
			_animatedSprite.Play("run");
		}
		else
		{
			_animatedSprite.Play("idle");
		}
	}

	public override void _PhysicsProcess(double _)
	{
		GetInput();
		Animate();
		MoveAndSlide();
	}
}
