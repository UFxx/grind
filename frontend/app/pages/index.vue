<script setup lang="ts">
	import { type ITodo } from '~/types/todo';

	const { todo } = useApi();
	const { logSomeString } = useTodosStore();

	const todoList = ref<ITodo[] | null>(null);

	const fetchTodos = async () =>
	{
		const data     = await todo.fetchTodo();
		todoList.value = data;
	}

	await fetchTodos();
	logSomeString();
</script>

<template></template>

<style lang='scss'></style>