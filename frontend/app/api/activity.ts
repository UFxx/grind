import { type SuccessResponse } from "~/types/common"
import { type ActivitiesResponse, type AddActivity } from "~/types/activity"

export default {
	fetchActivities: async (limit = 10, page = 1) =>
		await useRequest<SuccessResponse<ActivitiesResponse>>(`/users/me/activities?limit=${limit}&page=${page}`),

	// Actions
	createActivity: async (payload: AddActivity) => await useRequest<SuccessResponse<[]>>('/users/me/activities',
		{
			method: 'POST',
			body:
			{
				mode        : payload.mode,
				description : payload.description
			}
		}
	)
}