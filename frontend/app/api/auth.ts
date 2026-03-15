export default {
	telegramAuth: async () =>
	{
		return useRequest('/auth/telegram',
			{
				method: 'POST',
				// body: { invite_code:  }
			}
		)
	}
}