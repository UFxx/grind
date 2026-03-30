<script setup>
	const inviteCode = ref('');

	const auth = () => navigateTo(`/?inviteCode=${inviteCode.value}`);
</script>

<template>
	<div class="auth">
		<div class="auth-content">
			<div class="auth-content__title-wr">
				<p class="auth-content__title">Please enter the invitation code</p>
				<p class="auth-content__subtitle">I don't see you've been invited...</p>
			</div>
			<div class="auth-content__content">
				<UiInput
					placeholder="Invite code"
					:full-width="true"
					v-model="inviteCode"
				/>
				<UiButton
					color="white"
					:fullWidth="true"
					@click="auth"
				>
					Enter
				</UiButton>
			</div>
		</div>
	</div>
</template>

<style lang='scss'>
	.auth
	{
		height: calc(100vh - 40px);
		row-gap: 10px;
		padding: 0 10px;

		display: flex;
		align-items: center;
		justify-content: center;
	}

	.auth-content
	{
		width: 100%;
		row-gap: 10px;

		display: flex;
		flex-direction: column;
	}

	.auth-content__title-wr
	{
		display: flex;
		align-items: center;
		flex-direction: column;
	}

	.auth-content__title
	{
		font-weight: 500;
		line-height: 18px;
	}

	.auth-content__subtitle
	{
		color: $gray;
		font-size: 12px;
		line-height: 18px;
	}

	.auth-content__content
	{
		row-gap: 10px;
		padding: 0 10px;

		display: flex;
		flex-direction: column;
	}
</style>