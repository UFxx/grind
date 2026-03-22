<script setup lang="ts">
	const userStore = useUserStore();

	const fetchSkills = async () =>
	{
		try { await userStore.fetchUserSkills() }
		catch (err) { console.error(err) }
	};

	fetchSkills();
</script>

<template>
	<div
		v-if="userStore.userSkills.length"
		class="profile-skills"
	>
		<ProfileSkillsItem
			v-for="(skill, idx) in userStore.userSkills"
			:key="skill.id"
			:skill
			:idx="idx"
			v-motion-slide-visible-left
			:duration="200"
			:delay="idx * 50"
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