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
			<div
				class="user-avatar__xp-counter"
				:style="`--progress: ${levelPercent}%`"
			/>
			<img
				class="user-avatar__image"
				:src="image"
				alt="user photo"
			>
		</div>
		<div class="user-avatar__percentage">
			<span class="user-avatar__percentage-number">{{ levelPercent }}</span>
			<span class="user-avatar__percentage-label">%</span>
		</div>
		<div class="user-avatar__level">
			<span class="user-avatar__level-number">{{ level }}</span>
			<span class="user-avatar__level-label">lvl</span>
		</div>
	</div>
</template>

<style lang='scss' scoped>
	.user-avatar { position: relative; }

	.user-avatar__image-wr
	{
		width: 100px;
		height: 100px;

		position: relative;

		img
		{
			z-index: 2;
			width: 100%;
			height: 100%;
			object-fit: cover;
			border-radius: 100%;
			object-position: top;

			position: relative;
		}
	}

	.user-avatar__xp-counter
	{
		z-index: 1;
		width: 103.5px;
		height: 103.5px;
		border-radius: 100%;
		background-image: conic-gradient(
			from 0turn at 50% 50%,
			$white var(--progress),
			transparent 0%
		);

		top: 50%;
		left: 50%;
		position: absolute;
		transform: translate(-50%, -50%);
	}

	.user-avatar__percentage-number,
	.user-avatar__level-number { font-size: 12px; }

	.user-avatar__percentage,
	.user-avatar__level
	{
		font-size: 8px;

		bottom: -10px;
		position: absolute;

		&-label { color: $gray; }
	}

	.user-avatar__percentage { left: -10px; }
	.user-avatar__level { right: -10px; }
</style>