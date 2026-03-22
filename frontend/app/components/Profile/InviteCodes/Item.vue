<script setup lang="ts">
	import { type FormattedInviteCode } from '~/types/inviteCode';

	const props = defineProps<{
		inviteCode: FormattedInviteCode,
		loadingInviteCodeIds: string[]
	}>();

	const emit = defineEmits(['deleteCode']);
1
	const deleteCode = (id: string) => emit('deleteCode', id);
	const copyCode   = (inviteCode: FormattedInviteCode) => window.navigator.clipboard.writeText(inviteCode.code);

	const isInviteCodeLoading = computed(() => props.loadingInviteCodeIds.includes(props.inviteCode.id));
</script>

<template>
	<div
		class="invite-code"
		:class="{ 'invite-code--loading': isInviteCodeLoading }"
	>
		<div class="invite-code__info">
			<div
				@click="copyCode(inviteCode)"
				class="invite-code__code"
			>
				<span class="invite-code__code-icon"><IconsCopy /></span>
				<p class="invite-code__code-label">{{ inviteCode.code }}</p>
			</div>
			<div class="invite-code__uses">
				<span class="invite-code__uses-icon"><IconsUsers /></span>
				<span class="invite-code__uses-label">{{ inviteCode.uses }}/{{ inviteCode.maxUses }}</span>
			</div>
		</div>
		<button
			@click="deleteCode(inviteCode.id)"
			class="invite-code__delete"
		>
			<IconsDelete />
		</button>
	</div>
</template>

<style lang='scss' scoped>
	.invite-code
	{
		row-gap: 5px;
		padding: 10px;
		font-weight: 500;
		border-radius: 5px;
		background-color: $black;

		display: flex;
		justify-content: space-between;

		@include tr(.3, transform, background-color, border);

		&--loading
		{
			transform: scale(0.9);
			background-color: $darkGray;
			border: 1px solid $black;
		}
	}

	.invite-code__info
	{
		display: flex;
		flex-direction: column;
	}

	.invite-code__code { cursor: pointer; }

	.invite-code__code,
	.invite-code__uses
	{
		column-gap: 5px;

		display: flex;
		align-items: center;

		&-icon { color: $gray; }
	}

	.invite-code__delete { align-self: flex-end; }

	.invite-codes__button { font-weight: 500; }
</style>