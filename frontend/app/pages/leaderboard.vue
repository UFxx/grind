<script setup lang="ts">
	import { type FormattedSeason } from '~/types/leaderboard';

	const { leaderboard: leaderboardApi } = useApi();

	const seasons = ref<FormattedSeason[]>([]);

	const fetchSeasons = async () =>
	{
		try
		{
			const { data } = await leaderboardApi.fetchSeasons();
			seasons.value = leaderboardSeasonsSerializer(data);
		}
		catch (err) { console.error(err); }
	};

	await fetchSeasons();
</script>

<template>
	<div class="leaderboard-page">
		<LeaderboardSeason
			v-if="seasons.length"
			:season="seasons[0]"
		/>
	</div>
</template>

<style lang="scss">
</style>