import { type FormattedActivity, type AddActivity } from "~/types/activity";

const { activity } = useApi();

export const useActivityStore = defineStore('activity', () =>
	{
		// Data
		const activities = ref<FormattedActivity[]>([]);

		// Inner functions
		const setActivities = (value: FormattedActivity[]) => activities.value = value;

		// Fetchs + setters
		const fetchActivities = async () =>
		{
			const response = await activity.fetchActivities()
			setActivities(activitySerialilzer(response.data.items));
		};

		// Actions
		const createActivity = async (payload: AddActivity) =>
		{
			const response = await activity.createActivity(payload);
			await fetchActivities();

			return response;
		};

		return {
			// Data
			activities,

			// Fetchs + setters
			fetchActivities,

			// Actions
			createActivity
		}
	}
);