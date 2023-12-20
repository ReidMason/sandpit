local M = {}

M.setup = function()
	vim.keymap.set("n", "<C-b>", function()
		-- local fp = "~/Documents/cleanup-info.md"
		-- local x = vim.api.nvim_exec("glow " .. fp, true)
		-- print(x)
		-- local curr_file = vim.api.nvim_buf_get_name(0)

		local currWin = vim.api.nvim_get_current_win()
		vim.api.nvim_command("vsplit")
		local win = vim.api.nvim_get_current_win()
		local buf = vim.api.nvim_create_buf(true, true)
		vim.api.nvim_win_set_buf(win, buf)

		vim.api.nvim_set_current_win(currWin)

		-- 	local bufnr = vim.api.nvim_create_buf(false, true)
		-- 	local win_id = vim.api.nvim_open_win(bufnr, true, {
		-- 		relative = "editor",
		-- 		title = "MD preview",
		-- 		row = 3,
		-- 		col = 10,
		-- 		width = 10,
		-- 		height = 10,
		-- 		style = "minimal",
		-- 		border = "rounded",
		-- 	})
	end, {})
end

return M
