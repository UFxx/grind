<script setup lang="ts">
	import { type FormattedInviteCode } from '~/types/inviteCode';

	const userStore = useUserStore();

	const copyCode = (code: FormattedInviteCode) => window.navigator.clipboard.writeText(code.code);

	const deleteCode = async (id: string) =>
	{
		try { await userStore.deleteInviteCode(id); }
		catch (err) { console.error(err); }
	};

	const addCode = async () =>
	{
		try { await userStore.addInviteCode(`${Date.now()}`, 1); }
		catch (err) { console.error(err); }
	};
</script>

<template>
	<div class="invite-codes-wr">
		<div
			v-if="userStore.userInviteCodes.length"
			class="invite-codes"
		>

			<TransitionGroup name="fade">
				<div
					v-for="code in userStore.userInviteCodes"
					:key="code.id"
					class="invite-code"
				>
					<div class="invite-code__info">
						<div
							@click="copyCode(code)"
							class="invite-code__code"
						>
							<span class="invite-code__code-icon"><IconsCopy /></span>
							<p class="invite-code__code-label">{{ code.code }}</p>
						</div>
						<div class="invite-code__uses">
							<span class="invite-code__uses-icon"><IconsUsers /></span>
							<span class="invite-code__uses-label">{{ code.uses }}/{{ code.maxUses }}</span>
						</div>
					</div>
					<button
						@click="deleteCode(code.id)"
						class="invite-code__delete"
					>
						<IconsDelete />
					</button>
				</div>
			</TransitionGroup>

		</div>

		<div
			v-else
			class="invite-codes__empty"
		>
			<p class="invite-codes__empty-text">Пока что нет пригласительных кодов</p>
		</div>
		<UiButton @click="addCode" color="white" class="invite-codes__button">Добавить новый код</UiButton>
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

	.invite-codes__empty
	{
		padding: 10px;

		display: flex;
		align-items: center;
		justify-content: center;
	}

	.invite-code
	{
		row-gap: 5px;
		padding: 10px;
		font-weight: 500;
		border-radius: 5px;
		background-color: $black;

		display: flex;
		justify-content: space-between;
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