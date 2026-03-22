import { type PopupName } from "~/types/ui/popups";

export const usePopupsStore = defineStore('popups', () =>
	{
		const activePopup = ref<PopupName | null>(null);

		const togglePopup = (name: PopupName, value: boolean) =>
			value
				? activePopup.value = name
				: activePopup.value = null;

		return {
			activePopup,

			togglePopup
		}
	}
)