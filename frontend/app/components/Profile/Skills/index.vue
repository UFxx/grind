<script setup lang="ts">
	import { type Skill } from '~/types/skill';

	const userStore = useUserStore();

	const mockSkills: Skill[]  = [
		{
			id: "550e8400-e29b-41d4-a716-446655440004",
			name: "Endurance",
			total_xp: 1500,
			current_level: 3,
			next_level: 4,
			current_level_start_xp: 1000,
			next_level_start_xp: 2000,
			xp_to_next_level: 500,
			items: [
				{
					id: "550e8400-e29b-41d4-a716-446655440005",
					name: "Running",
					total_xp: 800,
					current_level: 2,
					next_level: 4,
					current_level_start_xp: 500,
					next_level_start_xp: 1500,
					xp_to_next_level: 700
				},
				{
					id: "550e8400-e29b-41d4-a716-446655440006",
					name: "Swimming",
					total_xp: 700,
					current_level: 2,
					next_level: 3,
					current_level_start_xp: 300,
					next_level_start_xp: 1200,
					xp_to_next_level: 500
				}
			]
		}
	];

	userStore.setUserSkills(skillSerializer(mockSkills));
</script>

<template>
	<div v-if="userStore.userSkills.length" class="profile-skills">
		<ProfileSkillsItem
			v-for="skill in userStore.userSkills"
			:key="skill.id"
			:skill
		/>
	</div>

	<div v-else class="profile-skills__empty">
		<p class="profile-skills__empty-text">Пока что нет навыков</p>
	</div>
</template>

<style lang='scss' scoped>
	.profile-skills
	{
		row-gap: 4px;

		display: flex;
		flex-direction: column;
	}

	.profile-skills__empty
	{
		padding: 10px;

		display: flex;
		align-items: center;
		justify-content: center;
	}
</style>