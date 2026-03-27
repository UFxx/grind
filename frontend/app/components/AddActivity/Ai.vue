<script setup>
	const model = defineModel();
</script>

<template>
	<div class="add-activity-ai">
		<UiInput
			placeholder="Activity description"
			v-model="model"
			:fullWidth="true"
		/>
	</div>
</template>