<script setup lang="ts">
	import { type AddActivity } from '~/types/activity';
	import { type UiMiniSwitcherItem } from '~/types/ui/miniSwitcher';

	const activityStore = useActivityStore();
	const { addToast }  = useToastsStore();
	const { togglePopup } = usePopupsStore();

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

	const closePopup = () => togglePopup('AddActivity', false);

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
			{
				activityDescription.value = '';
				addToast('success', 'Activity added successfully');
			}
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
			<UiMiniSwitcher
				:items="miniSwitcherItems"
				v-model="currentItem"
			/>
			<IconsClose
				class="add-activity__close"
				@click="closePopup"
			/>
		</div>

		<div class="add-activity__title">
			<p class="add-activity__header-title">Add activity</p>
		</div>

		<div class="add-activity__content">
			<PopupsAddActivityAi
				v-if="currentItem.id === 'ai'"
				v-model="activityDescription"
			/>
		</div>
		<UiButton
			color="white"
			:fullWidth="true"
			:disabled="isLoading"
			class="add-activity__add-button"
			@click="createActivity"
		>
				<span v-if="!isLoading">ADD</span>
				<UiLoader v-else />
		</UiButton>
	</div>
</template>

<style lang='scss' scoped>
	.add-activity
	{
		width: 100%;
		padding: 10px;
		border-radius: 10px;
		background-color: $primary;
		box-shadow: 0 0 4px 2px $darkGray;
	}

	.add-activity__close
	{
		opacity: 0.5;
		cursor: pointer;
		margin-bottom: 5px;

		display: flex;
		justify-content: flex-end;

		@include tr(.3, opacity);

		&:hover { opacity: 1; }
	}

	.add-activity__header
	{
		margin-bottom: 5px;

		display: flex;
		align-items: center;
		justify-content: space-between;
	}

	.add-activity__title
	{
		padding-bottom: 5px;
		border-bottom: 1px solid $darkGray;
	}

	.add-activity__header-title { font-weight: 700; }

	.add-activity__content { padding: 10px 0; }

	.add-activity__add-button
	{
		text-transform: uppercase;
		font-weight: 700;
	}
</style>