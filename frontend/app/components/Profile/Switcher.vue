<script setup lang="ts">
	import { type UiTabsHeaderButton } from '~/types/ui/tabs';

	const userStore = useUserStore();

	const buttons: UiTabsHeaderButton[] =
	[
		{ title: 'My skills' },
		{ title: 'Invite codes' }
	];

	const activeTabIdx = ref(0);

	const fetchInviteCodes = async () =>
	{
		try { await userStore.fetchInviteCodes() }
		catch (err) { console.error(err); }
	}

	fetchInviteCodes();
</script>

<template>
	<UiTabs v-model="activeTabIdx" :items="buttons">
		<div class="profile-switcher-content__wr">
			<ProfileInviteCodes v-if="activeTabIdx === 1" />
			<ProfileSkills v-if="activeTabIdx === 0" />
		</div>
	</UiTabs>
</template>

<style lang="scss">
	.profile-switcher-content__wr
	{
		row-gap: 5px;
		padding: 3px;
		overflow: hidden;
		border-radius: 5px;
		background-color: $darkGray;

		display: flex;
		flex-direction: column;
	}
</style>