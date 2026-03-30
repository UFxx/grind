<script setup lang="ts">
	import { type AddActivity } from '~/types/activity';
	import { type UiMiniSwitcherItem } from '~/types/ui/miniSwitcher';

	const activityStore = useActivityStore();
	const { addToast }  = useToastsStore();

	const isLoading           = ref(false);
	const activityDescription = ref('');
	const miniSwitcherItems   = ref<[UiMiniSwitcherItem, UiMiniSwitcherItem]>(
		[
			{
				id    : 'ai',
				label : 'AI'
			},
			{
				id    : 'manually',
				label : 'Manual'
			}
		]
	);

	const currentItem = ref<UiMiniSwitcherItem>(miniSwitcherItems.value[0]);

	const createActivity = async () =>
	{
		const payload: AddActivity =
		{
			mode        : currentItem.value.id,
			description : activityDescription.value
		}

		isLoading.value = true;

		try
		{
			const response = await activityStore.createActivity(payload);

			if (!response.data.length)
				addToast('success', 'Activity added successfully');
		}
		catch(err)
		{
			console.error(err);
			addToast('error', 'Failed to add activity')
		}
		finally { isLoading.value = false; }
	};
</script>

<template>
	<div class="add-activity">
		<div class="add-activity__header">
			<p class="add-activity__header-title">Add activity</p>
			<UiMiniSwitcher
				:items="miniSwitcherItems"
				v-model="currentItem"
			/>
		</div>

		<div class="add-activity__content">
			<AddActivityAi
				v-if="currentItem.id === 'ai'"
				v-model="activityDescription"
			/>
		</div>
	</div>

	<UiButton
		color="white"
		:fullWidth="true"
		:disabled="isLoading"
		class="add-activity__add-button"
		@click="createActivity"
	>
		Add
	</UiButton>
</template>

<style lang='scss' scoped>
	.add-activity { padding: 10px 10px 0 10px; }

	.add-activity__header
	{
		padding-bottom: 10px;
		border-bottom: 1px solid $darkGray;

		display: flex;
		align-items: center;
		justify-content: space-between;
	}

	.add-activity__header-title { font-weight: 700; }

	.add-activity__content { padding: 10px 0; }

	.add-activity__add-button
	{
		text-transform: uppercase;
		font-weight: 700;
	}
</style>