<script setup lang="ts">
	import { useFetchErrors } from '~/composables/useFetchErrors';
	import { type AddInviteCode } from '~/types/inviteCode';

	const userStore       = useUserStore();
	const { togglePopup } = usePopupsStore();
	const { addToast }    = useToastsStore();

	const inviteCodeData = ref<AddInviteCode>(
		{
			code    : undefined,
			maxUses : undefined
		}
	);

	const isLoading = ref(false);

	const addInviteCode = async () =>
	{
		isLoading.value = true;

		try
		{
			const response = await userStore.addInviteCode(inviteCodeData.value);

			if (!response.data.length)
			{
				addToast('success', 'Code added successfully');
				closePopup();
			}
		}
		catch (err)
		{
			const { errors, statusCode } = useFetchErrors(err);

			if (statusCode.value !== 200 && errors.value?.other?.length)
				addToast('error', errors.value.other);
			else
				addToast('error', 'An error occurred');
		}
		finally { isLoading.value = false; }
	};

	const closePopup = () => togglePopup('AddInviteCode', false);
</script>

<template>
	<div class="add-invite-code">
		<div class="add-invite-code__header">
			<p class="add-invite-code__header-text">Add new code</p>
			<IconsClose @click="closePopup" class="add-invite-code__header-icon" />
		</div>
		<div class="add-invite-code__content">
			<UiInput
				placeholder="Enter code"
				v-model="inviteCodeData.code"
			/>
			<UiInput
				type="number"
				:onlyNumbers="true"
				placeholder="Max. uses number"
				inputmode="numeric"
				v-model.number="inviteCodeData.maxUses"
			/>
			<UiButton
				@click="addInviteCode"
				color="white"
				:disabled="isLoading"
			>
				ADD
			</UiButton>
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
		background-color: $primary;
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
		opacity: 0.5;
		cursor: pointer;

		top: 13px;
		right: 11px;
		position: absolute;

		@include tr(.3, opacity);

		&:hover { opacity: 1; }
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