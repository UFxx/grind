<script setup lang="ts">
	import { type FormattedUser } from '~/types/user';

	const { auth } = useTelegramAuth();
	const { user: userApi } = useApi();

	const user = ref<FormattedUser | null>(null);

	const getProfile = async () =>
	{
		try
		{
			const { data } = await userApi.getProfile();

			user.value = userSerializer(data)
		}
		catch (err) { console.error(err); }
	}

	await auth();
	await getProfile();
</script>

<template>
	<UserInfo
		v-if="user"
		:user
	/>
</template>

<style lang='scss' scoped></style>