import { type LoginResponse } from "~/types/auth"

export default {
	telegramAuth: async (
		inviteCode : string | undefined,
		initData   : string
	) =>
		await useRequest<LoginResponse>('/auth/telegram',
			{
				method  : 'POST',
				body    : { invite_code: inviteCode },
				headers : { Authorization: `tma ${initData}` }
			}
		)
}