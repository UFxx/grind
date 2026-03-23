import { type PopupName } from "~/types/ui/popups";

import AddInviteCode from "~/components/popups/AddInviteCode.vue";

const popups = { 'AddInviteCode': AddInviteCode }

export const usePopupsRegistry = (name: PopupName | null) =>
{
	if (!name) return
	return popups[name];
}