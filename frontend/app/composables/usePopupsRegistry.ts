import { type PopupName } from "~/types/ui/popups";

import UserInfo from "~/components/popups/UserInfo.vue";
import Feedback from "~/components/popups/Feedback.vue";
import AddActivity from '~/components/popups/AddActivity/index.vue';
import AddInviteCode from "~/components/popups/AddInviteCode.vue";

const popups =
{
	'UserInfo'      : UserInfo,
	'Feedback'      : Feedback,
	'AddActivity'   : AddActivity,
	'AddInviteCode' : AddInviteCode
}

export const usePopupsRegistry = (name: PopupName | null) =>
{
	if (!name) return;

	return popups[name];
}