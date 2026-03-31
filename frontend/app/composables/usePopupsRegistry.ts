import { type PopupName } from "~/types/ui/popups";

import UserInfo from "~/components/popups/UserInfo.vue";
import Feedback from "~/components/popups/Feedback.vue";
import AddInviteCode from "~/components/popups/AddInviteCode.vue";

const popups =
{
	'AddInviteCode' : AddInviteCode,
	'UserInfo'      : UserInfo,
	'Feedback'      : Feedback
}

export const usePopupsRegistry = (name: PopupName | null) =>
{
	if (!name) return;

	return popups[name];
}