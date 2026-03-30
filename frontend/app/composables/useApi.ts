import auth from "~/api/auth";
import user from "~/api/user";
import activity from "~/api/activity";

export const useApi = () =>
{
	return {
		auth,
		user,
		activity
	};
}