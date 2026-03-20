<script setup lang="ts">
	const { auth }          = useTelegramAuth();
	const userStore         = useUserStore();

	const getProfile = async () =>
	{
		try { await userStore.fetchProfile(); }
		catch (err) { console.error(err); }
	};

	await auth();
	await getProfile();
</script>

<template>
	<div class="profile-page">
		<ProfileUserInfo v-if="userStore.userData" />
		<ProfileSwitcher />
	</div>
</template>

<style lang="scss">
	.profile-page
	{
		row-gap: 20px;

		display: flex;
		flex-direction: column;
	}
</style>