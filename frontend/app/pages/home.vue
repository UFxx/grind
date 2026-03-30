<script setup>
	const activityStore = useActivityStore();

	const fetchActivities = async () =>
	{
		try { await activityStore.fetchActivities(); }
		catch (err) { console.error(err); }
	}

	await fetchActivities();
</script>

<template>
	<div class="home-page">
		<AddActivity />

		<div class="home-activities-list">
			<Activity
				v-for="activity in activityStore.activities"
				:key="activity.id"
				:activity
			/>
		</div>
	</div>
</template>

<style lang='scss' scoped>
	.home-page
	{
		row-gap: 10px;

		display: flex;
		flex-direction: column;
	}

	.home-activities-list
	{
		row-gap: 10px;
		margin-top: 10px;

		display: flex;
		flex-direction: column;
	}
</style>