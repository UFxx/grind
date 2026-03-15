import { type FetchContext } from 'ofetch';

export const useRequest = <T>(request: string, opts = {}) =>
{
	const config    = useRuntimeConfig();
	const { token } = useTelegramAuth();

	const options: Object =
	{
		baseURL: config.public.api,
		headers: { Authorization: token.value !== '' ? token.value : null },
		...opts,

		onResponseError(e: FetchContext) { console.log(e) },
	};

	return $fetch<T>(request, options);
};
