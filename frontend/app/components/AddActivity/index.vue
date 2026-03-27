<script setup lang="ts">
	import { type UiMiniSwitcherItem } from '~/types/ui/miniSwitcher';

	const activityDescription = ref('');

	const miniSwitcherItems = ref<[UiMiniSwitcherItem, UiMiniSwitcherItem]>(
		[
			{
				id: 'AI',
				label: 'AI',
			},
			{
				id: 'Manual',
				label: 'Manual',
			}
		]
	);

	const currentItem = ref(miniSwitcherItems.value[0]);
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
				v-if="currentItem.id === 'AI'"
				v-model="activityDescription"
			/>
		</div>
	</div>

	<UiButton
		color="white"
		:fullWidth="true"
		class="add-activity__add-button"
	>
		Add
	</UiButton>
</template>

<style lang='scss' scoped>
	.add-activity { padding: 10px; }

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