<script setup lang="ts">
	const { auth } = useTelegramAuth();
	await auth();
</script>

<template>
	<UserInfo />
</template>

<style lang='scss' scoped>
</style>