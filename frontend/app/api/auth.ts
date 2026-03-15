import { type ILoginResponse } from "~/types/login"

export default {
	telegramAuth: async (
		inviteCode: string | undefined,
		initData: string
	) =>
		await useRequest<ILoginResponse>('/auth/telegram',
			{
				method: 'POST',
				body: { invite_code: inviteCode },
				headers: { Authorization: `tma ${initData}` }
			}
		)
}