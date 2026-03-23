import { type FetchError } from '~/types/common';

type NormalizedErrors = Record<string, string>;

export const useFetchErrors = (err: unknown) =>
{
	const errors     = ref<NormalizedErrors>({});
	const statusCode = ref<number | undefined>(undefined);

	console.error(err);

	if (err && typeof err === 'object' && 'statusCode' in err)
	{
		const error = err as FetchError;

		statusCode.value = error.statusCode;

		if (error.data?.errors && typeof error.data.errors === 'object')
			errors.value = error.data.errors;
	}

	return {
		errors,
		statusCode,
	};
};