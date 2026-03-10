import { type FetchContext } from 'ofetch';

export const useRequest = <T>(request: string, opts = {}) =>
{
	const config = useRuntimeConfig();

	const options: Object =
	{
		...opts,
		baseURL: config.public.api,
		onResponseError(e: FetchContext) { console.log(e) },
	};

	return $fetch<T>(request, options);
};
