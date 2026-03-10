export const useTodosStore = defineStore('todos', () =>
{
	const someString = ref<string>('I\'m string');

	const logSomeString = () => console.log(someString.value);

	return { logSomeString };
});
