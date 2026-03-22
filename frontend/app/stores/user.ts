import { type FormattedUser } from "~/types/user";
import { type FormattedSkill } from "~/types/skill";
import { type FormattedInviteCode, type AddInviteCode } from "~/types/inviteCode";

const { user } = useApi();

export const useUserStore = defineStore('user', () =>
	{
		// Data
		const userData        = ref<FormattedUser | null>(null);
		const userInviteCodes = ref<FormattedInviteCode[]>([]);
		const userSkills      = ref<FormattedSkill[]>([]);

		// Inner functions
		const setInviteCodes = (inviteCodes: FormattedInviteCode[]) => userInviteCodes.value = inviteCodes;
		const setUserData    = (data: FormattedUser) => userData.value = data;
		const setUserSkills  = (skills: FormattedSkill[]) => userSkills.value = skills;

		// Actions
		const addInviteCode = async (payload: AddInviteCode ) =>
		{
			const response = await user.addInviteCode(payload);
			await fetchInviteCodes();

			return response;
		};

		const deleteInviteCode = async (id: string) =>
		{
			await user.deleteInviteCode(id);
			await fetchInviteCodes();
		};

		// Fetchs + setters
		const fetchProfile = async () =>
		{
			const { data } = await user.fetchProfile();
			setUserData(userSerializer(data));
		};

		const fetchInviteCodes = async () =>
		{
			const response = await user.fetchCodes();
			setInviteCodes(inviteCodeSerializer(response.data));
		};

		const fetchUserSkills = async () =>
		{
			const response = await user.fetchSkills();
			setUserSkills(skillSerializer(response.data));
		};

		return {
			// Data
			userData,
			userSkills,
			userInviteCodes,

			// !-- temp --!
			setUserSkills,

			// Actions
			addInviteCode,
			deleteInviteCode,

			// Fetchs + setters
			fetchProfile,
			fetchUserSkills,
			fetchInviteCodes,
		};
	}
)