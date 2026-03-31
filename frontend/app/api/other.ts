import { type Feedback } from "~/types/feedback"
import { type SuccessResponse } from "~/types/common"

export default {
	sendFeedback: async (payload: Feedback) => await useRequest<SuccessResponse<[]>>('/app/feedback',
		{
			method: 'POST',
			body: payload
		}
	)
}