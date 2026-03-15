import WebApp from '@twa-dev/sdk';

export const useTelegramAuth = () =>
{
	const initData = useCookie('initData',
		{
			default: () => '',
			maxAge: 60 * 60,
		}
	);

	const token = useCookie('token',
		{
			default: () => '',
			maxAge: 60 * 60 * 24,
		}
	);

	const { auth: authApi } = useApi();

	const inviteCode = ref<string | undefined>('');

	const init = () =>
	{
		if (!process.client) return;

		if (!initData.value)
			initData.value = WebApp.initData;

		inviteCode.value = WebApp.initDataUnsafe?.start_param;

		WebApp.ready();
	};

	const auth = async () =>
	{
		try
		{
			init();

			if (token.value) return;

			const response = await authApi.telegramAuth(inviteCode.value, initData.value);

			token.value = response.data.token;
		}
		catch (err) { console.error(err); }
	};

	init();

	return {
		initData,
		token,
		inviteCode,
		auth,
	};
};