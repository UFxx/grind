<script setup lang="ts">
	import { type UiTabsHeaderButton } from '~/types/ui/tabs';

	const userStore = useUserStore();

	const buttons: UiTabsHeaderButton[] =
	[
		{ title: 'Мои навыки' },
		{ title: 'Коды приглашений' }
	];

	const activeTabIdx = ref(1);

	const fetchInviteCodes = async () =>
	{
		try { await userStore.fetchInviteCodes() }
		catch (err) { console.error(err); }
	}

	fetchInviteCodes();
</script>

<template>
	<UiTabs v-model="activeTabIdx" :items="buttons">
		<TransitionGroup name="fade">
			<ProfileInviteCodes v-if="activeTabIdx === 1" />
			<div v-if="activeTabIdx === 0" />
		</TransitionGroup>
	</UiTabs>
</template>