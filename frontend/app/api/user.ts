import { type UserResponse } from "~/types/user"

export default {
	getProfile: async () => await useRequest<UserResponse>('/users/me/profile')
};