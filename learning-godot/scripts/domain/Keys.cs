using System;
using Godot;

namespace scripts.domain;

public static class Keys
{
	public const string Left = "left";
	public const string Right = "right";
	public const string Up = "up";
	public const string Down = "down";
	
	private static bool InputMapInitialised { get; set; } = false;
	
	public static void SetupInputMap()
	{
		if (InputMapInitialised) return;
		
		var keyEventA = new InputEventKey();
		keyEventA.Keycode = Key.A;
		InputMap.AddAction(Left);
		InputMap.ActionAddEvent(Left, keyEventA);
		
		InputMapInitialised = true;
	}
}
