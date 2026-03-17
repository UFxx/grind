<script setup lang="ts">
	import { type UiTabsHeaderButton } from '~/types/ui/tabs';
	import { type FormattedInviteCode } from '~/types/inviteCode';

	const buttons: UiTabsHeaderButton[] =
	[
		{ title: 'Мои навыки' },
		{ title: 'Коды приглашений' }
	];

	const inviteCodes: FormattedInviteCode[] =
	[
		{
			id: "550e8400-e29b-41d4-a716-446655440001",
			code: "INVITE67890",
			uses: 0,
			maxUses: 3,
			createdAt: "2024-01-02T15:30:00Z"
		},
		{
			id: "550e8400-e29b-41d4-a716-446655440001",
			code: "INVITE67890",
			uses: 0,
			maxUses: 3,
			createdAt: "2024-01-02T15:30:00Z"
		}
	]

	const activeTabIdx = ref(0);
</script>

<template>
	<UiTabs v-model="activeTabIdx" :items="buttons">
		<InviteCodes :invite-codes />
	</UiTabs>
</template>

<style lang='scss'></style>