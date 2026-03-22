<script setup lang="ts">
	import { type FormattedSubskill } from '~/types/skill';

	const props = defineProps<{ subItem: FormattedSubskill }>();

	const levelPercent = computed(() => getLevelPercent(props.subItem));
</script>

<template>
	<div class="profile-skills__subitem">
		<p>{{ subItem.name }}</p>
		<div class="profile-skills__subitem-progress">
			<div class="profile-skills__subitem-progress-header">
				<div class="profile-skills__subitem-progress-header-item">
					<p>
						<span>{{ subItem.totalXp }}</span>
						<span class="profile-skills__subitem-progress-header-item-xp">xp</span>
					</p>
					<p class="profile-skills__subitem-progress-header-percent">
						<span>{{ levelPercent }}</span>
						<span class="profile-skills__subitem-progress-header-item-xp">%</span>
					</p>
					<p>
						<span>{{ subItem.nextLevelStartXp }}</span>
						<span class="profile-skills__subitem-progress-header-item-xp">xp</span>
					</p>
				</div>
			</div>
			<UiProgressBar :progress="levelPercent" color="blue" />
		</div>
	</div>
</template>

<style lang='scss'>
	.profile-skills__subitem
	{
		row-gap: 5px;
		font-size: 12px;

		display: flex;
		flex-direction: column;
	}

	.profile-skills__subitem-progress
	{
		row-gap: 5px;

		display: flex;
		flex-direction: column;
	}

	.profile-skills__subitem-progress-header-item
	{
		column-gap: 5px;

		display: flex;
		justify-content: space-between;
	}

	.profile-skills__subitem-progress-header-item-xp,
	.profile-skills__subitem-progress-header-percent { color: $gray; }

	.profile-skills__subitem-progress-header-percent .profile-skills__subitem-progress-header-item-xp { color: $darkGray; };
</style>