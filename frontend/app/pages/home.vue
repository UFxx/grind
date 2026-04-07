<script setup>
	const activityStore = useActivityStore();

	const fetchActivities = async () =>
	{
		try { await activityStore.fetchActivities(); }
		catch (err) { console.error(err); }
	}

	fetchActivities();
</script>

<template>
	<div class="home-page">
		<p
			class="home-page__title"

			v-motion-pop
			:duration="300"
		>
			List of activities
		</p>
		<TransitionGroup name="activity" tag="div" class="home-activities-list">
			<Activity
				v-for="(activity, idx) in activityStore.activities"
				:key="activity.id"
				:activity
			/>
		</TransitionGroup>
	</div>
</template>

<style lang='scss' scoped>
	.home-page
	{
		row-gap: 10px;

		display: flex;
		flex-direction: column;
	}

	.home-page__title
	{
		text-align: center;
		font-weight: 700;
	}

	.home-activities-list
	{
		position: relative;
		row-gap: 10px;
		margin-top: 10px;

		display: flex;
		flex-direction: column;
	}
</style>