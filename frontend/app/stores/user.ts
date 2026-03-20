import { type FormattedInviteCode } from "~/types/inviteCode";
import { type FormattedUser } from "~/types/user";

const { user } = useApi();

export const useUserStore = defineStore('user', () =>
	{
		const userData        = ref<FormattedUser | null>(null);
		const userInviteCodes = ref<FormattedInviteCode[]>([]);

		const setInviteCodes = (inviteCodes: FormattedInviteCode[]) => userInviteCodes.value = inviteCodes;
		const setUserData    = (data: FormattedUser) => userData.value = data;

		const fetchInviteCodes = async () =>
		{
			const response = await user.fetchCodes();
			setInviteCodes(inviteCodeSerializer(response.data));
		};

		const addInviteCode = async (code: string, maxUses: number) =>
		{
			await user.addCode(code, maxUses);
			await fetchInviteCodes();
		};

		const deleteInviteCode = async (id: string) =>
		{
			await user.deleteInviteCode(id);
			fetchInviteCodes();
		};

		const fetchProfile = async () =>
		{
			const { data } = await user.fetchProfile();
			setUserData(userSerializer(data));
		};

		return {
			userData,
			userInviteCodes,

			fetchProfile,
			addInviteCode,
			deleteInviteCode,
			fetchInviteCodes
		};
	}
)