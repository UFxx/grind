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
		<AddActivity
			v-motion-slide-top
			:duration="300"
			:delay="0"
		/>

		<div class="home-activities-list">
			<Activity
				v-for="(activity, idx) in activityStore.activities"
				:key="activity.id"
				:activity

				v-motion-fade
				:duration="150"
				:delay="(idx + 1) * 50"
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