import { type UserResponse } from "~/types/user"

export default {
	getProfile: async () => await useRequest<UserResponse>('/users/me/profile'),

	deleteInviteCodeByID: async (id: string) => await useRequest(`/users/me/invite-codes/${id}`, { method: 'DELETE' })
};