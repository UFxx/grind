<script setup lang="ts">
	import { useFetchErrors } from '~/composables/useFetchErrors';
	import { type AddInviteCode } from '~/types/inviteCode';

	const userStore       = useUserStore();
	const { togglePopup } = usePopupsStore();

	const inviteCodeData = ref<AddInviteCode>(
		{
			code    : undefined,
			maxUses : undefined
		}
	);

	const addInviteCode = async () =>
	{
		try { await userStore.addInviteCode(inviteCodeData.value); }
		catch (err) { useFetchErrors(err) }
		finally { closePopup(); }
	};

	const closePopup = () => togglePopup('AddInviteCode', false);
</script>

<template>
	<div class="add-invite-code">
		<div class="add-invite-code__header">
			<p class="add-invite-code__header-text">Добавить новый код</p>
			<button @click="closePopup" class="add-invite-code__header-icon">
				<IconsClose />
			</button>
		</div>
		<div class="add-invite-code__content">
			<UiInput
				placeholder="Введите код"
				v-model="inviteCodeData.code"
			/>
			<UiInput
				type="number"
				:onlyNumbers="true"
				placeholder="Макс. применений"
				inputmode="numeric"
				v-model.number="inviteCodeData.maxUses"
			/>
			<UiButton @click="addInviteCode" color="white">Добавить</UiButton>
		</div>
	</div>
</template>

<style lang='scss'>
	.add-invite-code
	{
		width: 100%;
		row-gap: 20px;
		padding: 10px;
		border-radius: 10px;
		background-color: $black;
		box-shadow: 0 0 6px 2px $darkGray;

		display: flex;
		position: relative;
		flex-direction: column;
	}

	.add-invite-code__header-text
	{
		font-size: 12px;
		text-align: center;
	}

	.add-invite-code__header-icon
	{
		cursor: pointer;
		color: $darkGray;

		top: 13px;
		right: 11px;
		position: absolute;
	}

	.add-invite-code__content
	{
		row-gap: 10px;

		display: flex;
		flex-direction: column;

		.ui-button
		{
			width: 100%;
			padding: 5px 10px;
			font-weight: 700;
		}
	}
</style>