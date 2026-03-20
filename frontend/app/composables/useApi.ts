import auth from "~/api/auth";
import user from "~/api/user";

export const useApi = () =>
{
	return {
		auth,
		user,
	};
}