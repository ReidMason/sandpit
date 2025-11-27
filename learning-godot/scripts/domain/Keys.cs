using System;
using System.Collections.Generic;
using Godot;

namespace scripts.domain;

public static class Keys
{
	public const string Left = "left";
	public const string Right = "right";
	public const string Up = "up";
	public const string Down = "down";
	
	private static bool InputMapInitialised { get; set; } = false;
	
	private static readonly Dictionary<string, Key[]> ActionKeyMap = new()
	{
		{ Left, new[] { Key.A, Key.Left } },
		{ Right, new[] { Key.D, Key.Right } },
		{ Up, new[] { Key.W, Key.Up } },
		{ Down, new[] { Key.S, Key.Down } }
	};
	
	public static void SetupInputMap()
	{
		if (InputMapInitialised) return;
		
		foreach (var (action, keys) in ActionKeyMap)
			AddKeysToActionKeyMap(action, keys);
			
		InputMapInitialised = true;
	}
	
	private static void AddKeysToActionKeyMap(string action, Key[] keys)
	{
		InputMap.AddAction(action);
			
		foreach (var key in keys)
			AddKeyToActionKeyMap(action, key);
	}
	
	private static void AddKeyToActionKeyMap(string action, Key key)
	{
		var keyEvent = new InputEventKey();
		keyEvent.Keycode = key;
		InputMap.ActionAddEvent(action, keyEvent);	
	}
}
