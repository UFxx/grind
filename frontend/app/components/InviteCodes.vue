<script setup lang="ts">
	import { type FormattedInviteCode } from '~/types/inviteCode';

	defineProps<{ inviteCodes: FormattedInviteCode[] }>();

	const copyCode = (code: FormattedInviteCode) => window.navigator.clipboard.writeText(code.code);
</script>

<template>
	<div class="invite-codes-wr">
		<div class="invite-codes">
			<div
				v-for="code in inviteCodes"
				class="invite-code"
			>
				<div
					@click="copyCode(code)"
					class="invite-code__code"
				>
					<span class="invite-code__code-icon"><IconsCopy /></span>
					<span class="invite-code__code-label">{{ code.code }}</span>
				</div>
				<div class="invite-code__uses">
					<span class="invite-code__uses-icon"><IconsUsers /></span>
					<span class="invite-code__uses-label">{{ code.uses }}/{{ code.maxUses }}</span>
				</div>
			</div>
		</div>
		<UiButton color="white" class="invite-codes__button">Добавить новый код</UiButton>
	</div>
</template>

<style lang='scss' scoped>
	.invite-codes-wr
	{
		row-gap: 5px;
		padding: 3px;
		border-radius: 5px;
		background-color: $darkGray;

		display: flex;
		flex-direction: column;
	}

	.invite-codes
	{
		gap: 4px;

		display: grid;
		grid-template-rows: auto;
		grid-template-columns: repeat(2, 1fr);
	}

	.invite-code
	{
		row-gap: 5px;
		padding: 10px;
		font-weight: 500;
		border-radius: 5px;
		background-color: $black;

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

	.invite-codes__button { font-weight: 500; }
</style>