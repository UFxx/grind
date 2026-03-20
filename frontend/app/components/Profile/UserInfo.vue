<script setup lang="ts">
	const userStore = useUserStore();
</script>

<template>
	<div class="user-info">
		<UserAvatar
			:image="userStore.userData?.avatarUrl"
			:level="userStore.userData?.level"
		/>
		<div class="user-info__text">
			<p class="user-info__tag">@{{ userStore.userData?.name }}</p>
			<p class="user-info__rank">{{ userStore.userData?.rank.name }}</p>
		</div>
	</div>
</template>

<style lang='scss' scoped>
	.user-info
	{
		row-gap: 10px;

		display: flex;
		align-items: center;
		flex-direction: column;
	}

	.user-info__text
	{
		row-gap: 2px;

		display: flex;
		align-items: center;
		flex-direction: column;
	}

	.user-info__tag { line-height: 19px; }

	.user-info__rank
	{
		color: $gray;
		font-size: 12px;
		line-height: 16px;
	}
</style>