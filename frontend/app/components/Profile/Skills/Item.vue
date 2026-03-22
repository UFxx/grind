<script setup lang="ts">
	import { type FormattedSkill } from '~/types/skill';
	const props = defineProps<{ skill: FormattedSkill }>();

	const levelPercent = computed(() => getLevelPercent(props.skill));
</script>

<template>
	<div
		class="profile-skills__item"
	>
		<div class="profile-skills__item-header">
			<p class="profile-skills__item-header-title">{{ skill.name }}</p>
			<div class="profile-skills__item-header-lvl">
				{{ skill.currentLevel }}
				<span class="profile-skills__item-header-lvl-label">lvl</span>
			</div>
		</div>

		<div class="profile-skills__item-progress">
			<div class="profile-skills__item-progress-header">
				<p>
					<span>{{ skill.totalXp }}</span>
					<span class="profile-skills__item-progress-header-xp-label">xp</span>
				</p>
				<p class="profile-skills__item-progress-header-percent">
					<span>{{ levelPercent }}</span>
					<span class="profile-skills__item-progress-header-xp-label">%</span>
				</p>
				<p>
					<span>{{ skill.nextLevelStartXp }}</span>
					<span class="profile-skills__item-progress-header-xp-label">xp</span>
				</p>
			</div>
			<UiProgressBar :progress="levelPercent" />
		</div>

		<div class="profile-skills__subitems">
			<ProfileSkillsSubItem
				v-for="subItem in skill.subskills"
				:key="subItem.id"
				:subItem
			/>
		</div>
	</div>
</template>

<style lang='scss' scoped>
	.profile-skills__item
	{
		row-gap: 10px;
		padding: 10px;
		border-radius: 5px;
		background-color: $black;
		border-left: 1px solid white;

		display: flex;
		flex-direction: column;
	}

	.profile-skills__item-header
	{
		display: flex;
		align-items: center;
		justify-content: space-between;
	}

	.profile-skills__item-header-title
	{
		font-weight: 500;
		line-height: 19px;
	}

	.profile-skills__item-header-lvl
	{
		padding: 5px 10px;
		border-radius: 5px;
		background-color: $darkGray;
	}

	.profile-skills__item-header-lvl-label { color: $gray; }

	.profile-skills__item-progress
	{
		row-gap: 5px;

		display: flex;
		flex-direction: column;
	}

	.profile-skills__item-progress-header
	{
		display: flex;
		justify-content: space-between;

		p { line-height: 16px; }
	}

	.profile-skills__item-progress-header-xp-label { color: $gray; }
	.profile-skills__item-progress-header-percent
	{
		color: $gray;

		.profile-skills__item-progress-header-xp-label { color: $darkGray; }
	}

	.profile-skills__subitems
	{
		row-gap: 10px;
		padding: 10px;
		border-left: 1px solid $darkGray;

		display: flex;
		flex-direction: column;
	}
</style>