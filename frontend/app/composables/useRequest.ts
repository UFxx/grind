export const useRequest = <T>(request: string, opts = {}) =>
{
	const config    = useRuntimeConfig();
	const { token } = useTelegramAuth();

	const options: Object =
	{
		baseURL: config.public.api,
		headers:
		{
			Authorization: token.value !== ''
				? `Bearer ${token.value}`
				: null
		},
		...opts,
	};

	return $fetch<T>(request, options);
};