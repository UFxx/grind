<script setup lang="ts">
	import { type FormattedActivityTag } from '~/types/activity';

	const props = defineProps<{ tag: FormattedActivityTag }>();

	const tagColor = computed(() =>
		{
			switch (props.tag.name)
			{
				case 'New':
					return 'orange'
				case 'Hard':
					return 'red'
				case 'Impact':
					return 'green'
				default:
					break;
			}
		}
	);
</script>

<template>
	<div
		v-if="tag.isActive"
		class="activity-tag"
		:class="`activity-tag--${tagColor}`"
	>
		<slot />
	</div>
</template>

<style lang='scss' scoped>
	.activity-tag
	{
		font-size: 10px;
		padding: 3px 5px;
		line-height: 14px;
		border-radius: 5px;
		width: max-content;

		&--green
		{
			color: $green;
			background-color: rgba($green, 0.1);
		}

		&--red
		{
			color: $red;
			background-color: rgba($red, 0.1);
		}

		&--orange
		{
			color: $orange;
			background-color: rgba($orange, 0.1);
		}
	}
</style>