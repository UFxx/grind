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
	const route = useRoute();

	const inviteCode = ref<string | undefined>('');

	const init = () =>
	{
		if (!process.client) return;

		if (!initData.value)
			initData.value = WebApp.initData;

		if (WebApp.initDataUnsafe?.start_param)
			inviteCode.value = WebApp.initDataUnsafe?.start_param;
		else
		{
			const queryValue = route.query.inviteCode;
			inviteCode.value = Array.isArray(queryValue)
				? queryValue[0] ?? undefined
				: queryValue ?? undefined;
		}

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
		catch (err)
		{
			console.error(err);
			navigateTo('/auth');
		}
	};

	return {
		initData,
		token,
		inviteCode,
		auth,
	};
};