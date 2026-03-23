<script setup lang="ts">
	const userStore       = useUserStore();
	const { togglePopup } = usePopupsStore();

	const loadingInviteCodeIds = ref<string[]>([]);

	const deleteCode = async (id: string) =>
	{
		loadingInviteCodeIds.value.push(id);

		try { await userStore.deleteInviteCode(id); }
		catch (err) { console.error(err); }
		finally { loadingInviteCodeIds.value = loadingInviteCodeIds.value.filter(loadingId => loadingId !== id); }
	};

	const openAddInviteCodePopup = () => togglePopup('AddInviteCode', true);
</script>

<template>
	<div class="invite-codes">
		<div
			v-if="!userStore.userInviteCodes.length"
			class="invite-codes__empty"
		>
			<p class="invite-codes__empty-text">There are no invitation codes yet</p>
		</div>

		<TransitionGroup name="fade">
			<ProfileInviteCodesItem
				v-for="(inviteCode, idx) in userStore.userInviteCodes"
				:key="inviteCode.id"
				:inviteCode
				:loadingInviteCodeIds
				@deleteCode="deleteCode"

				v-motion-slide-top
				:duration="200"
				:delay="idx * 50"
			/>
		</TransitionGroup>
	</div>

	<UiButton
		@click="openAddInviteCodePopup"
		color="white"
		class="invite-codes__button"

		v-motion-pop
		:duration="200"
		:delay="300"
	>
		Add new code
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