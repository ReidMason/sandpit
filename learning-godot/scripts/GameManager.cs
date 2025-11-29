using Godot;
using System;

public partial class GameManager : Node
{
	public static GameManager Instance { get; private set; }

	[Export]
	private Label _scoreLabel; 
	
	private int Score { get; set; } = 0;
	
	public override void _EnterTree()
	{
		Instance = this;
		AddScore(0);
	}
	
	public void AddScore(int value)
	{
		Score += value;
		_scoreLabel.Text = $"Score: {Score}";
	}
	
	public int GetScore()
	{
		return Score;
	}
}
