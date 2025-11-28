using Godot;
using System;

public partial class Slime : CharacterBody2D
{
	public override void _Ready()
	{
		AddToGroup(CollisionGroups.Enemies);
	}
}
