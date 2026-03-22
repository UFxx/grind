<script setup lang="ts">
	const userStore       = useUserStore();
	const { togglePopup } = usePopupsStore();

	const deleteCode = async (id: string) =>
	{
		try { await userStore.deleteInviteCode(id); }
		catch (err) { console.error(err); }
	};

	const openAddInviteCodePopup = () => togglePopup('AddInviteCode', true);
</script>

<template>
	<div class="invite-codes">
		<div
			v-if="!userStore.userInviteCodes.length"
			class="invite-codes__empty"
		>
			<p class="invite-codes__empty-text">Пока что нет пригласительных кодов</p>
		</div>

		<TransitionGroup name="fade">
			<ProfileInviteCodesItem
				v-for="inviteCode in userStore.userInviteCodes"
				:key="inviteCode.id"
				:inviteCode
				@deleteCode="deleteCode"
			/>
		</TransitionGroup>
	</div>

	<UiButton
		@click="openAddInviteCodePopup"
		color="white"
		class="invite-codes__button"
	>
		Добавить новый код
	</UiButton>
</template>

<style lang='scss' scoped>
	.invite-codes
	{
		gap: 4px;

		display: grid;
		grid-template-rows: auto;
		grid-template-columns: repeat(2, 1fr);

		&:has(&__empty) { display: block; }
	}

	.invite-codes__empty
	{
		padding: 10px;

		display: flex;
		align-items: center;
		justify-content: center;
	}
</style>