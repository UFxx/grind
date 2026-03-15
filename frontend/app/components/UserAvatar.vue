<script setup lang="ts">
	const props = defineProps(
		{
			image:
			{
				type: String,
				required: true
			},
			levelPercent:
			{
				type: Number,
				required: true
			},
			level:
			{
				type: Number,
				required: true
			}
		}
	);
</script>

<template>
	<div class="user-avatar">
		<div class="user-avatar__image-wr">
			<img
				class="user-avatar__image"
				src=""
				alt="user photo"
			>
		</div>
		<span class="user-avatar__percentage"></span>
		<span class="user-avatar__level"></span>
	</div>
</template>

<style lang='scss' scoped>
	.user-avatar
	{

	}
</style>