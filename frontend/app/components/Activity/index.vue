<script setup lang="ts">
	import { type FormattedActivity } from '~/types/activity';

	const props = defineProps<{ activity: FormattedActivity }>();

	const { deleteActivity } = useActivityStore();
	const { addToast }       = useToastsStore();

	const isLoading = ref(false);

	const deleteActivityHandler = async () =>
	{
		isLoading.value = true;

		try
		{
			const response = await deleteActivity(props.activity.id);

			if (!response.data.length)
				addToast('success', 'Activity successfully deleted');
		}
		catch(err)
		{
			console.error(err);
			addToast('error', 'Error deleting activity');
		}
		finally { isLoading.value = false; }
	};
</script>

<template>
	<div class="activity">
		<div class="activity__header">
			<div class="activity__header-tags-wr">
				<div class="activity__header-tags">
					<ActivityTag
						v-for="(tag, idx) in activity.tags"
						:key="idx"
						:tag
					>
						{{ tag.name }}
					</ActivityTag>
				</div>

				<div class="activity__header-delete" @click="deleteActivityHandler">
					<IconsClose />
				</div>
			</div>
			<div class="activity__header-title-wr">
				<span class="activity__header-title">
					{{ activity.description }}
				</span>
				<span class="activity__header-date">
					{{ formatDate(activity.createdAt) }}
				</span>
			</div>
		</div>

		<div class="activity__rewards">
			<div
				v-for="reward in activity.rewards"
				class="activity__reward"
			>
				<span class="activity__reward-skill-name">{{ reward.skillName }}</span>
				<span class="activity__reward-xp-amount">+{{ reward.xpAmount }}xp</span>
			</div>
		</div>
	</div>
</template>

<style lang='scss' scoped>
	.activity
	{
		padding: 10px;
		border-radius: 10px;
		border: 2px solid rgba($gray, 0.25);
	}

	.activity__header
	{
		row-gap: 10px;

		display: flex;
		flex-direction: column;
	}

	.activity__header-delete
	{
		cursor: pointer;
		color: rgba($red, 0.3);

		@include tr(.3, color);

		&:hover { color: $red; }
	}

	.activity__header-tags-wr
	{
		display: flex;
		align-items: center;
		justify-content: space-between;
	}

	.activity__header-tags
	{
		column-gap: 5px;

		display: flex;
	}

	.activity__header-title-wr
	{
		column-gap: 10px;
		padding-bottom: 10px;
		border-bottom: 1px solid $darkGray;

		display: flex;
		align-items: center;
		justify-content: space-between;
	}

	.activity__header-title
	{
		max-width: 80%;
		font-weight: 700;
		word-break: break-word;
	}

	.activity__header-date
	{
		color: $gray;
		font-size: 12px;
		font-weight: 500;
	}

	.activity__rewards
	{
		column-gap: 20px;
		padding: 10px 10px 0 10px;

		display: flex;
	}

	.activity__reward
	{
		display: flex;
		align-items: center;
		flex-direction: column;
	}

	.activity__reward-skill-name { font-size: 12px; }
	.activity__reward-xp-amount
	{
		color: $green;
		font-size: 10px;
		font-weight: 300;
	}
</style>