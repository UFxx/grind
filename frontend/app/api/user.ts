import { type User } from "~/types/user";
import { type Skill } from "~/types/skill";
import { type SuccessResponse } from "~/types/common";
import { type InviteCode, type AddInviteCode } from "~/types/inviteCode";

export default {
	// Fetchs
	fetchProfile : async () => await useRequest<SuccessResponse<User>>('/users/me/profile'),
	fetchCodes   : async () => await useRequest<SuccessResponse<InviteCode[]>>('/users/me/invite-codes'),
	fetchSkills  : async () => await useRequest<SuccessResponse<Skill[]>>('/users/me/skills-progress'),

	// Actions
	deleteInviteCode : async (id: string) => await useRequest<SuccessResponse<[]>>(`/users/me/invite-codes/${id}`, { method: 'DELETE' }),
	addInviteCode    : async (payload: AddInviteCode) =>
		await useRequest<SuccessResponse<[]>>('/users/me/invite-codes',
		{
			method: 'POST',
			body:
			{
				code     : payload.code,
				max_uses : Number(payload.maxUses)
			}
		}
	),
	deleteAccount: async () => await useRequest<SuccessResponse<[]>>('/users/me', { method: 'DELETE' })
};