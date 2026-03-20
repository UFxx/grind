import { type User } from "~/types/user"
import { type InviteCode } from "~/types/inviteCode"
import { type SuccessResponse } from "~/types/common";

export default {
	fetchProfile     : async () => await useRequest<SuccessResponse<User>>('/users/me/profile'),
	deleteInviteCode : async (id: string) => await useRequest<SuccessResponse<[]>>(`/users/me/invite-codes/${id}`, { method: 'DELETE' }),
	fetchCodes       : async () => await useRequest<SuccessResponse<InviteCode[]>>('/users/me/invite-codes'),
	addCode          : async (
		code    : string,
		maxUses : number
	) =>
		await useRequest<SuccessResponse<[]>>('/users/me/invite-codes',
		{
			method: 'POST',
			body:
			{
				code,
				max_uses: maxUses
			}
		}
	)
};