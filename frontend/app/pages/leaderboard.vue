<script setup lang="ts">
	import { type FormattedEntry, type FormattedSeason, type FormattedSeasonDetails, type SeasonDetails } from '~/types/leaderboard';

	const { leaderboard: leaderboardApi } = useApi();

	const seasonDetails = ref<FormattedSeasonDetails | null>(null);

	const hasSeason = computed(() => seasonDetails.value);
	const currentSeason = computed(() => seasonDetails.value?.season ?? null);
	const callToActionText = computed(() => seasonDetails.value?.cta ?? '');

	const entries = computed(() => seasonDetails.value?.entries ?? []);
	const myEntry = computed(() => seasonDetails.value?.me ?? null);

	const podiumEntries = computed(() => entries.value.slice(0, 3));
	const otherEntries = computed(() => entries.value.slice(3));

	const isMe = (entry: FormattedEntry) => myEntry.value ? entry.position === myEntry.value.position : false;

	const fetchSeasonDetails = async () =>
	{
		try {
			const { data: seasonsData } = await leaderboardApi.fetchSeasons();
			const seasons = leaderboardSeasonsSerializer(seasonsData);

			const firstSeason = seasons[0];
			if (!firstSeason) {
				return;
			}

			const { data: detail } = await leaderboardApi.fetchSeasonDetails(firstSeason.id);
			seasonDetails.value = leaderboardSeasonDetailsSerializer(detail);

		} catch (err) { console.error(err); }
	};

	await fetchSeasonDetails();
</script>

<template>
	<div class="leaderboard-page">
		<div
			v-if="!hasSeason"
			class="leaderboard__empty"
		>
			<p class="leaderboard__empty-text">There is no active season at the moment</p>
		</div>
		<div
			v-else
			class="leaderboard-wr"
		>
			<LeaderboardSeason
				v-if="currentSeason"
				:season="currentSeason"
				v-motion-fade
			/>
			<LeaderboardCTA
				:text="callToActionText"
				v-motion-fade
			/>
			<div
				v-if="!entries.length"
				class="leaderboard__empty"
			>
				<p class="leaderboard__empty-text">At the moment, no one has joined the season yet</p>
			</div>
			<LeaderboardPodium
				:entries="podiumEntries"
			/>
			<div
				class="leaderboard-entries"
			>
				<LeaderboardEntry
					v-for="(entry, idx) in otherEntries"
					:key="idx"
					:entry="entry"
					:isMe="isMe(entry)"
					v-motion-slide-left
					:duration="200"
					:delay="idx * 50"
				/>
			</div>
		</div>
	</div>
</template>

<style lang="scss">
	.leaderboard__empty {
		padding: 20px;

		color: $gray;
		font-size: 16px;
		text-align: center;
	}

	.leaderboard-wr {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 20px;
	}

	.leaderboard-entries {
		width: 100%;

		display: flex;
		flex-direction: column;
	}
</style>